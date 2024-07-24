
# envutils Package

The `envutils` package provides utility functions for retrieving environment variables with default values and type conversions.

## Functions

### GetEnvWithDefault

The `GetEnvWithDefault` function retrieves the value of the environment variable named by the key. If the variable is present in the environment, it returns the trimmed value. Otherwise, it returns the default value.

```go
func GetEnvWithDefault(key string, defaultVal string) string
```

#### Parameters

- **key string**: The name of the environment variable.
- **defaultVal string**: The default value to return if the environment variable is not present.

#### Returns

- **string**: The value of the environment variable or the default value if the variable is not present.

### GetEnv

The `GetEnv` function retrieves the value of the environment variable named by the key. If the variable is not present in the environment, it returns an empty string.

```go
func GetEnv(key string) string
```

#### Parameters

- **key string**: The name of the environment variable.

#### Returns

- **string**: The value of the environment variable or an empty string if the variable is not present.

### GetEnvAsBool

The `GetEnvAsBool` function retrieves the value of the environment variable named by the key and parses it as a boolean. If the variable is not present or cannot be parsed as a boolean, it returns the default value.

```go
func GetEnvAsBool(key string, defaultVal bool) bool
```

#### Parameters

- **key string**: The name of the environment variable.
- **defaultVal bool**: The default value to return if the environment variable is not present or cannot be parsed as a boolean.

#### Returns

- **bool**: The parsed boolean value of the environment variable or the default value if the variable is not present or cannot be parsed.

### GetEnvAsInt

The `GetEnvAsInt` function retrieves the value of the environment variable named by the key and parses it as an integer. If the variable is not present or cannot be parsed as an integer, it returns the default value.

```go
func GetEnvAsInt(key string, defaultVal int) int
```

#### Parameters

- **key string**: The name of the environment variable.
- **defaultVal int**: The default value to return if the environment variable is not present or cannot be parsed as an integer.

#### Returns

- **int**: The parsed integer value of the environment variable or the default value if the variable is not present or cannot be parsed.

### GetEnvAsInt32

The `GetEnvAsInt32` function retrieves the value of the environment variable named by the key and parses it as an int32. If the variable is not present or cannot be parsed as an int32, it returns the default value.

```go
func GetEnvAsInt32(key string, defaultVal int32) int32
```

#### Parameters

- **key string**: The name of the environment variable.
- **defaultVal int32**: The default value to return if the environment variable is not present or cannot be parsed as an int32.

#### Returns

- **int32**: The parsed int32 value of the environment variable or the default value if the variable is not present or cannot be parsed.

### GetEnvAsDuration

The `GetEnvAsDuration` function retrieves an environment variable as a `time.Duration`. If the variable is not set or cannot be parsed as a duration, it returns the default value.

```go
func GetEnvAsDuration(key string, defaultVal string) (time.Duration, error)
```

#### Parameters

- **key string**: The name of the environment variable.
- **defaultVal string**: The default value to return if the environment variable is not present or cannot be parsed as a duration.

#### Returns

- **time.Duration**: The parsed duration value of the environment variable or the default value if the variable is not present or cannot be parsed.
- **error**: An error if the value cannot be parsed as a duration.

### Usage Example

```go
package main

import (
    "fmt"
    "github.com/Sectoid-Systems/sectoid-go-kit/envutils"
)

func main() {
    // Example of GetEnvWithDefault
    dbHost := envutils.GetEnvWithDefault("DB_HOST", "localhost")
    fmt.Println("DB Host:", dbHost)

    // Example of GetEnvAsBool
    debug := envutils.GetEnvAsBool("DEBUG", false)
    fmt.Println("Debug mode:", debug)

    // Example of GetEnvAsInt
    maxRetries := envutils.GetEnvAsInt("MAX_RETRIES", 5)
    fmt.Println("Max retries:", maxRetries)

    // Example of GetEnvAsDuration
    timeout, err := envutils.GetEnvAsDuration("TIMEOUT", "30s")
    if err != nil {
        fmt.Println("Error parsing timeout:", err)
    } else {
        fmt.Println("Timeout:", timeout)
    }
}
```
