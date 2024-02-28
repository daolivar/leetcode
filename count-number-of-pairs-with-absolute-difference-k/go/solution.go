// countKDifference calculates the number of pairs of integers in the given slice 'nums'
// that have a difference of 'k'.
//
// Time Complexity: O(n) where n is the length of the input slice 'nums'.
//
// Space Complexity: O(n) as the function uses a map to store the frequency of each number in 'nums'.
//
// Parameters:
//   - nums: The input slice of integers.
//   - k: The target difference to search for.
//
// Returns:
//   - The count of pairs with a difference of 'k' in the input slice.
func countKDifference(nums []int, k int) int {
	count := 0
	freq := make(map[int]int)

	// Iterate through each number in the input slice
	for _, num := range nums {
		// Calculate the complements for the current number
		complement1 := num + k
		complement2 := num - k

		// Increment the count by the sum of frequencies of complements
		count += freq[complement1] + freq[complement2]

		// Increment the frequency of the current number
		freq[num]++
	}

	return count // Return the total count of pairs with a difference of 'k'
}
