// Package envutils provides utility functions for retrieving environment variables with default values and type conversions.
package envutils

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// GetEnvWithDefault retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, it returns the trimmed value. Otherwise, it returns the default value.
func GetEnvWithDefault(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return strings.TrimSpace(value)
	}
	return defaultVal
}

// GetEnv retrieves the value of the environment variable named by the key.
// If the variable is not present in the environment, it returns an empty string.
func GetEnv(key string) string {
	return GetEnvWithDefault(key, "")
}

// GetEnvAsBool retrieves the value of the environment variable named by the key and parses it as a boolean.
// If the variable is not present or cannot be parsed as a boolean, it returns the default value.
func GetEnvAsBool(key string, defaultVal bool) bool {
	b, err := strconv.ParseBool(GetEnv(key))
	if err != nil {
		return defaultVal
	}
	return b
}

// GetEnvAsInt retrieves the value of the environment variable named by the key and parses it as an integer.
// If the variable is not present or cannot be parsed as an integer, it returns the default value.
func GetEnvAsInt(key string, defaultVal int) int {
	i, err := strconv.Atoi(GetEnv(key))
	if err != nil {
		return defaultVal
	}
	return i
}

// GetEnvAsInt32 retrieves the value of the environment variable named by the key and parses it as an int32.
// If the variable is not present or cannot be parsed as an int32, it returns the default value.
func GetEnvAsInt32(key string, defaultVal int32) int32 {
	i, err := strconv.ParseInt(GetEnv(key), 10, 32)
	if err != nil {
		return defaultVal
	}
	return int32(i)
}

// GetEnvAsDuration retrieves an environment variable as a time.Duration.
// If the variable is not set or cannot be parsed as a duration, it returns the default value.
func GetEnvAsDuration(key string, defaultVal string) (time.Duration, error) {
	// Retrieve the environment variable
	s, ok := os.LookupEnv(key)
	if !ok {
		s = defaultVal
	}

	// Parse the duration
	d, err := time.ParseDuration(s)
	if err != nil {
		return 0, err
	}

	return d, nil
}
