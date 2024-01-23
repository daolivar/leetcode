// toLowerCase converts a string to its lowercase equivalent.
//
// Time Complexity: O(n) where n is the length of the input string 's'.
//
// Space Complexity: O(n) as the result string is built character by character.
//
// Parameters:
//   - s: The input string to convert to lowercase.
//
// Returns:
//   - The lowercase representation of the input string 's'.
func toLowerCase(s string) string {
	result := ""

	// Iterate through each character in the input string 's'
	for _, char := range s {
		// Check if the character is an uppercase letter
		if char >= 'A' && char <= 'Z' {
			// Convert the uppercase letter to lowercase
			char = char + 'a' - 'A'
		}

		// Append the character to the result string
		result += string(char)
	}

	return result
}
