
# supermath Package

The `supermath` package provides utility functions for mathematical operations, including truncating floating-point numbers to a specific number of decimal places.

## Constants

### FloatIntShiftSize

The `FloatIntShiftSize` constant defines the shift size for truncating to 4 decimal places.

```go
const FloatIntShiftSize = 10000
```

## Functions

### TruncateFloat

The `TruncateFloat` function truncates a floating-point number to 4 decimal places.

```go
func TruncateFloat(f float64) float64
```

#### Parameters

- **f float64**: The floating-point number to truncate.

#### Returns

- **float64**: The truncated floating-point number.

### Usage Example

```go
package main

import (
    "fmt"
    "github.com/Sectoid-Systems/sectoid-go-kit/supermath"
)

func main() {
    f := 123.456789
    truncated := supermath.TruncateFloat(f)
    fmt.Printf("Original: %f, Truncated: %f
", f, truncated) // Original: 123.456789, Truncated: 123.456700
}
```
