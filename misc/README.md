
# misc Package

The `misc` package provides utility functions for various common tasks such as checking for nil pointers and retrying operations with timeouts.

## Functions

### AnyNil

The `AnyNil` function checks if any of the provided pointers are nil. It returns true if at least one of the pointers is nil, otherwise false.

```go
func AnyNil(pointers ...interface{}) bool
```

#### Parameters

- **pointers ...interface{}**: A variadic parameter representing the pointers to be checked.

#### Returns

- **bool**: Returns true if any pointer is nil, otherwise false.

### IsNil

The `IsNil` function checks if the given pointer is nil. It handles both untyped nil and typed nil (such as nil pointers).

```go
func IsNil(p interface{}) bool
```

#### Parameters

- **p interface{}**: An interface representing the pointer to be checked.

#### Returns

- **bool**: Returns true if the pointer is nil, otherwise false.

### Retry

The `Retry` function runs the provided function a specified number of times, passing a context with each timeout as a parameter. It waits at least until the timeout before moving on to the next attempt. Returns the result of the last call.

```go
func Retry(ctx context.Context, f func(ctx context.Context) error, timeouts []time.Duration) error
```

#### Parameters

- **ctx context.Context**: The context for the retry operation.
- **f func(ctx context.Context) error**: The function to retry.
- **timeouts []time.Duration**: A slice of timeouts for each retry attempt.

#### Returns

- **error**: Returns nil if the function succeeds within the retry attempts, otherwise returns the error from the last attempt.

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "time"
    "github.com/Sectoid-Systems/sectoid-go-kit/misc"
)

func main() {
    // Example of AnyNil and IsNil
    var ptr1, ptr2 *int
    var ptr3 = new(int)
    fmt.Println("AnyNil:", misc.AnyNil(ptr1, ptr2, ptr3)) // true
    fmt.Println("IsNil (ptr1):", misc.IsNil(ptr1)) // true
    fmt.Println("IsNil (ptr3):", misc.IsNil(ptr3)) // false

    // Example of Retry
    ctx := context.Background()
    timeouts := []time.Duration{1 * time.Second, 2 * time.Second, 3 * time.Second}
    err := misc.Retry(ctx, func(ctx context.Context) error {
        // Simulate a task that may fail
        return errors.New("task failed")
    }, timeouts)

    if err != nil {
        fmt.Println("Retry failed:", err)
    }
}
```
