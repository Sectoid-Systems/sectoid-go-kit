package logmesh

import "strings"

// LogLevel represents a log level as a string.
type LogLevel string

// Define log levels as constants of type LogLevel.
const (
	LogLevelDebug   LogLevel = "DEBUG"
	LogLevelDPanic  LogLevel = "DPANIC"
	LogLevelError   LogLevel = "ERROR"
	LogLevelFatal   LogLevel = "FATAL"
	LogLevelInfo    LogLevel = "INFO"
	LogLevelPanic   LogLevel = "PANIC"
	LogLevelWarn    LogLevel = "WARN"
	InvalidLogLevel LogLevel = ""
)

// logLevelsSet is a set of valid log levels for quick lookup.
var logLevelsSet = map[LogLevel]bool{
	LogLevelDebug:  true,
	LogLevelDPanic: true,
	LogLevelError:  true,
	LogLevelFatal:  true,
	LogLevelInfo:   true,
	LogLevelPanic:  true,
	LogLevelWarn:   true,
}

// IsValidLogLevel checks if the provided log level is valid.
func IsValidLogLevel(ll string) bool {
	_, ok := logLevelsSet[LogLevel(strings.ToUpper(ll))]
	return ok
}

// ParseLogLevel parses the log level string, returning a valid uppercased log level or InvalidLogLevel.
func ParseLogLevel(ll string) LogLevel {
	upper := LogLevel(strings.ToUpper(ll))
	if IsValidLogLevel(string(upper)) {
		return upper
	}
	return InvalidLogLevel
}
