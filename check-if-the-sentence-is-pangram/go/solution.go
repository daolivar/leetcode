// checkIfPangram checks if the given sentence contains all 26 letters of the English alphabet at least once.
//
// Time Complexity: O(n) where n is the length of the input sentence.
//
// Space Complexity: O(1) as the function uses a map with at most 26 entries to track the occurrence of letters.
//
// Parameters:
//   - sentence: The input sentence to check for pangrammaticity.
//
// Returns:
//   - true if the sentence is a pangram (contains all 26 letters), false otherwise.
func checkIfPangram(sentence string) bool {
	alphabet := map[rune]bool{} // Initialize a map to store the occurrence of each letter

	// Iterate through each character in the sentence
	for _, c := range sentence {
		alphabet[c] = true // Mark the current character as present in the alphabet
	}

	return len(alphabet) == 26 // Return true if all 26 letters are present in the alphabet
}
