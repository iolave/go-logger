package logger

import (
	"encoding/json"
	"fmt"
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
