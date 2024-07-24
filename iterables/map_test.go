package iterables

import (
	"testing"
)

func TestMapHasKeys(t *testing.T) {
	tests := []struct {
		name     string
		m        map[int]string
		keys     []int
		expected bool
	}{
		{"All keys exist", map[int]string{1: "a", 2: "b", 3: "c"}, []int{1, 2, 3}, true},
		{"One key missing", map[int]string{1: "a", 2: "b", 3: "c"}, []int{1, 4}, false},
		{"Empty map", map[int]string{}, []int{1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapHasKeys(tt.m, tt.keys...)
			if result != tt.expected {
				t.Errorf("MapHasKeys(%v, %v) = %v; expected %v", tt.m, tt.keys, result, tt.expected)
			}
		})
	}
}

func TestMapHasString(t *testing.T) {
	tests := []struct {
		name     string
		m        map[int]string
		key      int
		expected bool
	}{
		{"Key exists and not empty", map[int]string{1: "a", 2: ""}, 1, true},
		{"Key exists but empty", map[int]string{1: "a", 2: ""}, 2, false},
		{"Key does not exist", map[int]string{1: "a", 2: ""}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapHasString(tt.m, tt.key)
			if result != tt.expected {
				t.Errorf("MapHasString(%v, %v) = %v; expected %v", tt.m, tt.key, result, tt.expected)
			}
		})
	}
}

func TestMapHasStrings(t *testing.T) {
	tests := []struct {
		name     string
		m        map[int]string
		keys     []int
		expected bool
	}{
		{"All keys have non-empty values", map[int]string{1: "a", 2: "b"}, []int{1, 2}, true},
		{"One key has empty value", map[int]string{1: "a", 2: ""}, []int{1, 2}, false},
		{"Key does not exist", map[int]string{1: "a"}, []int{1, 2}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapHasStrings(tt.m, tt.keys...)
			if result != tt.expected {
				t.Errorf("MapHasStrings(%v, %v) = %v; expected %v", tt.m, tt.keys, result, tt.expected)
			}
		})
	}
}

func TestMapHasStringValues(t *testing.T) {
	tests := []struct {
		name     string
		m        map[int]string
		expected bool
	}{
		{"All values are non-empty", map[int]string{1: "a", 2: "b"}, true},
		{"One value is empty", map[int]string{1: "a", 2: ""}, false},
		{"Empty map", map[int]string{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapHasStringValues(tt.m)
			if result != tt.expected {
				t.Errorf("MapHasStringValues(%v) = %v; expected %v", tt.m, result, tt.expected)
			}
		})
	}
}

func TestMapKeys(t *testing.T) {
	tests := []struct {
		name     string
		m        map[int]string
		expected []int
	}{
		{"Non-empty map", map[int]string{1: "a", 2: "b"}, []int{1, 2}},
		{"Empty map", map[int]string{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapKeys(tt.m)
			if len(result) != len(tt.expected) {
				t.Errorf("MapKeys(%v) = %v; expected %v", tt.m, result, tt.expected)
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("MapKeys(%v) = %v; expected %v", tt.m, result, tt.expected)
				}
			}
		})
	}
}

func TestInsensitiveMap(t *testing.T) {
	tests := []struct {
		name     string
		m        map[string]string
		expected map[string]string
	}{
		{"Mixed case keys", map[string]string{"KeyOne": "value1", "KEYTWO": "value2"}, map[string]string{"keyone": "value1", "keytwo": "value2"}},
		{"Lowercase keys", map[string]string{"keyone": "value1", "keytwo": "value2"}, map[string]string{"keyone": "value1", "keytwo": "value2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InsensitiveMap(tt.m)
			if len(result) != len(tt.expected) {
				t.Errorf("InsensitiveMap(%v) = %v; expected %v", tt.m, result, tt.expected)
			}
			for k, v := range result {
				if v != tt.expected[k] {
					t.Errorf("InsensitiveMap(%v) = %v; expected %v", tt.m, result, tt.expected)
				}
			}
		})
	}
}

func TestMapValuesFromKeys(t *testing.T) {
	tests := []struct {
		name     string
		m        map[int]string
		keys     []int
		expected []string
	}{
		{"All keys exist", map[int]string{1: "a", 2: "b"}, []int{1, 2}, []string{"a", "b"}},
		{"One key does not exist", map[int]string{1: "a", 2: "b"}, []int{1, 3}, []string{"a", ""}},
		{"Empty map", map[int]string{}, []int{1, 2}, []string{"", ""}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapValuesFromKeys(tt.m, tt.keys...)
			if len(result) != len(tt.expected) {
				t.Errorf("MapValuesFromKeys(%v, %v) = %v; expected %v", tt.m, tt.keys, result, tt.expected)
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("MapValuesFromKeys(%v, %v) = %v; expected %v", tt.m, tt.keys, result, tt.expected)
				}
			}
		})
	}
}
