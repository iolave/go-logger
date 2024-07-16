package logger

import (
	"context"
	"encoding/json"
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

func (log Logger) Info(msg string) {
	entry := log.buildLogEntryBase()
	entry.Msg = strutils.ToSnakeCase(msg)
	entry.print()
}

func (log Logger) InfoWithContext(ctx context.Context, msg string) {
	fmt.Println(ctx.Value("test"))
	entry := log.buildLogEntryBase()
	entry.Msg = strutils.ToSnakeCase(msg)
	entry.print()
}

func (log Logger) Warn(msg string) {
	entry := log.buildLogEntryBase()
	entry.Msg = strutils.ToSnakeCase(msg)
	entry.print()
}

func (log Logger) Error(msg string) {
	entry := log.buildLogEntryBase()
	entry.Msg = strutils.ToSnakeCase(msg)
	entry.print()
}

func (log Logger) Debug(msg string) {
	entry := log.buildLogEntryBase()
	entry.Msg = strutils.ToSnakeCase(msg)
	entry.print()
}

// TODO: add exit 1
func (log Logger) Fatal(msg string) {
	entry := log.buildLogEntryBase()
	entry.Msg = strutils.ToSnakeCase(msg)
	entry.print()
}

func (log Logger) buildLogEntryBase() logEntry {
	entry := new(logEntry)

	hostname, _ := os.Hostname()

	entry.Name = log.name
	entry.Level = 0                  // TODO: Add the proper log level using LOG_LEVEL env
	entry.Time = time.Now().String() // TODO: Format it properly
	entry.Pid = os.Getpid()
	entry.Hostname = hostname
	entry.SchemaVersion = "v1.0.0" // TODO: maybe remove this
	entry.Duration = -1            // TODO: Add the proper duration
	return *entry
}

type action struct {
	Source string `json:"source"`
	Args   any    `json:"args"` // TODO: Write it properly
}

type logEntry struct {
	Name          string `json:"name"`
	Level         int    `json:"level"`
	Msg           string `json:"msg"`
	Time          string `json:"time"`
	Pid           int    `json:"pid"`
	Hostname      string `json:"hostname"`
	SchemaVersion string `json:"schemaVersion"`
	Duration      int    `json:"duration"`
	Info          any    `json:"info"` // TODO: Write it properly
	Action        action `json:"action"`
}

func (entry logEntry) print() {
	data, _ := json.Marshal(entry)
	fmt.Println(string(data))
}
