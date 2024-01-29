// reverseOnlyLetters reverses the alphabetical characters in a given string,
// while leaving non-alphabetical characters in their original positions.
//
// Time Complexity: O(n) where n is the length of the input string 's'.
//                  The function iterates through the string once.
//
// Space Complexity: O(n) as the string is converted to a rune slice.
//
// Parameters:
//   - s: The input string to reverse only the alphabetical characters.
//
// Returns:
//   - The string with reversed alphabetical characters.
func reverseOnlyLetters(s string) string {
	// Convert the string to a rune slice for in-place modifications
	runes := []rune(s)
	i, j := 0, len(runes)-1

	// Iterate through the string from both ends
	for i < j {
		// Find the next alphabetical character from the beginning
		for i < j && !isAlpha(runes[i]) {
			i++
		}

		// Find the next alphabetical character from the end
		for i < j && !isAlpha(runes[j]) {
			j--
		}

		// Swap alphabetical characters
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	// Convert the rune slice back to a string and return
	return string(runes)
}

// isAlpha checks if a rune is an alphabetical character.
//
// Time Complexity: O(1) as the function performs a constant-time operation.
//
// Parameters:
//   - c: The rune to check for being alphabetical.
//
// Returns:
//   - true if the rune is alphabetical, false otherwise.
func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
