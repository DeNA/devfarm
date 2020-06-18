package logging

import (
	"bytes"
	"fmt"
)

func NullSeverityLogger() SeverityLogger {
	return nullSeverityLogger{}
}

type nullSeverityLogger struct{}

var _ SeverityLogger = nullSeverityLogger{}

func (nullSeverityLogger) Debug(string) {}

func (nullSeverityLogger) Info(string) {}

func (nullSeverityLogger) Warn(string) {}

func (nullSeverityLogger) Error(string) {}

func (nullSeverityLogger) Log(Severity, string) {}

func SpySeverityLogger() *SeverityLoggerSpy {
	return &SeverityLoggerSpy{
		Logs: &bytes.Buffer{},
	}
}

type SeverityLoggerSpy struct {
	Logs *bytes.Buffer
}

var _ SeverityLogger = &SeverityLoggerSpy{}

func (s *SeverityLoggerSpy) Debug(message string) {
	s.Log(Debug, message)
}

func (s *SeverityLoggerSpy) Info(message string) {
	s.Log(Info, message)
}

func (s *SeverityLoggerSpy) Warn(message string) {
	s.Log(Warn, message)
}

func (s *SeverityLoggerSpy) Error(message string) {
	s.Log(Error, message)
}

func (s *SeverityLoggerSpy) Log(severity Severity, message string) {
	s.Logs.WriteString(fmt.Sprintf("%s: %s\n", string(severity), message))
}
