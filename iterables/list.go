package iterables

// ExistsIn checks if an element 'o' exists in the given list 'list'.
// It returns true if 'o' is found in the list, otherwise false.
func ExistsIn[T comparable](o T, list []T) bool {
	// Iterate through each element in the list
	for _, l := range list {
		// If the current element 'l' is equal to 'o', return true
		if o == l {
			return true
		}
	}

	// If the loop completes without finding 'o', return false
	return false
}
