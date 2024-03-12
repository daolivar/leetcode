// isFascinating checks if the given number 'n' is fascinating, which means when it's multiplied by 2 and 3,
// concatenating the results results in a number that contains each digit from 1 to 9 exactly once.
//
// Time Complexity: O(1) as the function performs constant-time operations.
//
// Space Complexity: O(1) as the function uses a fixed-size array to track encountered digits.
//
// Parameters:
//   - n: The input number to check for fascination.
//
// Returns:
//   - true if the number is fascinating, false otherwise.
func isFascinating(n int) bool {
	// Generate the concatenated number
	n2, n3 := 2*n, 3*n
	num := fmt.Sprintf("%d%d%d", n, n2, n3)
	// Validate the concatenated number
	return validate(num)
}

// validate checks if the given number 'num' contains each digit from 1 to 9 exactly once.
//
// Time Complexity: O(1) as the function performs constant-time operations.
//
// Space Complexity: O(1) as the function uses a fixed-size array to track encountered digits.
//
// Parameters:
//   - num: The input number as a string to validate.
//
// Returns:
//   - true if the number contains each digit from 1 to 9 exactly once, false otherwise.
func validate(num string) bool {
	encountered := [10]bool{} // Initialize an array to track encountered digits
	// Iterate through each digit in the number
	for _, digit := range num {
		// Check if the digit is '0' or if it has already been encountered
		if digit == '0' || encountered[digit-'0'] {
			return false // Return false if the condition is met
		}
		encountered[digit-'0'] = true // Mark the digit as encountered
	}
	return true // Return true if all digits are unique
}
