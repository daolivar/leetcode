// truncateSentence truncates a given string 's' to contain only the first 'k' words.
//
// Time Complexity: O(n) where n is the length of the input string 's'.
//
// Space Complexity: O(1) as the function uses constant space for variables.
//
// Parameters:
//   - s: The input string to be truncated.
//   - k: The number of words to retain in the truncated string.
//
// Returns:
//   - The truncated string containing the first 'k' words.
func truncateSentence(s string, k int) string {
	// Initialize word count and result string
	wordCount := 0
	result := ""

	// Iterate through the characters of the input string
	for _, c := range s {
		// Increment word count when encountering a space character
		if c == ' ' {
			wordCount++
			// Break the loop if the desired number of words is reached
			if wordCount == k {
				break
			}
		}
		// Append the current character to the result string
		result += string(c)
	}

	// Return the truncated string
	return result
}
