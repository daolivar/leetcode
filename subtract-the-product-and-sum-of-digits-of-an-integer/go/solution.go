// subtractProductAndSum calculates the difference between the product and sum of the digits of a given integer.
//
// Time Complexity: O(log10(n)) where n is the input integer.
//                  The function iterates through each digit of the integer.
//
// Space Complexity: O(1) as only a constant amount of space is used.
//
// Parameters:
//   - n: The input integer for which to calculate the difference.
//
// Returns:
//   - The difference between the product and sum of the digits of the input integer.
func subtractProductAndSum(n int) int {
	// Initialize product and sum variables
	product, sum := 1, 0

	// Iterate through each digit of the integer
	for n > 0 {
		digit := n % 10
		product *= digit
		sum += digit
		n /= 10
	}

	// Return the difference between the product and sum
	return product - sum
}
