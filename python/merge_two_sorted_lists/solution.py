class Solution(object):
    def mergeTwoLists(self, list1, list2):
        """
        Merge two sorted linked lists into one sorted linked list.

        :type list1: Optional[ListNode]  # The head of the first sorted linked list
        :type list2: Optional[ListNode]  # The head of the second sorted linked list
        :rtype: Optional[ListNode]  # The function returns the head of the merged sorted linked list
        """

        # Create a dummy node to act as the starting point of the merged list
        dummy = ListNode()
        # `current` pointer will be used to build the new list, starting from the dummy node
        current = dummy

        # Traverse both lists while neither list1 nor list2 is empty
        while list1 and list2:
            # Compare the current nodes of both lists
            if list1.val <= list2.val:
                # If the value of list1 is smaller or equal, append list1's node to the merged list
                current.next = list1
                # Move to the next node in list1
                list1 = list1.next
            else:
                # If list2's value is smaller, append list2's node to the merged list
                current.next = list2
                # Move to the next node in list2
                list2 = list2.next
            # Move the `current` pointer to the newly added node
            current = current.next

        # At the end of the while loop, one of the lists might still have remaining nodes
        # Attach the remaining nodes from either list1 or list2 to the merged list
        current.next = list1 if list1 else list2

        # Return the merged list, which starts after the dummy node
        return dummy.next
