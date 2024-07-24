package logmesh

import (
	"bytes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

// newTestZapLogger is a helper function to create a zapLogger instance for testing.
func newTestZapLogger(level LogLevel, sugared bool) (*zapLogger, *bytes.Buffer, error) {
	var buf bytes.Buffer
	writer := zapcore.AddSync(&buf)

	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config.EncoderConfig),
		writer,
		parseZapLogLevel(level),
	)

	logger := zap.New(core, zap.AddCaller())
	if sugared {
		return &zapLogger{logger: logger.Sugar()}, &buf, nil
	}
	return &zapLogger{logger: logger.Sugar()}, &buf, nil
}

func TestZapLogger_Info(t *testing.T) {
	logger, buf, err := newTestZapLogger(LogLevelInfo, true)
	if err != nil {
		t.Fatalf("failed to create logger: %v", err)
	}
	logger.Info("test info message")

	if !bytes.Contains(buf.Bytes(), []byte("INFO")) || !bytes.Contains(buf.Bytes(), []byte("test info message")) {
		t.Errorf("expected log message to contain 'INFO' and 'test info message', got %s", buf.String())
	}
}

func TestZapLogger_Debug(t *testing.T) {
	logger, buf, err := newTestZapLogger(LogLevelDebug, true)
	if err != nil {
		t.Fatalf("failed to create logger: %v", err)
	}
	logger.Debug("test debug message")

	if !bytes.Contains(buf.Bytes(), []byte("DEBUG")) || !bytes.Contains(buf.Bytes(), []byte("test debug message")) {
		t.Errorf("expected log message to contain 'DEBUG' and 'test debug message', got %s", buf.String())
	}
}

func TestZapLogger_Warn(t *testing.T) {
	logger, buf, err := newTestZapLogger(LogLevelWarn, true)
	if err != nil {
		t.Fatalf("failed to create logger: %v", err)
	}
	logger.Warn("test warn message")

	if !bytes.Contains(buf.Bytes(), []byte("WARN")) || !bytes.Contains(buf.Bytes(), []byte("test warn message")) {
		t.Errorf("expected log message to contain 'WARN' and 'test warn message', got %s", buf.String())
	}
}

func TestZapLogger_Error(t *testing.T) {
	logger, buf, err := newTestZapLogger(LogLevelError, true)
	if err != nil {
		t.Fatalf("failed to create logger: %v", err)
	}
	logger.Error("test error message")

	if !bytes.Contains(buf.Bytes(), []byte("ERROR")) || !bytes.Contains(buf.Bytes(), []byte("test error message")) {
		t.Errorf("expected log message to contain 'ERROR' and 'test error message', got %s", buf.String())
	}
}

func TestZapLogger_With(t *testing.T) {
	logger, buf, err := newTestZapLogger(LogLevelInfo, true)
	if err != nil {
		t.Fatalf("failed to create logger: %v", err)
	}
	logger = logger.With("key", "value").(*zapLogger)
	logger.Info("test with message")

	if !bytes.Contains(buf.Bytes(), []byte("key")) || !bytes.Contains(buf.Bytes(), []byte("value")) {
		t.Errorf("expected log message to contain 'key' and 'value', got %s", buf.String())
	}
}

func TestZapLogger_Child(t *testing.T) {
	logger, buf, err := newTestZapLogger(LogLevelInfo, true)
	if err != nil {
		t.Fatalf("failed to create logger: %v", err)
	}
	childLogger := logger.Child("childLogger").(*zapLogger)
	childLogger.Info("test child message")

	if !bytes.Contains(buf.Bytes(), []byte("ref")) || !bytes.Contains(buf.Bytes(), []byte("childlogger")) {
		t.Errorf("expected log message to contain 'ref' and 'childlogger', got %s", buf.String())
	}
}
