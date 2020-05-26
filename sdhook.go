// Package sdhook provides a logrus compatible logging hook for Google
// Stackdriver logging.
package sdhook

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"

	"cloud.google.com/go/errorreporting"
	"cloud.google.com/go/logging"
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	mrpb "google.golang.org/genproto/googleapis/api/monitoredres"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
)

const (
	// DefaultName is the default name passed to LogName when using serviceClient
	// account credentials.
	DefaultName = "default"
)

var MaxLogWriteDuration = 15 * time.Second

// StackdriverHook provides a logrus hook to Google Stackdriver logging.
type StackdriverHook struct {
	// levels are the levels that logrus will hook to.
	levels []logrus.Level

	// projectID is the projectID
	projectID string

	// serviceClient is the logging serviceClient.
	serviceClient *logging.Client

	logger *logging.Logger

	// serviceClient is the error reporting serviceClient.
	errorClient *errorreporting.Client

	// resource is the monitored resource.
	resource *mrpb.MonitoredResource

	// logName is the name of the log.
	logName string

	// labels are the labels to send with each log entry.
	labels map[string]string

	// partialSuccess allows partial writes of log entries if there is a badly
	// formatted log.
	partialSuccess bool

	// agentClient defines the fluentd logger object that can send data to
	// to the Google logging agent.
	agentClient *fluent.Fluent

	// errorReportingServiceName defines the value of the field <serviceClient>,
	// required for a valid error reporting payload. If this value is set,
	// messages where level/severity is higher than or equal to "error" will
	// be sent to Stackdriver error reporting.
	// See more at:
	// https://cloud.google.com/error-reporting/docs/formatting-error-messages
	errorReportingServiceName string

	// errorReportingLogName is the name of the log for error reporting.
	// It must contain the string "error"
	// If not given, the string "<logName>_error" is used.
	errorReportingLogName string

	// m synchronizes access to agentClient and logger for flushing and closing.
	m sync.RWMutex

	// googleOptions Google client options when creating StackDriver connection.
	googleOptions []option.ClientOption

	// syncLevel ensures that given level logs are send synchronously.
	syncLevel map[logrus.Level]bool
}

// New creates a StackdriverHook using the provided options that is suitible
// for using with logrus for logging to Google Stackdriver.
func New(opts ...Option) (*StackdriverHook, error) {
	var err error

	sh := &StackdriverHook{
		levels:    logrus.AllLevels,
		syncLevel: make(map[logrus.Level]bool),
	}

	// apply opts
	for _, o := range opts {
		err = o(sh)
		if err != nil {
			return nil, err
		}
	}

	// // check serviceClient, resource, logName set
	// if sh.serviceClient == nil && sh.agentClient == nil {
	// 	return nil, errors.New("no stackdriver serviceClient was provided")
	// }
	if sh.resource == nil && sh.agentClient == nil {
		return nil, errors.New("the monitored resource was not provided")
	}
	if sh.projectID == "" && sh.agentClient == nil {
		return nil, errors.New("the project id was not provided")
	}

	// set default project name
	if sh.logName == "" {
		if err = LogName(DefaultName)(sh); err != nil {
			return nil, err
		}
	}

	// If error reporting log name not set, set it to log name
	// plus string suffix
	if sh.errorReportingLogName == "" {
		sh.errorReportingLogName = sh.logName + "_errors"
	}

	if sh.serviceClient == nil {
		// create logging serviceClient
		l, err := logging.NewClient(context.Background(), sh.projectID, sh.googleOptions...)
		if err != nil {
			return nil, err
		}
		sh.serviceClient = l
	}

	if sh.errorClient == nil {
		// create error reporting serviceClient
		c, err := errorreporting.NewClient(context.Background(), sh.projectID, errorreporting.Config{ServiceName: sh.logName}, sh.googleOptions...)
		if err != nil {
			return nil, err
		}
		sh.errorClient = c
	}

	if sh.logger == nil {
		sh.logger = sh.serviceClient.Logger(sh.logName)
	}

	return sh, nil
}

func isError(entry *logrus.Entry) bool {
	if entry != nil {
		switch entry.Level {
		case logrus.ErrorLevel:
			return true
		case logrus.FatalLevel:
			return true
		case logrus.PanicLevel:
			return true
		}
	}
	return false
}

// Levels returns the logrus levels that this hook is applied to. This can be
// set using the Levels Option.
func (sh *StackdriverHook) Levels() []logrus.Level {
	return sh.levels
}

func severity(level logrus.Level) logging.Severity {
	switch level {
	case logrus.TraceLevel:
		return logging.Debug
	case logrus.DebugLevel:
		return logging.Debug
	case logrus.InfoLevel:
		return logging.Info
	case logrus.WarnLevel:
		return logging.Warning
	case logrus.ErrorLevel:
		return logging.Error
	case logrus.FatalLevel:
		return logging.Critical
	case logrus.PanicLevel:
		return logging.Emergency
	}

	return logging.Default
}

func (sh *StackdriverHook) send(entry *logrus.Entry, callstack []byte, sync bool) {
	defer sh.m.RUnlock()

	var httpReq *logging.HTTPRequest
	// convert entry data to labels
	labels := make(map[string]string, len(entry.Data))
	for k, v := range entry.Data {
		switch x := v.(type) {
		case string:
			labels[k] = x

		case *http.Request:
			httpReq = &logging.HTTPRequest{
				Request: x,
			}

		case *logging.HTTPRequest:
			httpReq = x

		default:
			labels[k] = fmt.Sprintf("%v", v)
		}
	}

	// write log entry
	if sh.agentClient != nil {
		sh.sendLogMessageViaAgent(entry, labels, httpReq, callstack, sync)
	} else {
		sh.sendLogMessageViaAPI(entry, labels, httpReq, callstack, sync)
	}
}

// Fire writes the message to the Stackdriver entry serviceClient.
func (sh *StackdriverHook) Fire(entry *logrus.Entry) error {
	sh.m.RLock()
	var callstack []byte
	if sh.errorReportingServiceName != "" && isError(entry) {
		// callstack = stack.Callers(8)
		var buf [20 * 1024]byte
		callstack = []byte(chopStack(buf[0:runtime.Stack(buf[:], false)]))
	}

	if sh.syncLevel[entry.Level] {
		sh.send(sh.copyEntry(entry), callstack, true)
	} else {
		go sh.send(sh.copyEntry(entry), callstack, false)
	}

	return nil
}

// Wait will return after all subroutines have returned.
// Use in conjunction with logrus return handling to ensure all of
// your logs are delivered before your program exits.
// `logrus.RegisterExitHandler(h.Wait)`
func (sh *StackdriverHook) Wait() {
	sh.m.Lock()
	defer sh.m.Unlock()
	if sh.agentClient != nil {
		if err := sh.agentClient.Close(); err != nil {
			log.Printf("failed to close agent client (via Fluentd): %s", err)
		}
	}
	if err := sh.logger.Flush(); err != nil {
		log.Printf("failed to flush logs (via API): %s", err)
	}
}

func (sh *StackdriverHook) copyEntry(entry *logrus.Entry) *logrus.Entry {
	e := *entry
	e.Data = make(logrus.Fields, len(entry.Data))
	for k, v := range entry.Data {
		e.Data[k] = v
	}
	return &e
}

func (sh *StackdriverHook) sendLogMessageViaAgent(entry *logrus.Entry,
	labels map[string]string, httpReq *logging.HTTPRequest,
	callstack []byte, sync bool) {

	// The log entry payload schema is defined by the Google fluentd
	// logging agent. See more at:
	// https://github.com/GoogleCloudPlatform/fluent-plugin-google-cloud
	logEntry := map[string]interface{}{
		"severity":         severity(entry.Level).String(),
		"timestampSeconds": strconv.FormatInt(entry.Time.Unix(), 10),
		"timestampNanos":   strconv.FormatInt(entry.Time.UnixNano()-entry.Time.Unix()*1000000000, 10),
		"message":          entry.Message,
	}
	for k, v := range labels {
		logEntry[k] = v
	}
	if httpReq != nil {
		logEntry["httpRequest"] = httpReq
	}
	// The error reporting payload JSON schema is defined in:
	// https://cloud.google.com/error-reporting/docs/formatting-error-messages
	// Which reflects the structure of the ErrorEvent type in:
	// https://godoc.org/google.golang.org/api/clouderrorreporting/v1beta1
	if sh.errorReportingServiceName != "" && isError(entry) {
		errorEvent := sh.buildErrorReportingEvent(entry, callstack, httpReq)
		errorStructPayload, err := json.Marshal(errorEvent)
		if err != nil {
			log.Printf("error marshaling error reporting data: %s", err.Error())
		}
		var errorJSONPayload map[string]interface{}
		err = json.Unmarshal(errorStructPayload, &errorJSONPayload)
		if err != nil {
			log.Printf("error parsing error reporting data: %s", err.Error())
		}
		for k, v := range logEntry {
			errorJSONPayload[k] = v
		}
		if err := sh.agentClient.Post(sh.errorReportingLogName, errorJSONPayload); err != nil {
			log.Printf("error posting error reporting entries to logging agent: %s", err.Error())
		}
	} else if err := sh.agentClient.Post(sh.logName, logEntry); err != nil {
		log.Printf("error posting log entries to logging agent: %s", err.Error())
	}
}

func chopStack(s []byte) string {
	toSkip := []byte("github.com/sirupsen/logrus.")

	lfFirst := bytes.IndexByte(s, '\n')
	if lfFirst == -1 {
		return string(s)
	}
	stack := s[lfFirst+1:]
	var found bool
	for true {
		nextLine := bytes.IndexByte(stack, '\n')
		if nextLine == -1 {
			return string(s)
		}
		toSkipIndex := bytes.Index(stack[:nextLine], toSkip)
		if found && toSkipIndex == -1 {
			break
		} else if toSkipIndex >= 0 {
			found = true
		}
		stack = stack[nextLine+1:]
		nextLine = bytes.IndexByte(stack, '\n')
		stack = stack[nextLine+1:]
	}
	return string(s[:lfFirst+1]) + string(stack)
}

func (sh *StackdriverHook) buildErrorReportingEvent(entry *logrus.Entry, callstack []byte, httpReq *logging.HTTPRequest) errorreporting.Entry {
	var r *http.Request
	if httpReq != nil {
		r = httpReq.Request
	}
	return errorreporting.Entry{
		Error: errors.New(entry.Message),
		Req:   r,
		User:  "",
		Stack: callstack,
	}
}

func (sh *StackdriverHook) sendLogMessageViaAPI(entry *logrus.Entry,
	labels map[string]string, httpReq *logging.HTTPRequest,
	callstack []byte, sync bool) {

	if sh.errorReportingServiceName != "" && isError(entry) {
		if sh != nil && sh.errorClient != nil {
			e := sh.buildErrorReportingEvent(entry, callstack, httpReq)
			if sync {
				ctx, canc := context.WithTimeout(context.Background(), MaxLogWriteDuration)
				defer canc()
				if err := sh.errorClient.ReportSync(ctx, e); err != nil {
					log.Println("cannot report event:", err)
				}
			} else {
				sh.errorClient.Report(e)
			}
		} else {
			log.Println("the error reporting serviceClient is not set")
		}
	} else {
		entrySd := logging.Entry{
			Timestamp:   entry.Time,
			Severity:    severity(entry.Level),
			Payload:     entry.Message,
			Labels:      labels,
			HTTPRequest: httpReq,
			LogName:     "",
			Resource:    sh.resource,
		}
		if c := entry.Caller; c != nil {
			entrySd.SourceLocation = &logpb.LogEntrySourceLocation{
				File:     c.File,
				Function: c.Function,
				Line:     int64(c.Line),
			}
		}
		if sync {
			ctx, canc := context.WithTimeout(context.Background(), MaxLogWriteDuration)
			defer canc()
			if err := sh.logger.LogSync(ctx, entrySd); err != nil {
				log.Printf("cannot write to log: %s", err)
			}
		} else {
			sh.logger.Log(entrySd)
		}
	}
}
