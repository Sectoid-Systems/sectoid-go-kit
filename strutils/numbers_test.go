package strutils

import "testing"

// TestIsNumber tests the IsNumber function.
func TestIsNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"12345", true},
		{"012345", true},
		{"", false},
		{"123a45", false},
		{" 12345", false},
		{"12345 ", false},
		{"123 45", false},
		{"-12345", false},
		{".12345", false},
	}

	for _, test := range tests {
		result := IsNumber(test.input)
		if result != test.expected {
			t.Errorf("IsNumber(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
