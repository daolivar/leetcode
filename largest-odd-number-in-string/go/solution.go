// largestOddNumber finds the largest odd number that can be formed by removing digits from the end of a given string 'num'.
//
// Time Complexity: O(n) where n is the length of the input string 'num'.
//
// Space Complexity: O(1) as the function uses constant space for variables.
//
// Parameters:
//   - num: The input string representing a number.
//
// Returns:
//   - The largest odd number that can be formed by removing digits from the end of 'num'. If no such number exists, an empty string is returned.
func largestOddNumber(num string) string {
	// Iterate through the characters of the input string from the end
	for i := len(num) - 1; i >= 0; i-- {
		// If an odd digit is encountered, return the substring up to that digit
		if num[i]%2 != 0 {
			return num[:i+1]
		}
	}
	// If no odd digits are found, return an empty string
	return ""
}
