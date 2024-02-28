// detectCapitalUse checks if the given word satisfies the rules for capitalization, which are:
// 1. All letters in the word are uppercase.
// 2. All letters in the word are lowercase.
// 3. Only the first letter in the word is uppercase, and the rest are lowercase.
//
// Time Complexity: O(n) where n is the length of the input word.
//
// Space Complexity: O(1) as the function performs constant-time operations and does not use extra space.
//
// Parameters:
//   - word: The input word to check for capitalization.
//
// Returns:
//   - true if the word satisfies the rules for capitalization, false otherwise.
func detectCapitalUse(word string) bool {
	// Check if the word is all uppercase or all lowercase, or if the first letter is uppercase and the rest are lowercase
	return word == strings.ToUpper(word) || // Rule 1: All letters in the word are uppercase.
		word == strings.ToLower(word) || // Rule 2: All letters in the word are lowercase.
		(word[0] >= 'A' && word[0] <= 'Z' && word[1:] == strings.ToLower(word[1:])) // Rule 3: Only the first letter is uppercase.
}
