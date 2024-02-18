// countSegments counts the number of segments in a string, where a segment is defined as a contiguous sequence of non-space characters.
//
// Time Complexity: O(n) where n is the length of the input string 's'.
//
// Space Complexity: O(1) as the function uses constant space for variables.
//
// Parameters:
//   - s: The input string to count segments in.
//
// Returns:
//   - The count of segments in the input string.
func countSegments(s string) int {
	count := 0
	prev := byte(' ') // Initialize the previous character to a space
	for i := 0; i < len(s); i++ {
		c := s[i]
		// If the current character is not a space and the previous character was a space,
		// increment the segment count as it indicates the start of a new segment
		if c != ' ' && prev == ' ' {
			count++
		}
		prev = c // Update the previous character for the next iteration
	}
	return count // Return the total count of segments
}
