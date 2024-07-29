package logger

import (
	"os"
	"time"

	strutils "github.com/iolave/go-logger/pkg/str_utils"
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
	entry := log.buildLogEntry(LOG_LEVEL_INFO, msg, customData)
	entry.print()
}

func (log Logger) Warn(msg string, customData map[string]interface{}) {
	entry := log.buildLogEntry(LOG_LEVEL_WARN, msg, customData)
	entry.print()
}

func (log Logger) Error(msg string, customData map[string]interface{}) {
	entry := log.buildLogEntry(LOG_LEVEL_ERROR, msg, customData)
	entry.print()
}

func (log Logger) Debug(msg string, customData map[string]interface{}) {
	entry := log.buildLogEntry(LOG_LEVEL_DEBUG, msg, customData)
	entry.print()
}

func (log Logger) Fatal(msg string, customData map[string]interface{}) {
	entry := log.buildLogEntry(LOG_LEVEL_FATAL, msg, customData)
	entry.print()
	os.Exit(1)
}

func (log Logger) buildLogEntry(level LogLevel, msg string, customData map[string]interface{}) LogEntry {
	entry := new(LogEntry)

	hostname, _ := os.Hostname()

	// Setting base log entry fields
	entry.Name = log.name
	entry.Level = level
	entry.Time = int(time.Now().Unix())
	entry.Pid = os.Getpid()
	entry.Hostname = hostname
	entry.SchemaVersion = "v1.0.0" // TODO: Add schema definition in README.md

	// Setting user log info
	entry.Msg = strutils.ToSnakeCase(msg)
	entry.CustomData = customData

	return *entry
}
