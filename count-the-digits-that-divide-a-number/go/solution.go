// countDigits counts the number of digits in a positive integer 'num' that evenly divide 'num'.
//
// Time Complexity: O(log10(num)) where num is the input integer. This is because
// the loop iterates through each digit of 'num'.
//
// Space Complexity: O(1) as only a constant amount of space is used.
//
// Parameters:
//   - num: The positive input integer to count the digits for.
//
// Returns:
//   - The count of digits in 'num' that evenly divide 'num'.
func countDigits(num int) int {
	count := 0
	original := num

	// Iterate through each digit of 'num'
	for num > 0 {
		// Extract the current digit
		digit := num % 10

		// Check if the digit divides the original number evenly
		if original%digit == 0 {
			count++
		}

		// Move to the next digit
		num /= 10
	}

	return count
}
