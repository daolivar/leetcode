// findTheDifference finds the differing character between two strings 's' and 't'.
// This version uses a map with rune keys for counting character occurrences.
//
// Time Complexity: O(max(len(s), len(t))) as the function iterates through both strings.
//
// Space Complexity: O(min(len(s), 26)) as the function uses a map to count character occurrences,
// and the number of unique characters is at most 26 (the number of lowercase and uppercase English letters).
//
// Parameters:
//   - s: The first input string.
//   - t: The second input string, which is the same as 's' with an additional character.
//
// Returns:
//   - The differing character that is present in 't' but not in 's'.
func findTheDifference(s string, t string) byte {
	mp := map[rune]int{}

	// Count character occurrences in 's'
	for _, c := range s {
		mp[c]++
	}

	// Subtract character occurrences in 't'
	for _, c := range t {
		mp[c]--

		// If the count becomes negative, return the differing character
		if mp[c] < 0 {
			return byte(c)
		}
	}

	// Default return if no difference is found
	return 0
}
