class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        """
        Determines if a linked list has a cycle in it.

        A cycle occurs if a node's `next` pointer references a previous node in the list, 
        forming a loop.

        Parameters:
            head (Optional[ListNode]): The head of the linked list to check for a cycle.

        Returns:
            bool: True if there is a cycle in the linked list, False otherwise.
        
        Explanation:
            This method uses the Floyd's Tortoise and Hare algorithm, which employs two pointers:
            - `slow`: Moves one step at a time.
            - `fast`: Moves two steps at a time.
            If there is a cycle, these two pointers will eventually meet.
        """
        
        # If the list is empty, there can't be a cycle
        if not head:
            return False
        
        # Initialize two pointers, slow and fast
        slow = head
        fast = head
        
        # Traverse the list with the two pointers
        while fast and fast.next:
            slow = slow.next            # Move slow pointer by one step
            fast = fast.next.next       # Move fast pointer by two steps
            
            # If the two pointers meet, a cycle exists
            if slow == fast:
                return True
        
        # If we exit the loop, it means there's no cycle
        return False
