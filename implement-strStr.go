package leetcode

import "strings"

// Using Go strings built-in function Index()
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	return strings.Index(haystack, needle)
}

// Using string range and for loop
// func strStr(haystack string, needle string) int {
// 	if needle == "" {
// 		return 0
// 	}
// 	for i, s := range haystack {
// 		if string(s) == string(needle[0]) {
// 			if i+len(needle) > len(haystack) {
// 				return -1
// 			}
// 			if haystack[i:i+len(needle)] == needle {
// 				return i
// 			}
// 		}
// 	}
// 	return -1
// }
