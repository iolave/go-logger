package logger

import (
	"encoding/json"
	"fmt"
)

// Field `Info` is a valid JSON type according to golang [JSON unmarsall]
// definition.
//
// [JSON Unmarshall]: https://pkg.go.dev/encoding/json#Unmarshal
type logEntry struct {
	Level         LogLevel               `json:"level"`
	Name          string                 `json:"name"`
	Msg           string                 `json:"msg"`
	Time          string                 `json:"time"`
	Pid           int                    `json:"pid"`
	Hostname      string                 `json:"hostname"`
	SchemaVersion string                 `json:"schemaVersion"`
	Info          map[string]interface{} `json:"info"`
}

// TODO: Add the proper log level using LOG_LEVEL env
func (entry logEntry) print() {
	data, _ := json.Marshal(entry)
	fmt.Println(string(data))
}
