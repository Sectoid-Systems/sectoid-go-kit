package logmesh

import "fmt"

// Provider represents a logging provider.
type Provider string

const (
	// ZapLogger represents the zap logging provider.
	ZapLogger Provider = "zap"
)

// Logger defines the interface for a logger with various log levels and methods.
type Logger interface {
	Info(args ...interface{})
	Infof(format string, v ...interface{})
	Debug(args ...interface{})
	Debugf(format string, v ...interface{})
	Warn(args ...interface{})
	Warnf(format string, v ...interface{})
	Error(args ...interface{})
	Errorf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
	DPanicf(format string, v ...interface{})
	With(key string, value string) Logger
	Child(name string) Logger
	Flush()
	Close() error
}

// NewLogger creates a new logger instance based on the provided level, sugared flag, and provider.
// It returns the created Logger and an error, if any.
func NewLogger(level string, sugared bool, provider Provider) (Logger, error) {
	ll := ParseLogLevel(level)
	switch provider {
	case ZapLogger:
		return newZapLogger(ll, sugared)
	default:
		return nil, fmt.Errorf("unsupported logger provider: %s", provider)
	}
}
