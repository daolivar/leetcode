// findTheArrayConcVal calculates the concatenated value of an array of integers by concatenating each pair of integers from the array.
//
// Time Complexity: O(n) where n is the length of the input array 'nums'.
//
// Space Complexity: O(1) as the function uses constant space for variables.
//
// Parameters:
//   - nums: The input array of integers.
//
// Returns:
//   - The concatenated value of the array as an int64.
func findTheArrayConcVal(nums []int) int64 {
	var value int64

	// Iterate through the array until it's empty
	for len(nums) > 0 {
		// If there are at least two elements in the array,
		// concatenate the first and last elements and add the result to the value.
		if len(nums) >= 2 {
			value += concatValue(nums[0], nums[len(nums)-1])
			nums = nums[1 : len(nums)-1]
		} else {
			// If there's only one element left in the array, add it to the value.
			value += int64(nums[0])
			nums = nums[:0]
		}
	}

	return value
}

// concatValue concatenates two integers and returns the result as an int64.
//
// Time Complexity: O(1) as the function performs a constant-time operation.
//
// Parameters:
//   - a: The first integer to concatenate.
//   - b: The second integer to concatenate.
//
// Returns:
//   - The concatenated value of the two integers as an int64.
func concatValue(a, b int) int64 {
	// Convert the concatenated string representation of the integers to an int64.
	num, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	if err != nil {
		// Handle any errors during conversion.
		fmt.Println("Error: ", err)
	}
	return int64(num)
}
