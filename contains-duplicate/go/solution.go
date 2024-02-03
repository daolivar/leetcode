// containsDuplicate checks if there are any duplicate elements in the given slice of integers.
//
// Time Complexity: O(n) where n is the length of the input slice 'nums'.
//
// Space Complexity: O(n) as the function uses a map to keep track of unique elements.
//
// Parameters:
//   - nums: The input slice of integers to check for duplicates.
//
// Returns:
//   - true if there are any duplicate elements, false otherwise.
func containsDuplicate(nums []int) bool {
	// Create a map to store seen elements
	seen := make(map[int]struct{})

	// Iterate through the slice of integers
	for _, num := range nums {
		// If the element is already in the map, it's a duplicate
		if _, ok := seen[num]; ok {
			return true
		}

		// Add the element to the map
		seen[num] = struct{}{}
	}

	// No duplicates found
	return false
}
