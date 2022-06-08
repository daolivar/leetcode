package leetcode

func isPalindrome(x int) bool {
	reverse, xCopy := 0, x
	if x < 0 {
		x *= -1
	}
	for xCopy != 0 {
		reverse = reverse*10 + (xCopy % 10)
		xCopy /= 10
	}
	return reverse == x
}
