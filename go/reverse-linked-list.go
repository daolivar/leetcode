package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev, next *ListNode = nil, nil
	curr := head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head = prev
	return head
}
