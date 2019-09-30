# About sdhook

Package sdhook provides a [logrus](https://github.com/sirupsen/logrus)
compatible hook for [Google Stackdriver logging](https://cloud.google.com/logging/).

This is a hard fork of [github.com/knq/sdhook](https://github.com/knq/sdhook).
It adds semantic versioning and replaces deprecated
[google.golang.org/api/logging/v2](https://google.golang.org/api/logging/v2)
with [cloud.google.com/go/logging](https://cloud.google.com/go/logging)

# Installation

Install in the usual Go way:
```sh
go get -u github.com/orian/sdhook
```

# Usage

Simply create the hook, and add it to a logrus logger:

```go
// create hook using service account credentials
h, err := sdhook.New(
	sdhook.GoogleServiceAccountCredentialsFile("./credentials.json"),
)

// create logger with extra fields
//
// logrus fields will be converted to Stackdriver labels
logger := logrus.New().WithFields(logrus.Fields{
	"field1": 15,
	"field2": 20,
})

// add hook
logger.Hooks.Add(h)

// log something
logger.Printf("something %d", 15)
```

The example above sends log entries directly to the logging API. If you have the logging agent running, you can send log entries to it instead, with the added benefit of having extra instance metadata added to your log entries by the agent. In the example above, the initialization would simply be:

```go
// create hook using the logging agent
h, err := sdhook.New(
	sdhook.GoogleLoggingAgent(),
)
```

Please also see [example/example.go](example/example.go) for a more complete
example.

## Error Reporting

If you'd like to enable sending errors to Google's Error Reporting (https://cloud.google.com/error-reporting/), you have to set the name of the service, app or system you're running. Following the example above, the initialization would then be:

```go
// create hook using the logging agent
h, err := sdhook.New(
	sdhook.GoogleLoggingAgent(),
	sdhook.ErrorReportingService("your-great-app"),
)
```

The value of the `ErrorReportingService` function parameter above corresponds to the string value you'd like to see in the `service` field of the Error Reporting payload, as defined by https://cloud.google.com/error-reporting/docs/formatting-error-messages

Also note that, if you enable error reporting, errors and messages of more severe levels go into the error log and will not be displayed in the regular log. The error log name is either defined by the `ErrorReportingLogName` function or defaults to `<regular-log-name>_errors`. This fulfills Google's Error Reporting requirement that the log name should have the string `err` in its name. See more in: https://cloud.google.com/error-reporting/docs/setup/ec2

See [GoDoc](https://godoc.org/github.com/knq/sdhook) for a full API listing.
