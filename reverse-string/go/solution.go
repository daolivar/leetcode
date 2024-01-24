// reverseString reverses a slice of bytes in place.
//
// Time Complexity: O(n) where n is the length of the input slice 's'.
//                  The function iterates through half of the slice.
//
// Space Complexity: O(1) as the reversal is performed in place without using additional space.
//
// Parameters:
//   - s: The input slice of bytes to reverse.
//
// Returns:
//   - None (modifies the input slice 's' in place).
func reverseString(s []byte) {
	// Iterate through half of the slice
	for i := 0; i < len(s)/2; i++ {
		// Swap elements symmetrically around the center of the slice
		temp := s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = temp
	}
}
