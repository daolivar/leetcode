// numIdenticalPairs calculates the number of "good" pairs in the given slice 'nums',
// where a "good" pair is defined as a pair (i, j) where i < j and nums[i] == nums[j].
//
// Time Complexity: O(n) where n is the length of the input slice 'nums'.
//
// Space Complexity: O(n) as the function uses a map to store the frequency of each number in 'nums'.
//
// Parameters:
//   - nums: The input slice of integers.
//
// Returns:
//   - The count of "good" pairs in the input slice.
func numIdenticalPairs(nums []int) int {
    goodPairs := 0           // Initialize the count of "good" pairs to 0
    frequencyMap := make(map[int]int) // Initialize a map to store the frequency of each number

    // Iterate through each number in the input slice
    for _, n := range nums {
        goodPairs += frequencyMap[n] // Increment the count of "good" pairs by the frequency of the current number
        frequencyMap[n]++            // Increment the frequency of the current number
    }

    return goodPairs // Return the total count of "good" pairs
}
