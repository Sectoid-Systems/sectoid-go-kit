package strutils

import (
	"testing"
)

// TestInsensitive tests the Insensitive function.
func TestInsensitive(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "hello"},
		{" Hello ", "hello"},
		{"HELLO", "hello"},
		{"HeLLo WoRLD", "hello world"},
		{"", ""},
		{" ", ""},
	}

	for _, test := range tests {
		result := Insensitive(test.input)
		if result != test.expected {
			t.Errorf("Insensitive(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

// TestIEqual tests the IEqual function.
func TestIEqual(t *testing.T) {
	tests := []struct {
		input    string
		values   []string
		expected bool
	}{
		{"hello", []string{"Hello", "world"}, true},
		{"Hello", []string{"hello", "world"}, true},
		{"world", []string{"Hello", "world"}, true},
		{"hi", []string{"Hello", "world"}, false},
		{"", []string{"Hello", "world"}, false},
	}

	for _, test := range tests {
		result := IEqual(test.input, test.values...)
		if result != test.expected {
			t.Errorf("IEqual(%q, %v) = %v; expected %v", test.input, test.values, result, test.expected)
		}
	}
}

// TestTrimQuote tests the TrimQuote function.
func TestTrimQuote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"hello"`, "hello"},
		{`"hello`, "hello"},
		{`hello"`, "hello"},
		{`hello`, "hello"},
		{`"he"llo"`, `he"llo`},
		{`""`, ""},
		{`"`, ""},
		{``, ""},
	}

	for _, test := range tests {
		result := TrimQuote(test.input)
		if result != test.expected {
			t.Errorf("TrimQuote(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

// TestGenerateRandomString tests the GenerateRandomString function.
func TestGenerateRandomString(t *testing.T) {
	tests := []struct {
		length   int
		expected int
	}{
		{10, 10},
		{20, 20},
		{0, 0},
		{1, 1},
	}

	for _, test := range tests {
		result, err := GenerateRandomString(test.length)
		if err != nil {
			t.Errorf("GenerateRandomString(%d) returned an error: %v", test.length, err)
		}
		if len(result) != test.expected {
			t.Errorf("GenerateRandomString(%d) = %q; expected length %d", test.length, result, test.expected)
		}
	}
}
