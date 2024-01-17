// smallestEvenMultiple returns the smallest even multiple of the given integer.
//
// Time Complexity: O(1) as the operation involves checking and computing a simple arithmetic expression.
//
// Space Complexity: O(1) as no additional data structures are used.
//
// Parameters:
//   - n: The input integer.
//
// Returns:
//   - The smallest even multiple of n.
//     If n is already even, the result is n; otherwise, the result is 2n.
func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return 2 * n
}
