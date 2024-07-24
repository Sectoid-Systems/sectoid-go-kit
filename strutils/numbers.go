package strutils

import "unicode"

// IsNumber checks if a given string represents a number.
// It returns true if the string consists solely of numeric digits, otherwise false.
func IsNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}
