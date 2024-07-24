// Package strutils provides utility functions for string conversions and manipulations.
package strutils

import (
	"crypto/rand"
	"strings"
)

// Insensitive returns the lowercase version of the input string with leading and trailing spaces removed.
func Insensitive(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// IEqual checks if the input string is equal (case-insensitive) to any of the provided values.
// It returns true if the input string matches any of the values after converting to lowercase and trimming spaces.
func IEqual(s string, values ...string) bool {
	for _, v := range values {
		if Insensitive(s) == Insensitive(v) {
			return true
		}
	}
	return false
}

// TrimQuote removes leading and trailing double quotes from the input string.
// If the string does not start or end with a double quote, it is returned unchanged.
func TrimQuote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

// GenerateRandomString generates a random string of the specified length using alphanumeric characters.
// It returns the generated string and an error if the random number generation fails.
func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	for i := 0; i < length; i++ {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b), nil
}
