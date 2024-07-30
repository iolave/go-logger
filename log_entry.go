package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	strutils "github.com/iolave/go-logger/pkg/str_utils"
)

// Field `Info` is a valid JSON type according to golang [JSON unmarsall]
// definition.
//
// [JSON Unmarshall]: https://pkg.go.dev/encoding/json#Unmarshal
type LogEntry struct {
	Level         LogLevel               `json:"level"`
	Name          string                 `json:"name"`
	Msg           string                 `json:"msg"`
	Time          int                    `json:"time"`
	Pid           int                    `json:"pid"`
	Hostname      string                 `json:"hostname"`
	SchemaVersion string                 `json:"schemaVersion"`
	CustomData    map[string]interface{} `json:"customData"`
}

// TODO: Add the proper log level using LOG_LEVEL env
func (entry LogEntry) print() {
	if !entry.shouldPrint() {
		return
	}

	data, _ := json.Marshal(entry)
	fmt.Println(string(data))
}

func (entry LogEntry) shouldPrint() bool {
	envLogLevel := getLogLevelFromEnv()
	if entry.Level >= envLogLevel {
		return true
	}

	return false
}

func newLogEntry(name string, level LogLevel, msg string, customData map[string]interface{}) LogEntry {
	entry := new(LogEntry)

	hostname, _ := os.Hostname()

	// Setting base log entry fields
	entry.Name = name
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
