package logger

import (
	"context"
)

type Logger interface {
	// Trace logs a message with the LEVEL_TRACE level.
	Trace(ctx context.Context, msg string)

	// TraceWithData logs a message with the LEVEL_TRACE level
	// and the given data.
	TraceWithData(ctx context.Context, msg string, data map[string]any)

	// Debug logs a message with the LEVEL_DEBUG level.
	Debug(ctx context.Context, msg string)

	// DebugWithData logs a message with the LEVEL_DEBUG level
	// and the given data.
	DebugWithData(ctx context.Context, msg string, data map[string]any)

	// Info logs a message with the LEVEL_INFO level.
	Info(ctx context.Context, msg string)

	// InfoWithData logs a message with the LEVEL_INFO level
	// and the given data.
	InfoWithData(ctx context.Context, msg string, data map[string]any)

	// Warn logs a message with the LEVEL_WARN level.
	Warn(ctx context.Context, msg string, err error)

	// WarnWithData logs a message with the LEVEL_WARN level
	// and the given data.
	WarnWithData(ctx context.Context, msg string, err error, data map[string]any)

	// Error logs a message with the LEVEL_ERROR level.
	Error(ctx context.Context, msg string, err error)

	// ErrorWithData logs a message with the LEVEL_ERROR level
	// and the given data.
	ErrorWithData(ctx context.Context, msg string, err error, data map[string]any)

	// Fatal logs a message with the LEVEL_FATAL level and it's followed by a
	// call to os.Exit(1).
	Fatal(ctx context.Context, msg string, err error)

	// FatalWithData logs a message with the LEVEL_FATAL level
	// and the given data and it's followed by a call to os.Exit(1).
	FatalWithData(ctx context.Context, msg string, err error, data map[string]any)
}
