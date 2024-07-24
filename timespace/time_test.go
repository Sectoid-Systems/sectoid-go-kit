package timespace

import (
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TestFromJSTimestamp tests the FromJSTimestamp function.
func TestFromJSTimestamp(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{"Valid timestamp", "1627938000000", false},
		{"Invalid timestamp", "invalid", true},
		{"Empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FromJSTimestamp(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("FromJSTimestamp(%v) error = %v, expectErr %v", tt.input, err, tt.expectErr)
			}

			if !tt.expectErr {
				expected := time.Unix(1627938000, 0) // 1627938000000 ms = 1627938000 s
				if !result.Equal(expected) {
					t.Errorf("FromJSTimestamp(%v) = %v, expected %v", tt.input, result, expected)
				}
			}
		})
	}
}

// TestFromJSTSToTime tests the FromJSTSToTime function.
func TestFromJSTSToTime(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected time.Time
	}{
		{"Valid timestamp", 1627938000000, time.Unix(1627938000, 0)},
		{"Timestamp with milliseconds", 1627938000123, time.Unix(1627938000, 123000000)},
		{"Zero timestamp", 0, time.Unix(0, 0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromJSTSToTime(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("FromJSTSToTime(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestFromUnixToTimestamp tests the FromUnixToTimestamp function.
func TestFromUnixToTimestamp(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected *timestamppb.Timestamp
	}{
		{"Valid timestamp", 1627938000, timestamppb.New(time.Unix(1627938000, 0))},
		{"Zero timestamp", 0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromUnixToTimestamp(tt.input)
			if result == nil && tt.expected == nil {
				return
			}
			if result == nil || tt.expected == nil || !result.AsTime().Equal(tt.expected.AsTime()) {
				t.Errorf("FromUnixToTimestamp(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
