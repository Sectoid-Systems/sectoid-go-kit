package timespace

import (
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// FromJSTimestamp converts a JavaScript timestamp string (in milliseconds) to a time.Time object.
// It parses the string to an int64 and converts it to time.Time.
// If parsing fails, it returns the current time and an error.
func FromJSTimestamp(s string) (time.Time, error) {
	ts, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Now(), err
	}
	return FromJSTSToTime(ts), nil
}

// FromJSTSToTime converts a JavaScript timestamp (in milliseconds) represented as an int64 to a time.Time object.
// It divides the timestamp by 1000 to get seconds and calculates nanoseconds for the remainder.
func FromJSTSToTime(ts int64) time.Time {
	return time.Unix(ts/1000, (ts%1000)*1e6)
}

// FromUnixToTimestamp converts a Unix timestamp (in seconds) represented as an int64 to a *timestamppb.Timestamp.
// If the timestamp is zero, it returns nil.
func FromUnixToTimestamp(ts int64) *timestamppb.Timestamp {
	if ts == 0 {
		return nil // Return nil if the timestamp is zero
	}
	return timestamppb.New(time.Unix(ts, 0))
}
