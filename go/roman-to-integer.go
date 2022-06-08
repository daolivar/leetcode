package leetcode

func romanToInt(s string) int {
	var result int
	mp := getMap()
	for i := 0; i <= len(s)-1; i++ {
		if i == len(s)-1 {
			result += mp[s[i]]
		} else {
			if mp[s[i]] < mp[s[i+1]] {
				result -= mp[s[i]]
			} else {
				result += mp[s[i]]
			}
		}
	}
	return result
}

// helper function that returns a map of roman numeral -> integer pairs
func getMap() map[byte]int {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	return m
}
