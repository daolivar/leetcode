// separateDigits separates the digits of each number in the input slice 'nums'.
// It uses the extractDigits function to extract individual digits and appends them
// to the result slice.
//
// Time Complexity: O(m * n) where m is the length of the input slice 'nums', and n is
// the maximum number of digits in any number from 'nums'.
//
// Space Complexity: O(m * n) as the result slice stores the extracted digits.
//
// Parameters:
//   - nums: The input slice containing integers.
//
// Returns:
//   - A new slice containing the separated digits of all numbers from 'nums'.
func separateDigits(nums []int) []int {
	result := []int{}
	for _, n := range nums {
		extractDigits(n, &result)
	}
	return result
}

// extractDigits extracts individual digits from the given integer 'n'
// and appends them to the result slice.
//
// Time Complexity: O(log10(n)) where n is the input integer.
//
// Space Complexity: O(log10(n)) as the digits are stored in a string before conversion.
//
// Parameters:
//   - n: The input integer from which digits are to be extracted.
//   - result: A pointer to the slice where the extracted digits will be appended.
func extractDigits(n int, result *[]int) {
	digits := strings.Split(strconv.Itoa(n), "")
	for _, digit := range digits {
		d, _ := strconv.Atoi(digit)
		*result = append(*result, d)
	}
}
