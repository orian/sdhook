package sdhook

import (
	"cloud.google.com/go/errorreporting"
	"cloud.google.com/go/logging"
	"context"
	"fmt"
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	mrpb "google.golang.org/genproto/googleapis/api/monitoredres"
)

// Option represents an option that modifies the Stackdriver hook settings.
type Option func(*StackdriverHook) error

// Levels is an option that sets the logrus levels that the StackdriverHook
// will create log entries for.
func Levels(levels ...logrus.Level) Option {
	return func(sh *StackdriverHook) error {
		sh.levels = levels
		return nil
	}
}

// ProjectID is an option that sets the project ID which is needed for the log
// name.
func ProjectID(projectID string) Option {
	return func(sh *StackdriverHook) error {
		sh.projectID = projectID
		return nil
	}
}

// EntriesService is an option that sets the Google API entry serviceClient to use
// with Stackdriver.
// func EntriesService(serviceClient *logging.EntriesService) Option {
// 	return func(sh *StackdriverHook) error {
// 		sh.serviceClient = serviceClient
// 		return nil
// 	}
// }

// LoggingService is an option that sets the Google API logging serviceClient to use.
func LoggingClient(service *logging.Client) Option {
	return func(sh *StackdriverHook) error {
		sh.serviceClient = service
		return nil
	}
}

// ErrorService is an option that sets the Google API error reporting serviceClient to use.
func ErrorService(errorService *errorreporting.Client) Option {
	return func(sh *StackdriverHook) error {
		sh.errorClient = errorService
		return nil
	}
}

// MonitoredResource is an option that sets the monitored resource to send with
// each log entry.
func MonitoredResource(resource *mrpb.MonitoredResource) Option {
	return func(sh *StackdriverHook) error {
		sh.resource = resource
		return nil
	}
}

// Resource is an option that sets the resource information to send with each
// log entry.
//
// Please see https://cloud.google.com/logging/docs/api/v2/resource-list for
// the list of labels required per ResType.
func Resource(typ ResType, labels map[string]string) Option {
	return func(sh *StackdriverHook) error {
		return MonitoredResource(&mrpb.MonitoredResource{
			Type:   string(typ),
			Labels: labels,
		})(sh)
	}
}

// LogName is an option that sets the log name to send with each log entry.
//
// Log names are specified as "projects/{projectID}/logs/{logName}"
// if the projectID is set. Otherwise, it's just "{logName}"
func LogName(name string) Option {
	return func(sh *StackdriverHook) error {
		// We don't need whole path, {logName} is enough as GCP writes rest of it by default.
		sh.logName = name
		return nil
	}
}

// ErrorReportingLogName is an option that sets the log name to send
// with each error message for error reporting.
// Only used when ErrorReportingService has been set.
func ErrorReportingLogName(name string) Option {
	return func(sh *StackdriverHook) error {
		sh.errorReportingLogName = name
		return nil
	}
}

// Labels is an option that sets the labels to send with each log entry.
func Labels(labels map[string]string) Option {
	return func(sh *StackdriverHook) error {
		sh.labels = labels
		return nil
	}
}

// PartialSuccess is an option that toggles whether or not to write partial log
// entries.
func PartialSuccess(enabled bool) Option {
	return func(sh *StackdriverHook) error {
		sh.partialSuccess = enabled
		return nil
	}
}

// ErrorReportingService is an option that defines the name of the serviceClient
// being tracked for Stackdriver error reporting.
// See:
// https://cloud.google.com/error-reporting/docs/formatting-error-messages
func ErrorReportingService(service string) Option {
	return func(sh *StackdriverHook) error {
		sh.errorReportingServiceName = service
		return nil
	}
}

// requiredScopes are the oauth2 scopes required for stackdriver logging.
var requiredScopes = []string{
	logging.WriteScope,
}

// sliceContains returns true if haystack contains needle.
func sliceContains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}

	return false
}

func GoogleLoggingAgent() Option {
	return func(sh *StackdriverHook) error {
		var err error
		// set agent client. It expects that the forward input fluentd plugin
		// is properly configured by the Google logging agent, which is by default.
		// See more at:
		// https://cloud.google.com/error-reporting/docs/setup/ec2
		sh.agentClient, err = fluent.New(fluent.Config{
			Async: true,
		})
		if err != nil {
			return fmt.Errorf("could not find fluentd agent on 127.0.0.1:24224: %v", err)
		}
		return nil
	}
}

func SyncLevels(levels ...logrus.Level) Option {
	return func(sh *StackdriverHook) error {
		if sh.syncLevel == nil {
			sh.syncLevel = make(map[logrus.Level]bool)
		}
		for _, l := range levels {
			sh.syncLevel[l] = true
		}
		return nil
	}
}

// GoogleDefaultCredentials returns the token source for
// "Application Default Credentials".
//
// It looks for credentials in the following places,
// preferring the first location found:
//
//   1. A JSON file whose path is specified by the
//      GOOGLE_APPLICATION_CREDENTIALS environment variable.
//   2. A JSON file in a location known to the gcloud command-line tool.
//      On Windows, this is %APPDATA%/gcloud/application_default_credentials.json.
//      On other systems, $HOME/.config/gcloud/application_default_credentials.json.
//   3. On Google App Engine standard first generation runtimes (<= Go 1.9) it uses
//      the appengine.AccessToken function.
//   4. On Google Compute Engine, Google App Engine standard second generation runtimes
//      (>= Go 1.11), and Google App Engine flexible environment, it fetches
//      credentials from the metadata server.
func GoogleDefaultCredentials() Option {
	return func(sh *StackdriverHook) error {
		creds, err := google.FindDefaultCredentials(context.Background())
		if err != nil {
			return err
		}

		sh.googleOptions = append(sh.googleOptions, option.WithCredentials(creds))
		sh.projectID = creds.ProjectID
		return nil
	}
}

func GoogleClientOption(opts ...option.ClientOption) Option {
	return func(sh *StackdriverHook) error {
		sh.googleOptions = append(sh.googleOptions, opts...)
		return nil
	}
}
