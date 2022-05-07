package leetcode

func getDecimalValue(head *ListNode) int {
	var sum int
	curr := head
	for curr != nil {
		sum = sum*2 + curr.Val
		curr = curr.Next
	}
	return sum
}
