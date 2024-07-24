
# logmesh Package

The `logmesh` package provides interfaces and implementations for logging with various log levels and methods.

## Types

### Provider

The `Provider` type represents a logging provider.

```go
type Provider string
```

#### Constants

- **ZapLogger**: Represents the zap logging provider.

### Config

The Config interface defines methods for retrieving logger configuration.

```go
type Config interface {
    GetProvider() Provider
    GetLogLevel() string
    IsSugaredLog() bool
}
```

### Logger

The `Logger` interface defines methods for a logger with various log levels and logging methods.

```go
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
```

### LogLevel

The `LogLevel` type represents a log level as a string.

```go
type LogLevel string
```

#### Constants

- **LogLevelDebug**: "DEBUG"
- **LogLevelDPanic**: "DPANIC"
- **LogLevelError**: "ERROR"
- **LogLevelFatal**: "FATAL"
- **LogLevelInfo**: "INFO"
- **LogLevelPanic**: "PANIC"
- **LogLevelWarn**: "WARN"
- **InvalidLogLevel**: ""

## Functions

### NewLogger

The `NewLogger` function creates a new logger instance based on the provided level, sugared flag, and provider. It returns the created Logger and an error, if any.

```go
func NewLogger(level string, sugared bool, provider Provider) (Logger, error)
```

#### Parameters

- **level string**: The log level for the logger.
- **sugared bool**: Flag indicating if the logger should be sugared.
- **provider Provider**: The logging provider.

#### Returns

- **Logger**: The created logger instance.
- **error**: An error if the logger creation fails.

### IsValidLogLevel

The `IsValidLogLevel` function checks if the provided log level is valid.

```go
func IsValidLogLevel(ll string) bool
```

#### Parameters

- **ll string**: The log level to check.

#### Returns

- **bool**: Returns true if the log level is valid, otherwise false.

### ParseLogLevel

The `ParseLogLevel` function parses the log level string, returning a valid uppercased log level or InvalidLogLevel.

```go
func ParseLogLevel(ll string) LogLevel
```

#### Parameters

- **ll string**: The log level string to parse.

#### Returns

- **LogLevel**: The parsed log level or InvalidLogLevel if invalid.

## Internal Functions

### parseZapLogLevel

The `parseZapLogLevel` function converts a LogLevel to a zapcore.Level.

```go
func parseZapLogLevel(level LogLevel) zapcore.Level
```

#### Parameters

- **level LogLevel**: The log level to convert.

#### Returns

- **zapcore.Level**: The corresponding zapcore level.

### newZapLogger

The `newZapLogger` function creates a new zapLogger instance.

```go
func newZapLogger(level LogLevel, sugared bool) (Logger, error)
```

#### Parameters

- **level LogLevel**: The log level for the logger.
- **sugared bool**: Flag indicating if the logger should be sugared.

#### Returns

- **Logger**: The created zapLogger instance.
- **error**: An error if the logger creation fails.

## zapLogger Methods

### Close

The `Close` method syncs the logger, flushing any buffered log entries.

```go
func (z *zapLogger) Close() error
```

#### Returns

- **error**: An error if the sync fails.

### Flush

The `Flush` method is a no-op for zap.SugaredLogger.

```go
func (z *zapLogger) Flush()
```

### With

The `With` method adds a key-value pair to the logger context.

```go
func (z *zapLogger) With(key, value string) Logger
```

#### Parameters

- **key string**: The key for the context.
- **value string**: The value for the context.

#### Returns

- **Logger**: The logger with the added context.

### Info

The `Info` method logs a message at Info level.

```go
func (z *zapLogger) Info(args ...interface{})
```

### Infof

The `Infof` method logs a formatted message at Info level.

```go
func (z *zapLogger) Infof(format string, v ...interface{})
```

### Debug

The `Debug` method logs a message at Debug level.

```go
func (z *zapLogger) Debug(args ...interface{})
```

### Debugf

The `Debugf` method logs a formatted message at Debug level.

```go
func (z *zapLogger) Debugf(format string, v ...interface{})
```

### Warn

The `Warn` method logs a message at Warn level.

```go
func (z *zapLogger) Warn(args ...interface{})
```

### Warnf

The `Warnf` method logs a formatted message at Warn level.

```go
func (z *zapLogger) Warnf(format string, v ...interface{})
```

### Error

The `Error` method logs a message at Error level.

```go
func (z *zapLogger) Error(args ...interface{})
```

### Errorf

The `Errorf` method logs a formatted message at Error level.

```go
func (z *zapLogger) Errorf(format string, v ...interface{})
```

### Panicf

The `Panicf` method logs a formatted message at Panic level and then panics.

```go
func (z *zapLogger) Panicf(format string, v ...interface{})
```

### DPanicf

The `DPanicf` method logs a formatted message at DPanic level. In development, the logger then panics.

```go
func (z *zapLogger) DPanicf(format string, v ...interface{})
```

### Child

The `Child` method creates a sub-logger with the given name.

```go
func (z *zapLogger) Child(name string) Logger
```

#### Parameters

- **name string**: The name for the sub-logger.

#### Returns

- **Logger**: The created sub-logger.

### Usage Example

```go
package main

import (
    "log"
    "github.com/Sectoid-Systems/sectoid-go-kit/logmesh"
)

func main() {
	config := logmesh.Config{
		Provider: "zap",
		LogLevel: "debug",
		Sugared:  true,
	}

	logger, err := logmesh.NewLogger(config)
    if err != nil {
        log.Fatalf("Failed to create logger: %v", err)
    }
    defer logger.Close()

    logger.Info("This is an info message")
    logger.Debugf("This is a debug message with value: %d", 42)
}
```
