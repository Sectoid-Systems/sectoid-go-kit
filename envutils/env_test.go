package envutils

import (
	"os"
	"testing"
	"time"
)

// TestGetEnvWithDefault tests the GetEnvWithDefault function.
func TestGetEnvWithDefault(t *testing.T) {
	os.Setenv("TEST_KEY", "value")
	defer os.Unsetenv("TEST_KEY")

	tests := []struct {
		key        string
		defaultVal string
		expected   string
	}{
		{"TEST_KEY", "default", "value"},
		{"MISSING_KEY", "default", "default"},
	}

	for _, test := range tests {
		result := GetEnvWithDefault(test.key, test.defaultVal)
		if result != test.expected {
			t.Errorf("GetEnvWithDefault(%q, %q) = %q; expected %q", test.key, test.defaultVal, result, test.expected)
		}
	}
}

// TestGetEnv tests the GetEnv function.
func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "value")
	defer os.Unsetenv("TEST_KEY")

	tests := []struct {
		key      string
		expected string
	}{
		{"TEST_KEY", "value"},
		{"MISSING_KEY", ""},
	}

	for _, test := range tests {
		result := GetEnv(test.key)
		if result != test.expected {
			t.Errorf("GetEnv(%q) = %q; expected %q", test.key, result, test.expected)
		}
	}
}

// TestGetEnvAsBool tests the GetEnvAsBool function.
func TestGetEnvAsBool(t *testing.T) {
	os.Setenv("BOOL_TRUE", "true")
	os.Setenv("BOOL_FALSE", "false")
	defer os.Unsetenv("BOOL_TRUE")
	defer os.Unsetenv("BOOL_FALSE")

	tests := []struct {
		key        string
		defaultVal bool
		expected   bool
	}{
		{"BOOL_TRUE", false, true},
		{"BOOL_FALSE", true, false},
		{"MISSING_BOOL", true, true},
		{"INVALID_BOOL", true, true},
	}

	for _, test := range tests {
		result := GetEnvAsBool(test.key, test.defaultVal)
		if result != test.expected {
			t.Errorf("GetEnvAsBool(%q, %v) = %v; expected %v", test.key, test.defaultVal, result, test.expected)
		}
	}
}

// TestGetEnvAsInt tests the GetEnvAsInt function.
func TestGetEnvAsInt(t *testing.T) {
	os.Setenv("INT_KEY", "123")
	defer os.Unsetenv("INT_KEY")

	tests := []struct {
		key        string
		defaultVal int
		expected   int
	}{
		{"INT_KEY", 0, 123},
		{"MISSING_INT", 456, 456},
		{"INVALID_INT", 789, 789},
	}

	for _, test := range tests {
		result := GetEnvAsInt(test.key, test.defaultVal)
		if result != test.expected {
			t.Errorf("GetEnvAsInt(%q, %d) = %d; expected %d", test.key, test.defaultVal, result, test.expected)
		}
	}
}

// TestGetEnvAsInt32 tests the GetEnvAsInt32 function.
func TestGetEnvAsInt32(t *testing.T) {
	os.Setenv("INT32_KEY", "123")
	defer os.Unsetenv("INT32_KEY")

	tests := []struct {
		key        string
		defaultVal int32
		expected   int32
	}{
		{"INT32_KEY", 0, 123},
		{"MISSING_INT32", 456, 456},
		{"INVALID_INT32", 789, 789},
	}

	for _, test := range tests {
		result := GetEnvAsInt32(test.key, test.defaultVal)
		if result != test.expected {
			t.Errorf("GetEnvAsInt32(%q, %d) = %d; expected %d", test.key, test.defaultVal, result, test.expected)
		}
	}
}

// TestGetEnvAsDuration tests the GetEnvAsDuration function.
func TestGetEnvAsDuration(t *testing.T) {
	os.Setenv("DURATION_KEY", "1h30m")
	defer os.Unsetenv("DURATION_KEY")

	tests := []struct {
		key        string
		defaultVal string
		expected   time.Duration
		err        bool
	}{
		{"DURATION_KEY", "0s", 90 * time.Minute, false},
		{"MISSING_DURATION", "1h", 1 * time.Hour, false},
		{"INVALID_DURATION", "invalid", 0, true},
	}

	for _, test := range tests {
		result, err := GetEnvAsDuration(test.key, test.defaultVal)
		if (err != nil) != test.err {
			t.Errorf("GetEnvAsDuration(%q, %q) error = %v; expected error = %v", test.key, test.defaultVal, err, test.err)
		}
		if result != test.expected {
			t.Errorf("GetEnvAsDuration(%q, %q) = %v; expected %v", test.key, test.defaultVal, result, test.expected)
		}
	}
}
