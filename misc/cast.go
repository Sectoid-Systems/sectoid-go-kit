package misc

import "encoding/json"

// CastWithDefault attempts to cast a value to the specified type T.
// If the value is nil or cannot be cast, it returns the provided default value.
func CastWithDefault[T any](value interface{}, defaultValue T) T {
	if IsNil(value) {
		return defaultValue
	}

	// Attempt to cast the value to the desired type.
	parsed, ok := value.(T)
	if ok {
		return parsed
	}

	// If the cast fails, return the default value.
	return defaultValue
}

// SilentMarshal marshals an object to JSON, silently ignoring any errors.
// It returns the JSON byte slice.
func SilentMarshal[T any](obj T) []byte {
	bts, _ := json.Marshal(obj)
	return bts
}
