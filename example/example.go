package main

import (
	"fmt"
	"os"
	"time"

	"github.com/orian/sdhook"
	"github.com/sirupsen/logrus"
)

func main() {
	// create a logger with some fields
	logger := logrus.New()
	logger.WithFields(logrus.Fields{
		"my_field":  115888,
		"my_field2": 898858,
	})

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	logger.Infof("project id: %s", projectID)
	logger.SetReportCaller(true)

	{
		// create stackdriver hook
		hook, err := sdhook.New(
			sdhook.GoogleDefaultCredentials(),
			sdhook.ProjectID(projectID),
			sdhook.LogName("some_log"),
			sdhook.Resource(sdhook.ResTypeGenericTask, map[string]string{
				"project_id": projectID,
				"location":   "europe:poland:olsztyn:1",
				"namespace":  "default",
				"job":        "test",
				"task_id":    fmt.Sprintf("test-%d", time.Now().Unix()),
			}),
			sdhook.SyncLevels(logrus.PanicLevel, logrus.FatalLevel),
		)
		if err != nil {
			logger.Fatal(err)
		}

		// add to logrus
		logger.Hooks.Add(hook)
		logrus.RegisterExitHandler(hook.Wait)
	}

	{
		hook, err := sdhook.New(
			sdhook.GoogleDefaultCredentials(),
			sdhook.ErrorReportingLogName("some_log_error"),
			sdhook.ProjectID(projectID),
			sdhook.Resource(sdhook.ResTypeGenericTask, map[string]string{
				"project_id": projectID,
				"location":   "europe:poland:olsztyn:1",
				"namespace":  "default",
				"job":        "test",
				"task_id":    fmt.Sprintf("test-%d", time.Now().Unix()),
			}),
			sdhook.ErrorReportingService("generic-test-job"),
			sdhook.Levels(logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel),
			sdhook.SyncLevels(logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel),
		)
		if err != nil {
			logger.Fatal(err)
		}
		logger.Hooks.Add(hook)
		logrus.RegisterExitHandler(hook.Wait)
	}

	// log some message
	logger.Printf("a random message @ %s", time.Now().Format("15:04:05"))
	logger.Errorf("to jest sformatowany error %d", 123)

	// defer logrus.Exit(0)
	logger.Fatal("fatal error")
}
