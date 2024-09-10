class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        """
        Reverses a singly linked list.

        This function takes the head of a singly linked list and reverses the direction of its nodes,
        returning the new head of the reversed list.

        Parameters:
            head (Optional[ListNode]): The head of the linked list to be reversed.

        Returns:
            Optional[ListNode]: The new head of the reversed linked list, or None if the list is empty.
        """
        # Base case: if the list is empty, return None
        if not head:
            return None

        prev = None  # Initialize the 'prev' pointer to None (as the new tail of the list)
        curr = head  # Start the 'curr' pointer at the head of the list

        while curr:
            next_node = curr.next  # Save the next node to move forward later
            curr.next = prev       # Reverse the current node's link to point to the previous node
            prev = curr            # Move 'prev' forward to the current node
            curr = next_node       # Move 'curr' forward to the next node in the list

        return prev  # 'prev' will be the new head of the reversed list
