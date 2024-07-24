package logmesh

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name      string
		level     string
		sugared   bool
		provider  Provider
		expectErr bool
	}{
		{"Valid ZapLogger", "info", true, ZapLogger, false},
		{"Unsupported provider", "info", true, Provider("unsupported"), true},
		{"Invalid log level", "invalid", true, ZapLogger, false}, // Assuming default log level is set
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := NewLogger(tt.level, tt.sugared, tt.provider)
			if (err != nil) != tt.expectErr {
				t.Errorf("NewLogger() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !tt.expectErr && logger == nil {
				t.Error("Expected a logger instance, got nil")
			}
		})
	}
}
