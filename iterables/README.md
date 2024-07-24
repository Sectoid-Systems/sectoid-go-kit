
# iterables Package

The `iterables` package provides utility functions for working with lists and maps.

## Functions

### ExistsIn

The `ExistsIn` function checks if an element 'o' exists in the given list 'list'. It returns true if 'o' is found in the list, otherwise false.

```go
func ExistsIn[T comparable](o T, list []T) bool
```

#### Parameters

- **o T**: The element to check for existence in the list.
- **list []T**: The list in which to check for the element.

#### Returns

- **bool**: Returns true if the element exists in the list, otherwise false.

### MapHasKeys

The `MapHasKeys` function returns true if the map contains ALL keys passed.

```go
func MapHasKeys[K comparable, V any](m map[K]V, keys ...K) bool
```

#### Parameters

- **m map[K]V**: The map to check for the keys.
- **keys ...K**: The keys to check for in the map.

#### Returns

- **bool**: Returns true if the map contains all the keys passed, otherwise false.

### MapHasString

The `MapHasString` function returns TRUE when a string value is not empty.

```go
func MapHasString[K comparable](m map[K]string, key K) bool
```

#### Parameters

- **m map[K]string**: The map to check for the key.
- **key K**: The key to check for in the map.

#### Returns

- **bool**: Returns true if the string value of the key is not empty, otherwise false.

### MapHasStrings

The `MapHasStrings` function returns true if string values of ALL KEYS passed are not empty.

```go
func MapHasStrings[K comparable](m map[K]string, keys ...K) bool
```

#### Parameters

- **m map[K]string**: The map to check for the keys.
- **keys ...K**: The keys to check for in the map.

#### Returns

- **bool**: Returns true if the string values of all the keys are not empty, otherwise false.

### MapHasStringValues

The `MapHasStringValues` function returns true if all string values in the map are not empty.

```go
func MapHasStringValues[K comparable](m map[K]string) bool
```

#### Parameters

- **m map[K]string**: The map to check for the string values.

#### Returns

- **bool**: Returns true if all string values in the map are not empty, otherwise false.

### MapKeys

The `MapKeys` function returns an array containing all keys from a map.

```go
func MapKeys[K comparable, V any](m map[K]V) []K
```

#### Parameters

- **m map[K]V**: The map to retrieve the keys from.

#### Returns

- **[]K**: An array containing all keys from the map.

### InsensitiveMap

The `InsensitiveMap` function converts a map with string keys to a map with insensitive string keys using `strutils.Insensitive`.

```go
func InsensitiveMap[V any](m map[string]V) map[string]V
```

#### Parameters

- **m map[string]V**: The map with string keys to convert.

#### Returns

- **map[string]V**: The converted map with insensitive string keys.

### MapValuesFromKeys

The `MapValuesFromKeys` function returns all values of keys from a given map.

```go
func MapValuesFromKeys[K comparable, V any](m map[K]V, keys ...K) []V
```

#### Parameters

- **m map[K]V**: The map to retrieve the values from.
- **keys ...K**: The keys to retrieve the values for.

#### Returns

- **[]V**: An array containing all values of the specified keys from the map.

### Usage Example

```go
package main

import (
    "fmt"
    "github.com/Sectoid-Systems/sectoid-go-kit/iterables"
)

func main() {
    // Example of ExistsIn
    list := []int{1, 2, 3, 4, 5}
    exists := iterables.ExistsIn(3, list)
    fmt.Println("Exists in list:", exists)

    // Example of MapHasKeys
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    hasKeys := iterables.MapHasKeys(m, "a", "b")
    fmt.Println("Map has keys:", hasKeys)

    // Example of MapHasString
    strMap := map[string]string{"key1": "value1", "key2": ""}
    hasString := iterables.MapHasString(strMap, "key1")
    fmt.Println("Map has non-empty string:", hasString)
}
```
