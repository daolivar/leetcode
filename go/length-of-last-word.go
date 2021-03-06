package leetcode

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && length != 0 {
			return length
		} else if s[i] != ' ' {
			length++
		}
	}
	return length
}
