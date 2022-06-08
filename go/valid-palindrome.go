package leetcode

import (
	"log"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = processStringData(s)
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

// helper func:
// to lowercase
// removes special characters
func processStringData(s string) string {
	s = strings.ToLower(s)
	reg, err := regexp.Compile(`[^a-zA-Z0-9]`)
	if err != nil {
		log.Fatal(err)
	}
	s = reg.ReplaceAllString(s, "")
	return s
}
