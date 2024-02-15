// isCircularSentence checks if a given sentence forms a circular sentence, where the first and last characters are the same.
//
// Time Complexity: O(n) where n is the length of the input sentence.
//
// Space Complexity: O(1) as the function uses constant space for variables.
//
// Parameters:
//   - sentence: The input sentence to be checked.
//
// Returns:
//   - true if the sentence forms a circular sentence, false otherwise.
func isCircularSentence(sentence string) bool {
	// Check if the first and last characters of the sentence are different,
	// as a circular sentence should start and end with the same character.
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}

	// Iterate through the characters of the sentence starting from the third character
	for i := 2; i < len(sentence); i++ {
		// If a space character is followed by a different character,
		// the sentence does not form a circular sentence.
		if sentence[i-1] == ' ' && sentence[i-2] != sentence[i] {
			return false
		}
	}

	// If no inconsistencies are found, the sentence forms a circular sentence.
	return true
}
