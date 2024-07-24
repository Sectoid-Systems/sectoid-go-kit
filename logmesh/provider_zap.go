package logmesh

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

// parseZapLogLevel converts a LogLevel to a zapcore.Level.
func parseZapLogLevel(level LogLevel) zapcore.Level {
	switch level {
	case LogLevelDebug:
		return zapcore.DebugLevel
	case LogLevelInfo:
		return zapcore.InfoLevel
	case LogLevelWarn:
		return zapcore.WarnLevel
	case LogLevelError:
		return zapcore.ErrorLevel
	case LogLevelDPanic:
		return zapcore.DPanicLevel
	case LogLevelPanic:
		return zapcore.PanicLevel
	case LogLevelFatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

type zapLogger struct {
	logger *zap.SugaredLogger
}

// newZapLogger creates a new zapLogger instance.
func newZapLogger(level LogLevel, sugared bool) (Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(parseZapLogLevel(level))

	if sugared {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	config.DisableStacktrace = true // TODO: work with stack options in the future

	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return &zapLogger{logger: logger.Sugar()}, nil
}

// Close syncs the logger, flushing any buffered log entries.
func (z *zapLogger) Close() error {
	return z.logger.Sync()
}

// Flush is a no-op for zap.SugaredLogger.
func (z *zapLogger) Flush() {}

// With adds a key-value pair to the logger context.
func (z *zapLogger) With(key, value string) Logger {
	return &zapLogger{logger: z.logger.With(key, value)}
}

// Info logs a message at Info level.
func (z *zapLogger) Info(args ...interface{}) {
	z.logger.Info(args...)
}

// Infof logs a formatted message at Info level.
func (z *zapLogger) Infof(format string, v ...interface{}) {
	z.logger.Infof(format, v...)
}

// Debug logs a message at Debug level.
func (z *zapLogger) Debug(args ...interface{}) {
	z.logger.Debug(args...)
}

// Debugf logs a formatted message at Debug level.
func (z *zapLogger) Debugf(format string, v ...interface{}) {
	z.logger.Debugf(format, v...)
}

// Warn logs a message at Warn level.
func (z *zapLogger) Warn(args ...interface{}) {
	z.logger.Warn(args...)
}

// Warnf logs a formatted message at Warn level.
func (z *zapLogger) Warnf(format string, v ...interface{}) {
	z.logger.Warnf(format, v...)
}

// Error logs a message at Error level.
func (z *zapLogger) Error(args ...interface{}) {
	z.logger.Error(args...)
}

// Errorf logs a formatted message at Error level.
func (z *zapLogger) Errorf(format string, v ...interface{}) {
	z.logger.Errorf(format, v...)
}

// Panicf logs a formatted message at Panic level and then panics.
func (z *zapLogger) Panicf(format string, v ...interface{}) {
	z.logger.Panicf(format, v...)
}

// DPanicf logs a formatted message at DPanic level. In development, the logger then panics.
func (z *zapLogger) DPanicf(format string, v ...interface{}) {
	z.logger.DPanicf(format, v...)
}

// Child creates a sub-logger with the given name.
func (z *zapLogger) Child(name string) Logger {
	return &zapLogger{logger: z.logger.With("ref", strings.ToLower(name))}
}
