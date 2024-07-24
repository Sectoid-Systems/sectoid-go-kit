package iterables

import "github.com/Sectoid-Systems/sectoid-go-kit/strutils"

// MapHasKeys returns true if the map contains ALL keys passed.
func MapHasKeys[K comparable, V any](m map[K]V, keys ...K) bool {
	for _, key := range keys {
		if _, ok := m[key]; !ok {
			return false
		}
	}
	return true
}

// MapHasString returns TRUE when a string value is not empty.
func MapHasString[K comparable](m map[K]string, key K) bool {
	if s, ok := m[key]; ok {
		return !strutils.IsEmpty(s)
	}

	return false
}

// MapHasStrings returns true if string values of ALL KEYS passed are not empty.
func MapHasStrings[K comparable](m map[K]string, keys ...K) bool {
	for _, key := range keys {
		if !MapHasString(m, key) {
			return false
		}
	}

	return true
}

// MapHasStringValues returns true if all string values in the map are not empty.
func MapHasStringValues[K comparable](m map[K]string) bool {
	keys := MapKeys(m)
	return MapHasStrings(m, keys...)
}

// MapKeys returns an array containing all keys from a map.
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0

	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// InsensitiveMap converts a map with string keys to a map with insensitive string keys using strutils.Insensitive.
func InsensitiveMap[V any](m map[string]V) map[string]V {
	r := make(map[string]V)

	for k, v := range m {
		r[strutils.Insensitive(k)] = v
	}

	return r
}

// MapValuesFromKeys returns all values of keys from a given map.
func MapValuesFromKeys[K comparable, V any](m map[K]V, keys ...K) []V {
	values := make([]V, len(keys))

	for i, key := range keys {
		if value, ok := m[key]; ok {
			values[i] = value
		}
	}

	return values
}
