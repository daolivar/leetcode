package leetcode

// Recursive

// Iterative in-place
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var output *ListNode = nil
	if list1.Val < list2.Val {
		output = list1
		list1 = list1.Next
	} else {
		output = list2
		list2 = list2.Next
	}
	var curr *ListNode = output
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			curr.Next = list2
			list2 = list2.Next
		} else {
			curr.Next = list1
			list1 = list1.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}
	return output
}
