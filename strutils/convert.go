package strutils

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"strings"
	"time"
)

// AsToBool converts a string to a boolean value.
// It returns true if the string (case-insensitive) is "true", otherwise false.
func AsToBool(s string) bool {
	return strings.ToLower(s) == "true"
}

// AsFloat64 converts a string to a float64 value.
// It replaces commas with dots if necessary and then parses the float value.
// Returns the parsed float64 value and an error if the conversion fails.
func AsFloat64(s string) (float64, error) {
	n := s
	if strings.Contains(s, ",") {
		n = strings.Replace(s, ",", ".", -1)
	}

	return strconv.ParseFloat(n, 64)
}

// AsTimeRFC3339 parses a string in RFC3339 format and returns the corresponding time.Time value.
// If the parsing fails, it returns the current time and an error.
func AsTimeRFC3339(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Now(), err
	}

	return t, nil
}

// AsTimestamp converts a string in RFC3339 format to a *timestamppb.Timestamp.
// It returns the parsed timestamp and an error if the conversion fails.
func AsTimestamp(s string) (*timestamppb.Timestamp, error) {
	ts, err := AsTimeRFC3339(s)
	if err != nil {
		return nil, err
	}

	return timestamppb.New(ts), nil
}
