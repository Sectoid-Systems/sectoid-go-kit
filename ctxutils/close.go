package ctxutils

import (
	"fmt"
)

// Closeable defines an interface for resources that can be closed with error handling.
type Closeable interface {
	IsClosed() bool
	Close() error
}

// CloseResource attempts to close a Closeable resource if it is not already closed.
func CloseResource(c Closeable) error {
	if c.IsClosed() {
		return nil
	}
	if err := c.Close(); err != nil {
		return fmt.Errorf("error closing resource: %w", err)
	}
	return nil
}
