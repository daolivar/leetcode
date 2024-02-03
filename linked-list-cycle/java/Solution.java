/**
 * Solution to determine whether a linked list has a cycle using Floyd's
 * Tortoise and Hare algorithm.
 */
public class Solution {
    /**
     * Checks whether a linked list has a cycle.
     * 
     * Time Complexity: O(n) where n is the number of nodes in the linked list.
     * The algorithm uses two pointers, one moving twice as fast as the other. In
     * the worst case, where there is a cycle, the fast pointer will catch up with
     * the slow pointer after going around the cycle once.
     * 
     * Space Complexity: O(1) constant space. The algorithm only uses two pointers
     * (slow and fast), regardless of the size of the linked list.
     * 
     * @param head The head of the linked list.
     * @return {@code true} if there is a cycle in the linked list, {@code false}
     *         otherwise.
     */
    public boolean hasCycle(ListNode head) {
        // Initialize two pointers, slow and fast, both starting from the head of the
        // linked list.
        ListNode slow = head;
        ListNode fast = head != null ? head.next : null;

        // Move the pointers through the linked list.
        while (fast != null && fast.next != null) {
            // Check if the slow and fast pointers meet, indicating the presence of a cycle.
            if (slow == fast) {
                return true;
            }

            // Move the slow pointer one step at a time.
            slow = slow.next;

            // Move the fast pointer two steps at a time.
            fast = fast.next.next;
        }

        // If the fast pointer reaches the end of the linked list (null), there is no
        // cycle.
        return false;
    }
}
