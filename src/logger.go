package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	strhelpers "github.com/iolave/go-logger/helpers/string"
)

type Logger struct {
	name string
}

func New(name string) Logger {
	logger := new(Logger)

	logger.name = name

	return *logger
}

func (log Logger) Info(msg string) {
	entry := log.buildLogEntryBase()
	entry.Msg = strhelpers.ToSnakeCase(msg)
	entry.print()
}

func (log Logger) Warn(msg string) {
	entry := log.buildLogEntryBase()
	entry.Msg = strhelpers.ToSnakeCase(msg)
	entry.print()
}

func (log Logger) Error(msg string) {
	entry := log.buildLogEntryBase()
	entry.Msg = strhelpers.ToSnakeCase(msg)
	entry.print()
}

func (log Logger) Debug(msg string) {
	entry := log.buildLogEntryBase()
	entry.Msg = strhelpers.ToSnakeCase(msg)
	entry.print()
}

// TODO: add exit 1
func (log Logger) Fatal(msg string) {
	entry := log.buildLogEntryBase()
	entry.Msg = strhelpers.ToSnakeCase(msg)
	entry.print()
}

type Trace struct {
	MachId         string `json:"mach_id"`
	DeviceId       string `json:"device_id"`
	RequestId      string `json:"request_id"`
	BusinessMachId string `json:"business_mach_id"`
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
	Trace         Trace  `json:"trace"`
	Info          any    `json:"info"` // TODO: Write it properly
	Action        action `json:"action"`
}

func (log Logger) buildLogEntryBase() logEntry {
	entry := new(logEntry)

	hostname, _ := os.Hostname()

	entry.Duration = -1 // TODO: Add the proper duration
	entry.Hostname = hostname
	entry.Level = 0 // TODO: Add the proper log level using LOG_LEVEL env
	entry.Name = log.name
	entry.Pid = os.Getpid()
	entry.SchemaVersion = "v1.0.0"   // TODO: maybe remove this
	entry.Time = time.Now().String() // TODO: Format it properly

	return *entry
}

func (entry logEntry) print() {
	data, _ := json.Marshal(entry)
	fmt.Println(string(data))
}
