package logger

import (
	"fmt"
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

func (log Logger) Info(msg string, info map[string]interface{}) {
	entry := log.buildLogEntry(LOG_LEVEL_INFO, msg, info)
	entry.print()
}

func (log Logger) Warn(msg string, info map[string]interface{}) {
	entry := log.buildLogEntry(LOG_LEVEL_WARN, msg, info)
	entry.print()
}

func (log Logger) Error(msg string, info map[string]interface{}) {
	entry := log.buildLogEntry(LOG_LEVEL_ERROR, msg, info)
	entry.print()
}

func (log Logger) Debug(msg string, info map[string]interface{}) {
	entry := log.buildLogEntry(LOG_LEVEL_DEBUG, msg, info)
	entry.print()
}

// TODO: add exit 1
func (log Logger) Fatal(msg string, info map[string]interface{}) {
	entry := log.buildLogEntry(LOG_LEVEL_FATAL, msg, info)
	entry.print()
}

func (log Logger) buildLogEntry(level LogLevel, msg string, info map[string]interface{}) logEntry {
	entry := new(logEntry)

	hostname, _ := os.Hostname()

	// Setting base log entry fields
	entry.Name = log.name
	entry.Level = level.ToString()
	entry.Time = fmt.Sprintf("%d", time.Now().Unix())
	entry.Pid = os.Getpid()
	entry.Hostname = hostname
	entry.SchemaVersion = "v1.0.0" // TODO: Add schema definition in README.md

	// Setting user log info
	entry.Msg = strutils.ToSnakeCase(msg)
	entry.Info = info

	return *entry
}
