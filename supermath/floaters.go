package supermath

const (
	FloatIntShiftSize = 10000 // shift size for truncating to 4 decimal places
)

// TruncateFloat truncates a floating-point number to 4 decimal places.
func TruncateFloat(f float64) float64 {
	shifted := int64(f * FloatIntShiftSize)     // Shift the decimal point 4 places to the right and truncate
	return float64(shifted) / FloatIntShiftSize // Shift the decimal point back to its original position
}
