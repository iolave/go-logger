# JSON based logger for golang

> [!WARNING]
> All versions released prior to `v1.0.0` are to be considered [breaking changes](https://semver.org/#how-do-i-know-when-to-release-100) (I'll try my best to not push breaking changes btw).

## Install
```bash
go get github.com/iolave/go-logger
```

## Environment variables

| Variable    | Description | Default value |
|-------------|-------------|---------------|
| `LOG_LEVEL` | Accepts any name value of the [log level definition](#log-level-definition). | `info` |

## Usage

### Create a logger
```golang
// ...
import "github.com/iolave/go-logger"

// Creates a logger with an awesome name
logger := logger.New("my-awesome-logger")
```

### Write logs
```golang
logger.Debug("debug message", map[string]any{})
// {"level":10,"name":"my-awesome-logger","msg":"debug_message","time":1722216216,"pid":97541,"hostname":"ignacios-mm.home.iolave.com","schemaVersion":"v1.0.0","customData":{}}
logger.Info("info message", map[string]any{})
// {"level":20,"name":"my-awesome-logger","msg":"info_message","time":1722216216,"pid":97541,"hostname":"ignacios-mm.home.iolave.com","schemaVersion":"v1.0.0","customData":{}}
logger.Warn("warn message", map[string]any{})
// {"level":30,"name":"my-awesome-logger","msg":"warn_message","time":1722216216,"pid":97541,"hostname":"ignacios-mm.home.iolave.com","schemaVersion":"v1.0.0","customData":{}}
logger.Error("error message", map[string]any{})
// {"level":40,"name":"my-awesome-logger","msg":"error_message","time":1722216216,"pid":97541,"hostname":"ignacios-mm.home.iolave.com","schemaVersion":"v1.0.0","customData":{}}
logger.Fatal("fatal message", map[string]any{})
// {"level":50,"name":"my-awesome-logger","msg":"fatal_message","time":1722216216,"pid":97541,"hostname":"ignacios-mm.home.iolave.com","schemaVersion":"v1.0.0","customData":{}}
```

## Schema version: v1.0.0

### Log entry schema

| Field         | Description                      | JSON type |
|---------------|----------------------------------|-----------|
| level         | Log entry level                  | `number`    | 
| name          | Name of the app/logger           | `string`    |
| msg           | Log message in snake case format | `string`    |
| time          | Unix time                        | `number`    |
| pid           | Process id                       | `number`    |
| hostname      | System's hostname                | `string`    |
| schemaVersion | This schema version (v1.0.0)     | `string`    |
| customData    | Custom data in any form or shape. Feel free to use this field as you want | `Record<string, any>` |

### Log level definition
| Level | Name  |
|-------|-------|
| 50    | fatal |
| 40    | error |
| 30    | warn  |
| 20    | info  |
| 10    | debug |
