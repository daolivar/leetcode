func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	n := len(t)
	mp := make(map[byte]int)
	for i := 0; i < n; i++ {
		mp[s[i]]++
		mp[t[i]]--
	}
	for _, val := range mp {
		if val != 0 {
			return false
		}
	}
	return true
}
