// canConstruct checks if it's possible to construct the ransom note from the characters
// available in the magazine. Each character in the magazine can only be used once.
//
// Time Complexity: O(max(len(ransomNote), len(magazine))) as the function iterates through both strings.
//
// Space Complexity: O(min(len(magazine), 26)) as the function uses a map to count character occurrences,
// and the number of unique characters is at most 26 (the number of lowercase and uppercase English letters).
//
// Parameters:
//   - ransomNote: The ransom note string to be constructed.
//   - magazine: The magazine string containing available characters.
//
// Returns:
//   - true if it's possible to construct the ransom note, false otherwise.
func canConstruct(ransomNote string, magazine string) bool {
	mp := map[rune]int{}

	// Count character occurrences in 'magazine'
	for _, c := range magazine {
		mp[c]++
	}

	// Subtract character occurrences for each character in 'ransomNote'
	for _, c := range ransomNote {
		mp[c]--

		// If the count becomes negative, return false (cannot construct the ransom note)
		if mp[c] < 0 {
			return false
		}
	}

	// It's possible to construct the ransom note
	return true
}
