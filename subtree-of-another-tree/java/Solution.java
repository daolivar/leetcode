/**
 * Solution to check if a tree is a subtree of another tree.
 */
class Solution {
    /**
     * Checks if a given tree is a subtree of another tree.
     * 
     * Time Complexity: O(m * n), where 'm' is the number of nodes in the main tree
     * and 'n' is the number of nodes in the potential subtree. In the worst case,
     * the method may need to compare each node of the larger tree with the entire
     * subtree of the smaller tree.
     * 
     * Space Complexity: O(max(m, n)), where 'm' is the height of the main tree and
     * 'n' is the height of the potential subtree. In the worst case, where one of
     * the trees is completely unbalanced (skewed), the height 'h' could be equal to
     * 'm' or 'n', resulting in O(m) or O(n) space complexity. In a balanced tree,
     * the height 'h' is O(log m) or O(log n), depending on which tree is larger.
     * 
     * @param root    The root of the main tree.
     * @param subRoot The root of the potential subtree.
     * @return {@code true} if subRoot is a subtree of root, {@code false}
     *         otherwise.
     */
    public boolean isSubtree(TreeNode root, TreeNode subRoot) {
        // Base case: If both roots are the same node, they are considered the same
        // subtree.
        if (root == subRoot) {
            return true;
        }

        // Base case: If either root or subRoot is null, subRoot cannot be a subtree.
        if (root == null || subRoot == null) {
            return false;
        }

        // Check if the current nodes are equal and if the subtrees rooted at these
        // nodes are identical.
        if (root.val == subRoot.val && isSameTree(root, subRoot)) {
            return true;
        }

        // Recursively check the left and right subtrees.
        return isSubtree(root.left, subRoot) || isSubtree(root.right, subRoot);
    }

    /**
     * Helper method to check if two trees are identical.
     *
     * @param p The root of the first tree.
     * @param q The root of the second tree.
     * @return True if the trees are identical, false otherwise.
     */
    public boolean isSameTree(TreeNode p, TreeNode q) {
        // Base case: both trees are empty (null), indicating they are the same
        if (p == null && q == null) {
            return true;
        }

        // If one of the trees is empty while the other is not, or their values are
        // different, return false
        if ((p == null || q == null) || (p.val != q.val)) {
            return false;
        }

        // Recursively check the left and right subtrees
        return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
    }
}
