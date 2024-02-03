/**
 * Solution to check whether two binary trees are the same.
 */
public class Solution {
    /**
     * Checks if two binary trees are the same.
     * 
     * Time Complexity: O(n), where n is the total number of nodes in the larger
     * of the two trees. In the worst case, the method needs to visit each node once
     * to determine if the trees are the same.
     * 
     * Space Complexity: O(h), where h is the height of the taller tree. The space
     * complexity is determined by the recursive calls on the call stack. In the
     * worst case, the maximum depth of the call stack is equal to the height of the
     * tree.
     * 
     * @param p The root of the first binary tree.
     * @param q The root of the second binary tree.
     * @return {@code true} if the trees are the same, {@code false} otherwise.
     */
    public boolean isSameTree(TreeNode p, TreeNode q) {
        // Base case: both trees are empty (null), indicating they are the same.
        if (p == null && q == null) {
            return true;
        }

        // If one of the trees is empty while the other is not, or their values are
        // different, return false.
        if ((p == null || q == null) || (p.val != q.val)) {
            return false;
        }

        // Recursively check the left and right subtrees.
        return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
    }
}
