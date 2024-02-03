/**
 * Solution for reversing a linked list.
 */
class Solution {
    /**
     * Reverses a linked list in-place.
     * 
     * Time Complexity: O(n) where n is the number of nodes in the linked list.
     * This is because the method iterates through each node in the list exactly
     * once.
     * 
     * Space Complexity: O(1) as the algorithm uses a constant amount of extra
     * space regardless of the size of the input linked list.
     * 
     * @param head The head of the linked list to be reversed.
     * @return The new head of the reversed linked list.
     */
    public ListNode reverseList(ListNode head) {
        // Initialize the previous node to null since it will be the last node in the
        // reversed list.
        ListNode prev = null;

        // Traverse the original linked list, reversing the direction of each link.
        while (head != null) {
            // Save the next node in the original list before modifying the link.
            ListNode next = head.next;

            // Reverse the link to point to the previous node.
            head.next = prev;

            // Move the previous node and current node pointers one step forward.
            prev = head;
            head = next;
        }

        // The previous node is now the new head of the reversed linked list.
        return prev;
    }
}
