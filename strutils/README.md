
# strutils Package

The `strutils` package provides utility functions for string conversions and manipulations.

## Functions

### AsToBool

The `AsToBool` function converts a string to a boolean value. It returns true if the string (case-insensitive) is "true", otherwise false.

```go
func AsToBool(s string) bool
```

#### Parameters

- **s string**: The string to convert to a boolean.

#### Returns

- **bool**: Returns true if the string is "true" (case-insensitive), otherwise false.

### AsFloat64

The `AsFloat64` function converts a string to a float64 value. It replaces commas with dots if necessary and then parses the float value. Returns the parsed float64 value and an error if the conversion fails.

```go
func AsFloat64(s string) (float64, error)
```

#### Parameters

- **s string**: The string to convert to a float64.

#### Returns

- **float64**: The parsed float64 value.
- **error**: An error if the conversion fails.

### AsTimeRFC3339

The `AsTimeRFC3339` function parses a string in RFC3339 format and returns the corresponding time.Time value. If the parsing fails, it returns the current time and an error.

```go
func AsTimeRFC3339(s string) (time.Time, error)
```

#### Parameters

- **s string**: The string in RFC3339 format to parse.

#### Returns

- **time.Time**: The parsed time.Time value.
- **error**: An error if the parsing fails.

### AsTimestamp

The `AsTimestamp` function converts a string in RFC3339 format to a *timestamppb.Timestamp. It returns the parsed timestamp and an error if the conversion fails.

```go
func AsTimestamp(s string) (*timestamppb.Timestamp, error)
```

#### Parameters

- **s string**: The string in RFC3339 format to convert.

#### Returns

- **timestamppb.Timestamp**: The parsed timestamp.
- **error**: An error if the conversion fails.

### IsNumber

The `IsNumber` function checks if a given string represents a number. It returns true if the string consists solely of numeric digits, otherwise false.

```go
func IsNumber(s string) bool
```

#### Parameters

- **s string**: The string to check.

#### Returns

- **bool**: Returns true if the string consists solely of numeric digits, otherwise false.

### TLen

The `TLen` function returns the length of the string after trimming leading and trailing spaces.

```go
func TLen(s string) int
```

#### Parameters

- **s string**: The string to trim and measure.

#### Returns

- **int**: The length of the trimmed string.

### TLenPtr

The `TLenPtr` function returns the length of the string pointed to by the input pointer after trimming spaces. If the pointer is nil, it returns 0.

```go
func TLenPtr(s *string) int
```

#### Parameters

- **s *string**: The pointer to the string to trim and measure.

#### Returns

- **int**: The length of the trimmed string or 0 if the pointer is nil.

### IsEmpty

The `IsEmpty` function checks if a string is empty after trimming leading and trailing spaces. It returns true if the trimmed string has length 0, otherwise false.

```go
func IsEmpty(s string) bool
```

#### Parameters

- **s string**: The string to check.

#### Returns

- **bool**: Returns true if the trimmed string has length 0, otherwise false.

### IsEmptyPtr

The `IsEmptyPtr` function checks if the string pointed to by the input pointer is empty after trimming spaces. If the pointer is nil or the trimmed string has length 0, it returns true.

```go
func IsEmptyPtr(s *string) bool
```

#### Parameters

- **s *string**: The pointer to the string to check.

#### Returns

- **bool**: Returns true if the pointer is nil or the trimmed string has length 0, otherwise false.

### AnyEmpty

The `AnyEmpty` function checks if any of the provided strings are empty after trimming spaces. It returns true if at least one of the trimmed strings has length 0, otherwise false.

```go
func AnyEmpty(s ...string) bool
```

#### Parameters

- **s ...string**: The strings to check.

#### Returns

- **bool**: Returns true if at least one of the trimmed strings has length 0, otherwise false.

### Insensitive

The `Insensitive` function returns the lowercase version of the input string with leading and trailing spaces removed.

```go
func Insensitive(s string) string
```

#### Parameters

- **s string**: The string to convert.

#### Returns

- **string**: The lowercase version of the input string with spaces removed.

### IEqual

The `IEqual` function checks if the input string is equal (case-insensitive) to any of the provided values. It returns true if the input string matches any of the values after converting to lowercase and trimming spaces.

```go
func IEqual(s string, values ...string) bool
```

#### Parameters

- **s string**: The string to compare.
- **values ...string**: The values to compare against.

#### Returns

- **bool**: Returns true if the input string matches any of the provided values, otherwise false.

### TrimQuote

The `TrimQuote` function removes leading and trailing double quotes from the input string. If the string does not start or end with a double quote, it is returned unchanged.

```go
func TrimQuote(s string) string
```

#### Parameters

- **s string**: The string to trim.

#### Returns

- **string**: The trimmed string.

### GenerateRandomString

The `GenerateRandomString` function generates a random string of the specified length using alphanumeric characters. It returns the generated string and an error if the random number generation fails.

```go
func GenerateRandomString(length int) (string, error)
```

#### Parameters

- **length int**: The length of the random string to generate.

#### Returns

- **string**: The generated random string.
- **error**: An error if the random number generation fails.

### Usage Example

```go
package main

import (
    "fmt"
    "time"
    "github.com/Sectoid-Systems/sectoid-go-kit/strutils"
)

func main() {
    // Example of AsToBool
    fmt.Println("AsToBool:", strutils.AsToBool("true")) // true

    // Example of AsFloat64
    f, err := strutils.AsFloat64("123,45")
    if err != nil {
        fmt.Println("Error converting string to float64:", err)
    } else {
        fmt.Println("AsFloat64:", f)
    }

    // Example of AsTimeRFC3339
    t, err := strutils.AsTimeRFC3339("2023-01-01T12:00:00Z")
    if err != nil {
        fmt.Println("Error parsing time:", err)
    } else {
        fmt.Println("AsTimeRFC3339:", t)
    }

    // Example of GenerateRandomString
    randomStr, err := strutils.GenerateRandomString(10)
    if err != nil {
        fmt.Println("Error generating random string:", err)
    } else {
        fmt.Println("GenerateRandomString:", randomStr)
    }
}
```
