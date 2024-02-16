// countSymmetricIntegers counts the number of symmetric integers within the range [low, high] (inclusive).
//
// Time Complexity: O(n) where n is the number of integers within the range [low, high].
//
// Space Complexity: O(1) as the function uses constant space for variables.
//
// Parameters:
//   - low: The lower bound of the range.
//   - high: The upper bound of the range.
//
// Returns:
//   - The count of symmetric integers within the specified range.
func countSymmetricIntegers(low int, high int) int {
	result := 0

	// Iterate through each integer within the range [low, high]
	for num := low; num <= high; num++ {
		// If the integer is symmetric, increment the result count
		if isSymmetric(num) {
			result++
		}
	}

	return result
}

// isSymmetricInteger checks if an integer is symmetric.
//
// Time Complexity: O(1) as the function performs a constant-time operation based on the number of digits in 'num'.
//
// Parameters:
//   - num: The integer to check for symmetry.
//
// Returns:
//   - true if the integer is symmetric, false otherwise.
func isSymmetric(num int) bool {
	// Check if 'num' has two or four digits, as these are the only cases where symmetry is possible
	if num >= 10 && num <= 99 {
		// For two-digit integers, compare the tens and units digits
		return num/10 == num%10
	}
	if num >= 1000 && num <= 9999 {
		// For four-digit integers, split the number into two halves and compare the sums of their digits
		left := num / 100
		right := num % 100
		return left/10+left%10 == right/10+right%10
	}
	return false // For all other cases, the number is not symmetric
}
