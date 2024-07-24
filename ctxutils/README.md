
# ctxutils Package

The `ctxutils` package provides utilities for managing context, closing resources, and handling graceful shutdowns in a safe and error-handling manner.

## Types

### Closeable

The `Closeable` interface defines methods for resources that can be closed with error handling.

```go
type Closeable interface {
    IsClosed() bool
    Close() error
}
```

#### Methods

- **IsClosed() bool**: Checks if the resource is already closed. Returns `true` if closed, otherwise `false`.
- **Close() error**: Closes the resource and returns an error if any issues occur during the closing process.

### ShutdownFunc

The `ShutdownFunc` type defines the function signature for shutdown procedures.

```go
type ShutdownFunc func() error
```

## Functions

### CloseResource

The `CloseResource` function attempts to close a `Closeable` resource if it is not already closed.

```go
func CloseResource(c Closeable) error
```

#### Parameters

- **c Closeable**: The resource that implements the `Closeable` interface.

#### Returns

- **error**: Returns `nil` if the resource is already closed or closed successfully. If an error occurs during closing, it returns an error wrapped with a descriptive message.

#### Example

```go
err := ctxutils.CloseResource(myResource)
if err != nil {
    log.Fatalf("Failed to close resource: %v", err)
}
```

### WaitForShutdown

The `WaitForShutdown` function waits for an interrupt or context cancellation signal to perform a graceful shutdown.

```go
func WaitForShutdown(ctx context.Context, shutdown ShutdownFunc, cancel context.CancelFunc)
```

#### Parameters

- **ctx context.Context**: The context to wait for cancellation or interrupt signal.
- **shutdown ShutdownFunc**: The function to execute for shutdown procedures.
- **cancel context.CancelFunc**: The function to cancel the context.

#### Example

```go
ctx, cancel := context.WithCancel(context.Background())
shutdownFunc := func() error {
    // Perform your shutdown procedures here.
    return nil
}

go ctxutils.WaitForShutdown(ctx, shutdownFunc, cancel)

// Simulate some work
time.Sleep(10 * time.Second)

// To trigger shutdown from code
cancel()
```

### Usage Example

```go
package main

import (
    "context"
    "log"
    "time"
    "github.com/Sectoid-Systems/sectoid-go-kit/ctxutils"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    shutdownFunc := func() error {
        // Perform your shutdown procedures here.
        return nil
    }

    go ctxutils.WaitForShutdown(ctx, shutdownFunc, cancel)

    // Simulate some work
    log.Println("Working...")
    time.Sleep(10 * time.Second)

    // To trigger shutdown from code
    cancel()
}
```
