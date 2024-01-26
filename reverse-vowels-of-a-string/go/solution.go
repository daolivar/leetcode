// reverseVowels reverses the vowels in a given string.
//
// Time Complexity: O(n) where n is the length of the input string 's'.
//                  The function iterates through the string once.
//
// Space Complexity: O(n) as the string is converted to a rune slice.
//
// Parameters:
//   - s: The input string in which vowels will be reversed.
//
// Returns:
//   - The string with reversed vowels.
func reverseVowels(s string) string {
	// Convert the string to a rune slice for in-place modifications
	runes := []rune(s)
	i, j := 0, len(s)-1

	// Iterate through the string from both ends
	for i < j {
		// Find the next vowel from the beginning
		for i < j && !isVowel(runes[i]) {
			i++
		}

		// Find the next vowel from the end
		for i < j && !isVowel(runes[j]) {
			j--
		}

		// Swap vowels
		if i < j {
			runes[i], runes[j] = runes[j], runes[i]
			i++
			j--
		}
	}

	// Convert the rune slice back to a string and return
	return string(runes)
}

// isVowel checks if a rune is a vowel.
//
// Time Complexity: O(1) as the function performs a constant-time operation.
//
// Parameters:
//   - c: The rune to check for being a vowel.
//
// Returns:
//   - true if the rune is a vowel, false otherwise.
func isVowel(c rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, c)
}
