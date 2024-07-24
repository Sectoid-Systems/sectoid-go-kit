package supermath

import (
	"testing"
)

// TestTruncateFloat tests the TruncateFloat function for various scenarios.
func TestTruncateFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"Normal truncation", 123.456789, 123.4567},
		{"No truncation needed", 123.4567, 123.4567},
		{"Less than four decimal places", 123.4, 123.4},
		{"Exactly four decimal places", 123.4567, 123.4567},
		{"More than four decimal places rounding down", 123.45674, 123.4567},
		{"More than four decimal places rounding up", 123.45675, 123.4567},
		{"Negative number truncation", -123.456789, -123.4567},
		{"Zero value", 0.0, 0.0},
		{"Small positive number", 0.000123456, 0.0001},
		{"Small negative number", -0.000123456, -0.0001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TruncateFloat(tt.input)
			if result != tt.expected {
				t.Errorf("TruncateFloat(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
