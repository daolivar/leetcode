// isPalindrome checks if a given string is a palindrome, ignoring non-alphanumeric characters and case.
//
// Time Complexity: O(n) where n is the length of the input string 's'.
//                  The function iterates through the string once.
//
// Space Complexity: O(n) as the string is converted to a rune slice.
//
// Parameters:
//   - s: The input string to check for palindrome.
//
// Returns:
//   - true if the string is a palindrome, false otherwise.
func isPalindrome(s string) bool {
	// Convert the string to lowercase and create a rune slice for in-place modifications
	s = strings.ToLower(s)
	runes := []rune(s)
	i, j := 0, len(runes)-1

	// Iterate through the string from both ends
	for i < j {
		// Find the next alphanumeric character from the beginning
		for i < j && !isAlphaNumeric(runes[i]) {
			i++
		}

		// Find the next alphanumeric character from the end
		for i < j && !isAlphaNumeric(runes[j]) {
			j--
		}

		// Check if the characters at positions i and j are not equal
		if runes[i] != runes[j] {
			return false
		}

		// Move the indices towards the center
		i++
		j--
	}

	// The string is a palindrome
	return true
}

// isAlphaNumeric checks if a rune is an alphanumeric character.
//
// Time Complexity: O(1) as the function performs a constant-time operation.
//
// Parameters:
//   - c: The rune to check for being alphanumeric.
//
// Returns:
//   - true if the rune is alphanumeric, false otherwise.
func isAlphaNumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
