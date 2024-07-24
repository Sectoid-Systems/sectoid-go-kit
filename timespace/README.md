
# timespace Package

The `timespace` package provides utility functions for converting between various timestamp formats and time representations.

## Functions

### FromJSTimestamp

The `FromJSTimestamp` function converts a JavaScript timestamp string (in milliseconds) to a `time.Time` object. It parses the string to an `int64` and converts it to `time.Time`. If parsing fails, it returns the current time and an error.

```go
func FromJSTimestamp(s string) (time.Time, error)
```

#### Parameters

- **s string**: The JavaScript timestamp string (in milliseconds) to convert.

#### Returns

- **time.Time**: The converted `time.Time` object.
- **error**: An error if the parsing fails.

### FromJSTSToTime

The `FromJSTSToTime` function converts a JavaScript timestamp (in milliseconds) represented as an `int64` to a `time.Time` object. It divides the timestamp by 1000 to get seconds and calculates nanoseconds for the remainder.

```go
func FromJSTSToTime(ts int64) time.Time
```

#### Parameters

- **ts int64**: The JavaScript timestamp (in milliseconds) to convert.

#### Returns

- **time.Time**: The converted `time.Time` object.

### FromUnixToTimestamp

The `FromUnixToTimestamp` function converts a Unix timestamp (in seconds) represented as an `int64` to a `*timestamppb.Timestamp`. If the timestamp is zero, it returns nil.

```go
func FromUnixToTimestamp(ts int64) *timestamppb.Timestamp
```

#### Parameters

- **ts int64**: The Unix timestamp (in seconds) to convert.

#### Returns

- **timestamppb.Timestamp**: The converted `*timestamppb.Timestamp`, or nil if the timestamp is zero.

### Usage Example

```go
package main

import (
    "fmt"
    "github.com/Sectoid-Systems/sectoid-go-kit/timespace"
)

func main() {
    // Example of FromJSTimestamp
    jsTimestamp := "1625097600000"
    t, err := timespace.FromJSTimestamp(jsTimestamp)
    if err != nil {
        fmt.Println("Error converting JS timestamp:", err)
    } else {
        fmt.Println("Converted time:", t)
    }

    // Example of FromUnixToTimestamp
    unixTimestamp := int64(1625097600)
    ts := timespace.FromUnixToTimestamp(unixTimestamp)
    if ts != nil {
        fmt.Println("Converted timestamp:", ts)
    } else {
        fmt.Println("Timestamp is zero or invalid")
    }
}
```
