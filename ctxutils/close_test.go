package ctxutils

import (
	"errors"
	"testing"
)

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
		expectedErr bool
	}{
		{
			name: "Resource already closed",
			mock: &mockCloseable{
				closed: true,
			},
			expectedErr: false,
		},
		{
			name: "Resource closed successfully",
			mock: &mockCloseable{
				closed: false,
			},
			expectedErr: false,
		},
		{
			name: "Error closing resource",
			mock: &mockCloseable{
				closed:   false,
				closeErr: errors.New("close error"),
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CloseResource(tt.mock)
			if (err != nil) != tt.expectedErr {
				t.Errorf("CloseResource() error = %v, expectedErr %v", err, tt.expectedErr)
			}
			if !tt.expectedErr && !tt.mock.closed {
				t.Errorf("CloseResource() did not close the resource")
			}
		})
	}
}
