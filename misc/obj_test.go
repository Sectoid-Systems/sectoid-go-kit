package misc

import (
	"testing"
)

func TestAnyNil(t *testing.T) {
	tests := []struct {
		name     string
		pointers []interface{}
		expected bool
	}{
		{"No nil pointers", []interface{}{1, "string", 3.14}, false},
		{"One nil pointer", []interface{}{nil, "string", 3.14}, true},
		{"All nil pointers", []interface{}{nil, nil, nil}, true},
		{"Mixed pointers with nil", []interface{}{1, nil, "string"}, true},
		{"Empty input", []interface{}{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AnyNil(tt.pointers...)
			if result != tt.expected {
				t.Errorf("AnyNil(%v) = %v; expected %v", tt.pointers, result, tt.expected)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	tests := []struct {
		name     string
		pointer  interface{}
		expected bool
	}{
		{"Nil interface", nil, true},
		{"Non-nil interface", 1, false},
		{"Nil pointer", (*int)(nil), true},
		{"Non-nil pointer", new(int), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNil(tt.pointer)
			if result != tt.expected {
				t.Errorf("IsNil(%v) = %v; expected %v", tt.pointer, result, tt.expected)
			}
		})
	}
}
