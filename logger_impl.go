package logger

import (
	"context"
	"fmt"
	"os"

	"github.com/iolave/go-errors"
)

// Config is the configuration of the logger.
type Config struct {
	// Level is the level of the logger.
	Level Level
}

// defaultLogger is the implementation of the Logger interface.
type defaultLogger struct {
	// level is the level of the logger.
	level Level
	// name is the name of the app
	name string
	// version is the version of the app.
	version string
}

var _ Logger = &defaultLogger{}

// New creates a new logger.
//
// If cfg.Level is not valid, an error of type
// [github.com/iolave/go-errors.GenericError] is returned.
func New(level Level, name string, version string) (Logger, error) {
	if !level.IsValid() {
		return nil, errors.NewWithName(
			"logger_error",
			"invalid level",
		)
	}

	return &defaultLogger{
		level:   level,
		name:    name,
		version: version,
	}, nil
}

// Trace logs a debug entry.
func (l defaultLogger) Trace(ctx context.Context, msg string) {
	e := newEntry(ctx, l.name, l.version, LEVEL_TRACE, msg, map[string]any{}, nil)

	l.log(e)
}

// TraceWithData logs a debug entry.
func (l defaultLogger) TraceWithData(ctx context.Context, msg string, info map[string]any) {
	e := newEntry(ctx, l.name, l.version, LEVEL_TRACE, msg, info, nil)
	l.log(e)
}

// Debug logs a debug entry.
func (l defaultLogger) Debug(ctx context.Context, msg string) {
	e := newEntry(ctx, l.name, l.version, LEVEL_DEBUG, msg, map[string]any{}, nil)
	l.log(e)
}

// DebugWithData logs a debug entry.
func (l defaultLogger) DebugWithData(ctx context.Context, msg string, info map[string]any) {
	e := newEntry(ctx, l.name, l.version, LEVEL_DEBUG, msg, info, nil)
	l.log(e)
}

// Info logs an info entry.
func (l defaultLogger) Info(ctx context.Context, msg string) {
	e := newEntry(ctx, l.name, l.version, LEVEL_INFO, msg, map[string]any{}, nil)
	l.log(e)
}

// InfoWithData logs an info entry.
func (l defaultLogger) InfoWithData(ctx context.Context, msg string, info map[string]any) {
	e := newEntry(ctx, l.name, l.version, LEVEL_INFO, msg, info, nil)
	l.log(e)
}

// Warn logs a warn entry.
func (l defaultLogger) Warn(ctx context.Context, msg string, err error) {
	e := newEntry(ctx, l.name, l.version, LEVEL_WARN, msg, map[string]any{}, err)
	l.log(e)
}

// WarnWithData logs a warn entry.
func (l defaultLogger) WarnWithData(ctx context.Context, msg string, err error, info map[string]any) {
	e := newEntry(ctx, l.name, l.version, LEVEL_WARN, msg, info, err)
	l.log(e)
}

// Error logs an error entry.
func (l defaultLogger) Error(ctx context.Context, msg string, err error) {
	e := newEntry(ctx, l.name, l.version, LEVEL_ERROR, msg, map[string]any{}, nil)
	e.Error = err
	l.log(e)
}

// ErrorWithData logs an error entry.
func (l defaultLogger) ErrorWithData(ctx context.Context, msg string, err error, info map[string]any) {
	e := newEntry(ctx, l.name, l.version, LEVEL_ERROR, msg, info, nil)
	l.log(e)
}

// Fatal logs a fatal entry and it's followed by a os.Exit(1) call
func (l defaultLogger) Fatal(ctx context.Context, msg string, err error) {
	e := newEntry(ctx, l.name, l.version, LEVEL_FATAL, msg, map[string]any{}, err)
	l.log(e)
	os.Exit(1)
}

// FatalWithData logs a fatal entry and it's followed by a os.Exit(1) call
func (l defaultLogger) FatalWithData(ctx context.Context, msg string, err error, info map[string]any) {
	e := newEntry(ctx, l.name, l.version, LEVEL_FATAL, msg, info, err)
	l.log(e)
	os.Exit(1)
}

// log logs an entry
func (l defaultLogger) log(entry Entry) {
	// checks if the entry should be printed
	if entry.Level < l.level {
		return
	}

	fmt.Println(entry.serialize())
}
