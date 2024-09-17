package ctxutils

import (
	"errors"
	"testing"
)

type logMock struct {
	Called bool
}

func (l *logMock) Logf(format string, args ...interface{}) {
	l.Called = true
}

// mockCloseable is a mock implementation of the Closeable interface for testing purposes.
type mockCloseable struct {
	closed   bool
	closeErr error
}

func (m *mockCloseable) IsClosed() bool {
	return m.closed
}

func (m *mockCloseable) Close() error {
	if m.closeErr != nil {
		return m.closeErr
	}
	m.closed = true
	return nil
}

func TestCloseResource(t *testing.T) {
	tests := []struct {
		name        string
		mock        *mockCloseable
		lm          *logMock
		expectedErr bool
	}{
		{
			name: "Resource already closed",
			mock: &mockCloseable{
				closed: true,
			},
			lm:          &logMock{},
			expectedErr: false,
		},
		{
			name: "Resource closed successfully",
			mock: &mockCloseable{
				closed: false,
			},
			lm:          &logMock{},
			expectedErr: false,
		},
		{
			name: "Error closing resource",
			mock: &mockCloseable{
				closed:   false,
				closeErr: errors.New("close error"),
			},
			lm:          &logMock{},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CloseAsNeeded(tt.mock, tt.lm.Logf)
			if !tt.expectedErr && !tt.mock.closed {
				t.Errorf("CloseResource() did not close the resource")
			}
			if tt.expectedErr && !tt.lm.Called {
				t.Errorf("CloseResource() did not log the error")
			}
		})
	}
}
