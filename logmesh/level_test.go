package logmesh

import (
	"testing"
)

func TestIsValidLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid Debug", "DEBUG", true},
		{"Valid Info", "INFO", true},
		{"Valid Warn", "WARN", true},
		{"Valid Error", "ERROR", true},
		{"Valid DPanic", "DPANIC", true},
		{"Valid Fatal", "FATAL", true},
		{"Valid Panic", "PANIC", true},
		{"Valid lowercase debug", "debug", true},
		{"Invalid empty string", "", false},
		{"Invalid level", "INVALID", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidLogLevel(tt.input)
			if result != tt.expected {
				t.Errorf("IsValidLogLevel(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected LogLevel
	}{
		{"Parse Debug", "DEBUG", LogLevelDebug},
		{"Parse Info", "INFO", LogLevelInfo},
		{"Parse Warn", "WARN", LogLevelWarn},
		{"Parse Error", "ERROR", LogLevelError},
		{"Parse DPanic", "DPANIC", LogLevelDPanic},
		{"Parse Fatal", "FATAL", LogLevelFatal},
		{"Parse Panic", "PANIC", LogLevelPanic},
		{"Parse lowercase debug", "debug", LogLevelDebug},
		{"Parse empty string", "", InvalidLogLevel},
		{"Parse invalid level", "INVALID", InvalidLogLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseLogLevel(tt.input)
			if result != tt.expected {
				t.Errorf("ParseLogLevel(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
