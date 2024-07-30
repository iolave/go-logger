package logger

import (
	"os"
)

type Logger struct {
	name string
}

func New(name string) *Logger {
	return &Logger{
		name: name,
	}
}

func (log Logger) Info(msg string, customData map[string]interface{}) {
	entry := newLogEntry(log.name, LOG_LEVEL_INFO, msg, customData)
	entry.print()
}

func (log Logger) Warn(msg string, customData map[string]interface{}) {
	entry := newLogEntry(log.name, LOG_LEVEL_WARN, msg, customData)
	entry.print()
}

func (log Logger) Error(msg string, customData map[string]interface{}) {
	entry := newLogEntry(log.name, LOG_LEVEL_ERROR, msg, customData)
	entry.print()
}

func (log Logger) Debug(msg string, customData map[string]interface{}) {
	entry := newLogEntry(log.name, LOG_LEVEL_DEBUG, msg, customData)
	entry.print()
}

func (log Logger) Fatal(msg string, customData map[string]interface{}) {
	entry := newLogEntry(log.name, LOG_LEVEL_FATAL, msg, customData)
	entry.print()
	os.Exit(1)
}
