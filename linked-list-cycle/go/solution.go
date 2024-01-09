// hasCycle determines whether a linked list has a cycle using Floyd's Tortoise and Hare algorithm.
//
// Time Complexity: O(n), where 'n' is the number of nodes in the linked list. The algorithm uses
// two pointers, one moving twice as fast as the other. In the worst case, where there is a cycle,
// the fast pointer will catch up with the slow pointer after going around the cycle once.
//
// Space Complexity: O(1), constant space. The algorithm only uses two pointers (slow and fast),
// regardless of the size of the linked list.
//
// Parameters:
// - head: The head of the linked list.
//
// Returns:
// - true if there is a cycle in the linked list, false otherwise.
func hasCycle(head *ListNode) bool {
	// Initialize two pointers, slow and fast, both starting from the head of the linked list.
	slow, fast := head, head

	// Move the pointers through the linked list.
	for fast != nil && fast.Next != nil {
		// Move the slow pointer one step at a time.
		slow = slow.Next

		// Move the fast pointer two steps at a time.
		fast = fast.Next.Next

		// Check if the slow and fast pointers meet, indicating the presence of a cycle.
		if slow == fast {
			return true
		}
	}

	// If the fast pointer reaches the end of the linked list (nil), there is no cycle.
	return false
}
