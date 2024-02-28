// sortArrayByParity sorts the elements of the given slice 'nums' such that all even elements appear before
// all odd elements. The relative order of even and odd elements within the slice is preserved.
//
// Time Complexity: O(n) where n is the length of the input slice 'nums'.
//
// Space Complexity: O(1) as the function sorts the input slice in-place without using extra space.
//
// Parameters:
//   - nums: The input slice of integers to be sorted by parity.
//
// Returns:
//   - The sorted slice where even elements appear before odd elements.
func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1 // Initialize two pointers, 'left' and 'right', at the start and end of the slice respectively
	// Iterate until 'left' and 'right' pointers meet
	for left < right {
		// Move 'left' pointer to the right until it points to an odd element
		for nums[left]%2 == 0 && left < right {
			left++
		}
		// Move 'right' pointer to the left until it points to an even element
		for nums[right]%2 == 1 && left < right {
			right--
		}
		// Swap the elements pointed to by 'left' and 'right' pointers
		nums[left], nums[right] = nums[right], nums[left]
		left++  // Move 'left' pointer to the right
		right-- // Move 'right' pointer to the left
	}
	return nums // Return the sorted slice
}
