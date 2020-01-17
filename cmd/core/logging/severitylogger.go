package logging

import (
	"fmt"
	"io"
)

type Severity int

const (
	Debug Severity = 1
	Info  Severity = 2
	Warn  Severity = 3
	Error Severity = 4
)

type SeverityLogger interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Log(Severity, string)
}

func NewLogger(severity Severity, writer io.Writer) SeverityLogger {
	return &severityLogger{
		severity: severity,
		writer:   writer,
	}
}

type severityLogger struct {
	severity Severity
	writer   io.Writer
}

func (logger *severityLogger) Debug(message string) {
	logger.Log(Debug, message)
}

func (logger *severityLogger) Info(message string) {
	logger.Log(Info, message)
}

func (logger *severityLogger) Warn(message string) {
	logger.Log(Warn, message)
}

func (logger *severityLogger) Error(message string) {
	logger.Log(Error, message)
}

func (logger *severityLogger) Log(severity Severity, message string) {
	if logger.severity <= severity {
		_, err := fmt.Fprintln(logger.writer, message)

		if err != nil {
			panic(err)
		}
	}
}
