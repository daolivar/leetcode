class Solution(object):
    def hasCycle(self, head):
        """
        Determine if a linked list contains a cycle.

        :type head: ListNode  # head is the head node of the singly linked list
        :rtype: bool  # The function returns a boolean value (True if the list contains a cycle, False otherwise)
        """

        # If the list is empty or contains only one node, it can't have a cycle
        if not head or not head.next:
            return False

        # Initialize two pointers, slow and fast
        slow = head  # Moves one step at a time
        fast = head  # Moves two steps at a time

        # Traverse the linked list using the two pointers
        while fast and fast.next:
            slow = slow.next  # Move slow pointer one step forward
            fast = fast.next.next  # Move fast pointer two steps forward

            # If slow and fast meet, there's a cycle in the list
            if slow == fast:
                return True

        # If the fast pointer reaches the end (no cycle), return False
        return False
