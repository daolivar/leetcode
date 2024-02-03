// maximumValue finds the maximum value (either numeric or string length) among the elements in the given slice of strings.
//
// Time Complexity: O(k * n) where k is the maximum length of any string in 'strs' and n is the length of 'strs'.
//
// Space Complexity: O(1) as the function uses constant space for variables.
//
// Parameters:
//   - strs: The input slice of strings containing either numeric or non-numeric values.
//
// Returns:
//   - The maximum value among the elements in 'strs'.
func maximumValue(strs []string) int {
	mx := 0

	// Iterate through the slice of strings
	for _, s := range strs {
		// Attempt to convert the string to an integer
		val, err := strconv.Atoi(s)
		if err != nil {
			// If it's not a valid integer, consider the length of the string
			mx = max(mx, len(s))
		} else {
			// If it's a valid integer, consider the integer value
			mx = max(mx, val)
		}
	}

	// Return the maximum value
	return mx
}

// Max returns the maximum of two integers.
//
// Time Complexity: O(1) as the function performs a constant-time operation.
//
// Parameters:
//   - m: The first integer.
//   - n: The second integer.
//
// Returns:
//   - The maximum of the two integers.
func max(m, n int) int {
	if m < n {
		return n
	}
	return m
}
