class Solution(object):
    def reverseList(self, head):
        """
        Reverse a singly linked list.

        :type head: Optional[ListNode]  # head is the head node of the singly linked list
        :rtype: Optional[ListNode]  # The function returns the new head of the reversed linked list
        """

        # If the list is empty or has only one node, return the head as it is already "reversed"
        if not head or not head.next:
            return head

        # Initialize the previous node to None (this will become the new tail of the reversed list)
        previousNode = None
        # Set currentNode to the head of the list, which will traverse through the list
        currentNode = head

        # Iterate over the linked list until currentNode becomes None (end of the list)
        while currentNode:
            # Temporarily store the next node before changing the link direction
            nextNode = currentNode.next
            # Reverse the link: point the current node to the previous node
            currentNode.next = previousNode
            # Move previousNode and currentNode one step forward in the list
            previousNode = currentNode
            currentNode = nextNode

        # Return the previousNode, which will now be the new head of the reversed list
        return previousNode  # The new head of the reversed linked list
