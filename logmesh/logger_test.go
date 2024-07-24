package logmesh

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockConfig struct {
	provider   Provider
	logLevel   string
	sugaredLog bool
}

func (m *MockConfig) GetProvider() Provider {
	return m.provider
}

func (m *MockConfig) GetLogLevel() string {
	return m.logLevel
}

func (m *MockConfig) IsSugaredLog() bool {
	return m.sugaredLog
}

func TestNewLogger_ZapLogger(t *testing.T) {
	mockConfig := &MockConfig{
		provider:   ZapLogger,
		logLevel:   "info",
		sugaredLog: true,
	}

	logger, err := NewLogger(mockConfig)

	assert.NoError(t, err)
	assert.NotNil(t, logger)
}

func TestNewLogger_UnsupportedProvider(t *testing.T) {
	mockConfig := &MockConfig{
		provider:   "unsupported",
		logLevel:   "info",
		sugaredLog: false,
	}

	logger, err := NewLogger(mockConfig)

	assert.Error(t, err)
	assert.Nil(t, logger)
	assert.Equal(t, errors.New("unsupported logger provider: unsupported"), err)
}

func TestNewLogger_ValidProviderWithoutSugar(t *testing.T) {
	mockConfig := &MockConfig{
		provider:   ZapLogger,
		logLevel:   "debug",
		sugaredLog: false,
	}

	logger, err := NewLogger(mockConfig)

	assert.NoError(t, err)
	assert.NotNil(t, logger)
}
