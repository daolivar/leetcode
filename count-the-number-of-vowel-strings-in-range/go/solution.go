// vowels is a map representing vowels where the key is the vowel character and the value is true.
var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

// vowelStrings counts the number of strings in the given slice 'words' where both the first and last characters are vowels.
//
// Time Complexity: O(n) where n is the number of words in the given slice.
//
// Space Complexity: O(1) as the function uses constant space for variables.
//
// Parameters:
//   - words: The slice of strings to be checked.
//   - left: The index of the first word to consider.
//   - right: The index of the last word to consider.
//
// Returns:
//   - The count of strings where both the first and last characters are vowels.
func vowelStrings(words []string, left int, right int) int {
	count := 0

	// Iterate through the words in the given range [left, right]
	for i := left; i <= right; i++ {
		// Check if both the first and last characters of the current word are vowels
		if isVowel(words[i][0]) && isVowel(words[i][len(words[i])-1]) {
			count++
		}
	}

	return count
}

// isVowel checks if a given byte represents a vowel.
//
// Time Complexity: O(1) as the function performs a constant-time map lookup.
//
// Parameters:
//   - c: The byte to be checked.
//
// Returns:
//   - true if the byte represents a vowel, false otherwise.
func isVowel(c byte) bool {
	return vowels[c]
}
