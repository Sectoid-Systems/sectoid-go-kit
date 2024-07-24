package strutils

import (
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TestAsToBool tests the AsToBool function.
func TestAsToBool(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"True", true},
		{"TRUE", true},
		{"false", false},
		{"False", false},
		{"FALSE", false},
		{"other", false},
	}

	for _, test := range tests {
		result := AsToBool(test.input)
		if result != test.expected {
			t.Errorf("AsToBool(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

// TestAsFloat64 tests the AsFloat64 function.
func TestAsFloat64(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		err      bool
	}{
		{"123.45", 123.45, false},
		{"123,45", 123.45, false},
		{"123", 123, false},
		{"abc", 0, true},
	}

	for _, test := range tests {
		result, err := AsFloat64(test.input)
		if (err != nil) != test.err {
			t.Errorf("AsFloat64(%q) error = %v; expected error = %v", test.input, err, test.err)
		}
		if !test.err && result != test.expected {
			t.Errorf("AsFloat64(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

// TestAsTimeRFC3339 tests the AsTimeRFC3339 function.
func TestAsTimeRFC3339(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Time
		err      bool
	}{
		{"2021-01-01T00:00:00Z", time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), false},
		{"invalid", time.Now(), true},
	}

	for _, test := range tests {
		result, err := AsTimeRFC3339(test.input)
		if (err != nil) != test.err {
			t.Errorf("AsTimeRFC3339(%q) error = %v; expected error = %v", test.input, err, test.err)
		}
		if !test.err && !result.Equal(test.expected) {
			t.Errorf("AsTimeRFC3339(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

// TestAsTimestamp tests the AsTimestamp function.
func TestAsTimestamp(t *testing.T) {
	tests := []struct {
		input    string
		expected *timestamppb.Timestamp
		err      bool
	}{
		{"2021-01-01T00:00:00Z", timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)), false},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := AsTimestamp(test.input)
		if (err != nil) != test.err {
			t.Errorf("AsTimestamp(%q) error = %v; expected error = %v", test.input, err, test.err)
		}
		if !test.err && !result.AsTime().Equal(test.expected.AsTime()) {
			t.Errorf("AsTimestamp(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
