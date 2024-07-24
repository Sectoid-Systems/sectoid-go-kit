package strutils

import "testing"

// TestTLen tests the TLen function.
func TestTLen(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{" hello ", 5},
		{"", 0},
		{" ", 0},
		{"  ", 0},
		{"\t\nhello\n\t", 5},
	}

	for _, test := range tests {
		result := TLen(test.input)
		if result != test.expected {
			t.Errorf("TLen(%q) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

// TestTLenPtr tests the TLenPtr function.
func TestTLenPtr(t *testing.T) {
	tests := []struct {
		input    *string
		expected int
	}{
		{ptr("hello"), 5},
		{ptr(" hello "), 5},
		{ptr(""), 0},
		{ptr(" "), 0},
		{ptr("  "), 0},
		{ptr("\t\nhello\n\t"), 5},
		{nil, 0},
	}

	for _, test := range tests {
		result := TLenPtr(test.input)
		if result != test.expected {
			t.Errorf("TLenPtr(%v) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

// TestIsEmpty tests the IsEmpty function.
func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"hello", false},
		{" hello ", false},
		{"", true},
		{" ", true},
		{"  ", true},
		{"\t\nhello\n\t", false},
	}

	for _, test := range tests {
		result := IsEmpty(test.input)
		if result != test.expected {
			t.Errorf("IsEmpty(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

// TestIsEmptyPtr tests the IsEmptyPtr function.
func TestIsEmptyPtr(t *testing.T) {
	tests := []struct {
		input    *string
		expected bool
	}{
		{ptr("hello"), false},
		{ptr(" hello "), false},
		{ptr(""), true},
		{ptr(" "), true},
		{ptr("  "), true},
		{ptr("\t\nhello\n\t"), false},
		{nil, true},
	}

	for _, test := range tests {
		result := IsEmptyPtr(test.input)
		if result != test.expected {
			t.Errorf("IsEmptyPtr(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

// TestAnyEmpty tests the AnyEmpty function.
func TestAnyEmpty(t *testing.T) {
	tests := []struct {
		input    []string
		expected bool
	}{
		{[]string{"hello", "world"}, false},
		{[]string{" hello ", " world "}, false},
		{[]string{"", "world"}, true},
		{[]string{"hello", ""}, true},
		{[]string{"", ""}, true},
		{[]string{" ", "world"}, true},
		{[]string{"hello", " "}, true},
		{[]string{"\t\nhello\n\t", "\t\nworld\n\t"}, false},
	}

	for _, test := range tests {
		result := AnyEmpty(test.input...)
		if result != test.expected {
			t.Errorf("AnyEmpty(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

// Helper function to create a pointer to a string.
func ptr(s string) *string {
	return &s
}
