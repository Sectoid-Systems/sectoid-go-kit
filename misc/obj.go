package misc

import (
	"reflect"
)

// Nilable Constrain T to only types that can be nil (e.g., pointers)
type Nilable interface {
	~*any // Allows only pointer types
}

// AnyNil checks if any of the provided pointers are nil.
// It returns true if at least one of the pointers is nil, otherwise false.
//
// Parameters:
// - pointers: a variadic parameter of interface{} that represents the pointers to be checked.
//
// Returns:
// - bool: true if any pointer is nil, false otherwise.
func AnyNil(pointers ...interface{}) bool {
	for _, p := range pointers {
		if IsNil(p) {
			return true
		}
	}
	return false
}

// IsNil checks if the given pointer is nil.
// It handles both untyped nil and typed nil (such as nil pointers).
//
// Parameters:
// - p: an interface{} that represents the pointer to be checked.
//
// Returns:
// - bool: true if the pointer is nil, false otherwise.
func IsNil(p interface{}) bool {
	return p == nil || (reflect.ValueOf(p).Kind() == reflect.Ptr && reflect.ValueOf(p).IsNil())
}

func AnyNotNil(pointers ...interface{}) bool {
	for _, p := range pointers {
		if IsNotNil(p) {
			return true
		}
	}
	return false
}

func IsNotNil(p interface{}) bool {
	return !IsNil(p)
}
