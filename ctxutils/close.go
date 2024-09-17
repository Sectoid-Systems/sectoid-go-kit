package ctxutils

import "github.com/Sectoid-Systems/sectoid-go-kit/misc"

// Closeable is an interface that defines a resource with a Close method.
type Closeable interface {
	Close() error
}

// CloseableWithCheck extends Closeable, adding the IsClosed method to check if the resource is already closed.
type CloseableWithCheck interface {
	Closeable
	IsClosed() bool
}

// CloseAsNeeded closes the resource if it is not already closed.
// It logs any errors that occur during the closing process.
func CloseAsNeeded(c CloseableWithCheck, lf misc.LoggerFunc) {
	if !c.IsClosed() {
		CloseResource(c, lf)
	}
}

// CloseResource attempts to close the resource and logs any errors that occur.
func CloseResource(c Closeable, lf misc.LoggerFunc) {
	if err := c.Close(); err != nil {
		lf("error closing resource: %v", err)
	}
}
