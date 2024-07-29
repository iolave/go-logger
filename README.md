# JSON based logger for golang

## Schema version: v1.0.0

### Log entry schema

| Field         | Description                      |
|---------------|----------------------------------|
| level         | Log entry level                  |
| name          | Name of the app/logger           |
| msg           | Log message in snake case format |
| time          | Unix time                        |
| pid           | Process id                       |
| hostname      | System's hostname                |
| schemaVersion | This schema version (v1.0.0)     |
| customData    | Custom data in any form or shape. Feel free to use this field as you want |

### Log level definition
| Level | Name  |
|-------|-------|
| 50    | fatal |
| 40    | error |
| 30    | warn  |
| 20    | info  |
| 10    | debug |
