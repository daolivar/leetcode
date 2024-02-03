// lengthOfLastWord calculates the length of the last word in a given string 's'.
//
// Time Complexity: O(n) where n is the length of the input string 's'.
//
// Space Complexity: O(1) as the function uses constant space for variables.
//
// Parameters:
//   - s: The input string in which the length of the last word needs to be calculated.
//
// Returns:
//   - The length of the last word in the string. If there is no last word, the function returns 0.
func lengthOfLastWord(s string) int {
	lastWordLength := 0

	// Iterate through the string from the end
	for i := len(s) - 1; i >= 0; i-- {
		// Skip trailing spaces before the last word
		if s[i] == ' ' && lastWordLength == 0 {
			continue
		} else if s[i] == ' ' {
			// Break when encountering a space after the last word
			break
		} else {
			// Count the characters of the last word
			lastWordLength++
		}
	}

	// Return the length of the last word
	return lastWordLength
}
