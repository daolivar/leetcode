// twoSum finds two numbers in the given slice 'nums' that add up to the 'target' value.
// It uses a hash map to store the indices of numbers encountered during the iteration.
// The function returns a slice containing the indices of the two numbers that sum up to the
// target. If no such indices are found, an empty slice is returned.
//
// Time Complexity: O(n) where n is the length of the input array nums. In the worst case,
// the algorithm iterates through the entire array once.
//
// Space Complexity: O(n) where n is the length of the input array nums. In the worst case,
// the algorithm may store all elements of the array in the map along with their indices.
//
// Parameters:
//   - nums: The input slice of integers.
//   - target: The target sum value.
//
// Returns:
//   - result: A slice containing the indices of the two numbers that add up to the target.
//             If no such indices are found, an empty slice is returned.
func twoSum(nums []int, target int) []int {
	// Create a map to store the indices of numbers encountered during the iteration.
	indexMap := make(map[int]int)

	// Iterate through the numbers in the slice.
	for index, num := range nums {
		// Calculate the difference between the target and the current number.
		difference := target - num

		// Check if the difference exists in the map.
		if val, ok := indexMap[difference]; ok {
			// If found, return the indices of the two numbers that add up to the target.
			return []int{val, index}
		}

		// Otherwise, add the current number and its index to the map.
		indexMap[num] = index
	}

	// If no such indices are found, return an empty slice.
	return []int{}
}
