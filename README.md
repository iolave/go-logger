# JSON based logger for golang

A simple and efficient JSON-based logging library for Go applications, designed for easy integration and observability. It provides structured logging with various levels, automatic trace context propagation, and customizable output to streamline log management and analysis.

## Install
```bash
go get github.com/iolave/go-logger
```

## Features
- JSON formatted logs for easy parsing and querying.
- Log levels to control log verbosity.
- Automatic inclusion of trace information from context via [go-trace](https://github.com/iolave/go-trace).
- Handles circular references in log data.
- Customizable log output with additional data.

## Usage

### Create a logger
```golang
// ...
import "github.com/iolave/go-logger"

// Creates a logger with a name and version
logger, err := logger.New(logger.LEVEL_INFO, "my-awesome-logger", "1.0.0")
if err != nil {
    // Handle error
}
```

### Write logs
```golang
ctx := context.Background() // Or your request context

logger.Debug(ctx, "debug message")
// {"timestamp":1722216216,"level":"debug","name":"my-awesome-logger","version":"1.0.0","trace":{},"info":{},"msg":"debug message"}

logger.Info(ctx, "info message")
// {"timestamp":1722216216,"level":"info","name":"my-awesome-logger","version":"1.0.0","trace":{},"info":{},"msg":"info message"}

logger.Warn(ctx, "warn message", errors.New("something happened"))
// {"timestamp":1722216216,"level":"warn","name":"my-awesome-logger","version":"1.0.0","trace":{},"error":{},"info":{},"msg":"warn message"}

logger.Error(ctx, "error message", errors.New("something bad happened"))
// {"timestamp":1722216216,"level":"error","name":"my-awesome-logger","version":"1.0.0","trace":{},"error":{},"info":{},"msg":"error message"}

logger.Fatal(ctx, "fatal message", errors.New("something terrible happened"))
// Exits the application after logging
// {"timestamp":1722216216,"level":"fatal","name":"my-awesome-logger","version":"1.0.0","trace":{},"error":{},"info":{},"msg":"fatal message"}
```

### With trace information
The logger will automatically include trace information from the context if it's provided via [go-trace](https://github.com/iolave/go-trace).

```golang
import (
	"context"
	"errors"

	"github.com/iolave/go-logger"
	"github.com/iolave/go-trace"
)

// ...

// Create a new trace
t := trace.Trace{}
t.Set("trace_id", "my-trace-id")
t.Set("span_id", "my-span-id")

// Add the trace to the context
ctx := t.SetInContext(context.Background())

// Log a message with the context
logger.Info(ctx, "This log includes trace information")
// {"timestamp":1722216216,"level":"info","name":"my-awesome-logger","version":"1.0.0","trace":{"trace_id":"my-trace-id","span_id":"my-span-id"},"info":{},"msg":"This log includes trace information"}
```

## Schema version: v1.0.0

### Log entry schema

| Field         | Description                      | JSON type |
|---------------|----------------------------------|-----------|
| timestamp     | Unix time                        | `number`    |
| level         | Log entry level                  | `string`    | 
| name          | Name of the app/logger           | `string`    |
| version       | Version of the app/logger        | `string`    |
| trace         | Trace information from context   | `object`    |
| error         | Error information                | `object`    |
| info          | Custom data in any form or shape | `object`    |
| msg           | Log message                      | `string`    |

### Log level definition
| Level | Name  |
|-------|-------|
| 60    | fatal |
| 50    | error |
| 40    | warn  |
| 30    | info  |
| 20    | debug |
| 10    | trace |

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

### Development

To get started with development, you'll need Go installed.

1. Fork the repository.
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/go-logger.git`
3. Create a new branch: `git checkout -b my-feature-branch`
4. Make your changes.
5. Run tests: `make test`
6. Push to your branch: `git push origin my-feature-branch`
7. Open a pull request.
