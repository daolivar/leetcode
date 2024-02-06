// removeTrailingZeros removes trailing zeros from a given string representation of a number.
//
// Time Complexity: O(n) where n is the length of the input string 'num'.
//
// Space Complexity: O(1) as the function uses constant space for variables.
//
// Parameters:
//   - num: The input string representing a number with potentially trailing zeros.
//
// Returns:
//   - The input string with trailing zeros removed. If the resulting string is empty, "0" is returned.
func removeTrailingZeros(num string) string {
	// Iterate through the characters of the input string from the end
	for i := len(num) - 1; i >= 0; i-- {
		// If a non-zero character is encountered, return the substring up to that character
		if num[i] != '0' {
			return num[:i+1]
		}
	}
	// If all characters are zeros or the input string is empty, return "0"
	return "0"
}
