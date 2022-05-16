package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := getListSize(head)
	if n == size {
		return head.Next
	} else {
		curr := head
		nodes := size - n - 1
		for nodes > 0 {
			curr = curr.Next
			nodes--
		}
		curr.Next = curr.Next.Next
	}
	return head
}

// helper function to get size of linked list
func getListSize(head *ListNode) int {
	var size int
	curr := head
	for curr != nil {
		size++
		curr = curr.Next
	}
	return size
}
