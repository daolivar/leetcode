class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        """
        Merges two sorted linked lists into one sorted linked list.

        Given two sorted linked lists, this function combines them into a single sorted linked list 
        by comparing node values from both lists and linking them in the appropriate order.

        Parameters:
            list1 (Optional[ListNode]): The head of the first sorted linked list.
            list2 (Optional[ListNode]): The head of the second sorted linked list.

        Returns:
            Optional[ListNode]: The head of the merged sorted linked list, or None if both lists are empty.
        """
        # Handle edge cases where one or both lists are None
        if not list1:
            return list2
        if not list2:
            return list1

        # Create a dummy node to simplify the merge process
        dummy = ListNode(0)
        current = dummy  # Pointer to track the current position in the merged list

        # Iterate through both lists as long as there are elements in both
        while list1 and list2:
            if list1.val <= list2.val:
                current.next = list1  # Attach the smaller node to the merged list
                list1 = list1.next    # Move the pointer to the next node in list1
            else:
                current.next = list2  # Attach the smaller node from list2
                list2 = list2.next    # Move the pointer to the next node in list2
            current = current.next    # Move to the next position in the merged list

        # If one list is exhausted, attach the remaining nodes from the other list
        current.next = list1 if list1 else list2

        # Return the merged list, skipping the dummy node
        return dummy.next
