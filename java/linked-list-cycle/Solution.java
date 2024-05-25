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
        if (head == null || head.next == null) {
            return false; // If the list is empty or has only one node, no cycle can exist
        }

        ListNode slow = head; // Slow pointer moves one step at a time
        ListNode fast = head.next; // Fast pointer moves two steps at a time

        while (fast != null && fast.next != null) {
            if (slow == fast) {
                return true; // If the slow and fast pointers meet, a cycle exists
            }
            slow = slow.next; // Move slow pointer one step
            fast = fast.next.next; // Move fast pointer two steps
        }

        return false; // If we reach the end of the list, no cycle exists
    }
}
