package leetcode

func findTheDifference(s string, t string) byte {
	var sumS, sumT byte
	sb := []byte(s)
	tb := []byte(t)
	for _, c := range sb {
		sumS += c
	}
	for _, c := range tb {
		sumT += c
	}
	return sumT - sumS
}
