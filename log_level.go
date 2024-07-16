package logger

type LogLevel int

const (
	LOG_LEVEL_NOTSET LogLevel = iota * 10 // 0
	LOG_LEVEL_DEBUG           = iota * 10 // 10
	LOG_LEVEL_INFO            = iota * 10 // 20
	LOG_LEVEL_WARN            = iota * 10 // 30
	LOG_LEVEL_ERROR           = iota * 10 // 40
	LOG_LEVEL_FATAL           = iota * 10 // 50
)

// toString returns the string representation of an int
// log level.
func (lvl LogLevel) ToString() string {
	switch level := lvl; {
	case level >= LOG_LEVEL_FATAL:
		return "fatal"
	case level >= LOG_LEVEL_ERROR:
		return "error"
	case level >= LOG_LEVEL_WARN:
		return "warn"
	case level >= LOG_LEVEL_INFO:
		return "info"
	case level >= LOG_LEVEL_DEBUG:
		return "debug"
	default:
		return "not_set"
	}
}
