package iterables

import (
	"testing"
)

func TestExistsIn(t *testing.T) {
	tests := []struct {
		name     string
		element  interface{}
		list     interface{}
		expected bool
	}{
		{
			name:     "Integer exists in list",
			element:  5,
			list:     []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Integer does not exist in list",
			element:  6,
			list:     []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "String exists in list",
			element:  "apple",
			list:     []string{"banana", "orange", "apple", "grape"},
			expected: true,
		},
		{
			name:     "String does not exist in list",
			element:  "cherry",
			list:     []string{"banana", "orange", "apple", "grape"},
			expected: false,
		},
		{
			name:     "Empty list",
			element:  1,
			list:     []int{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result bool

			switch v := tt.list.(type) {
			case []int:
				result = ExistsIn(tt.element.(int), v)
			case []string:
				result = ExistsIn(tt.element.(string), v)
			}

			if result != tt.expected {
				t.Errorf("ExistsIn(%v, %v) = %v; expected %v", tt.element, tt.list, result, tt.expected)
			}
		})
	}
}
