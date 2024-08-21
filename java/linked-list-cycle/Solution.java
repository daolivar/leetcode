/**
 * Solution class for LeetCode problem 141. Linked List Cycle
 */
public class Solution {

    /**
     * Detects if a cycle exists in a linked list using Floyd’s Tortoise and Hare algorithm.
     *
     * @param head the head node of the linked list
     * @return true if there is a cycle in the linked list, false otherwise
     */
    public boolean hasCycle(ListNode head) {
        if (head == null) {
            return false; // If the list is empty no cycle can exist
        }

        ListNode slow = head;
        ListNode fast = head;

        while (fast != null && fast.next != null) {
            slow = slow.next; // Move slow pointer one step
            fast = fast.next.next; // Move fast pointer two steps

            if (slow == fast) {
                return true; // If the slow and fast pointers meet, a cycle exists
            }
        }

        return false; // If we reach the end of the list, no cycle exists
    }
}
