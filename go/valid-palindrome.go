package leetcode

import "strings"

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !validCharacter(s[i]) {
			i++
			continue
		}
		if !validCharacter(s[j]) {
			j--
			continue
		}
		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func validCharacter(a byte) bool {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}
