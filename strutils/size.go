package strutils

import "strings"

// TLen returns the length of the string after trimming leading and trailing spaces.
func TLen(s string) int {
	return len(strings.TrimSpace(s))
}

// TLenPtr returns the length of the string pointed to by the input pointer after trimming spaces.
// If the pointer is nil, it returns 0.
func TLenPtr(s *string) int {
	if s == nil {
		return 0
	}
	return TLen(*s)
}

// IsEmpty checks if a string is empty after trimming leading and trailing spaces.
// It returns true if the trimmed string has length 0, otherwise false.
func IsEmpty(s string) bool {
	return TLen(s) == 0
}

// IsEmptyPtr checks if the string pointed to by the input pointer is empty after trimming spaces.
// If the pointer is nil or the trimmed string has length 0, it returns true.
func IsEmptyPtr(s *string) bool {
	return TLenPtr(s) == 0
}

// AnyEmpty checks if any of the provided strings are empty after trimming spaces.
// It returns true if at least one of the trimmed strings has length 0, otherwise false.
func AnyEmpty(s ...string) bool {
	for _, v := range s {
		if IsEmpty(v) {
			return true
		}
	}
	return false
}
