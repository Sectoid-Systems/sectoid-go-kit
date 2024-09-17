package misc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCastWithDefault_NilValue checks if CastWithDefault returns the default value when the input is nil.
func TestCastWithDefault_NilValue(t *testing.T) {
	// Act
	result := CastWithDefault(nil, "default")

	// Assert
	assert.Equal(t, "default", result, "CastWithDefault should return the default value when the input is nil")
}

// TestCastWithDefault_Success checks if CastWithDefault successfully casts the value to the desired type.
func TestCastWithDefault_Success(t *testing.T) {
	// Arrange
	value := "testString"

	// Act
	result := CastWithDefault(value, "default")

	// Assert
	assert.Equal(t, "testString", result, "CastWithDefault should return the value if the cast is successful")
}

// TestCastWithDefault_FailedCast checks if CastWithDefault returns the default value when the cast fails.
func TestCastWithDefault_FailedCast(t *testing.T) {
	// Act
	result := CastWithDefault(123, "default")

	// Assert
	assert.Equal(t, "default", result, "CastWithDefault should return the default value when the cast fails")
}

// TestSilentMarshal_Success checks if SilentMarshal successfully marshals a valid object to JSON.
func TestSilentMarshal_Success(t *testing.T) {
	// Arrange
	obj := map[string]string{"key": "value"}

	// Act
	result := SilentMarshal(obj)

	// Assert
	expected := `{"key":"value"}`
	assert.JSONEq(t, expected, string(result), "SilentMarshal should successfully marshal the object to JSON")
}

// TestSilentMarshal_EmptyObject checks if SilentMarshal works with an empty object.
func TestSilentMarshal_EmptyObject(t *testing.T) {
	// Arrange
	obj := map[string]string{}

	// Act
	result := SilentMarshal(obj)

	// Assert
	expected := `{}` // Expected JSON for an empty map
	assert.JSONEq(t, expected, string(result), "SilentMarshal should marshal an empty object to empty JSON")
}
