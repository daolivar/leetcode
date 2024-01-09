/**
 * Solution to get the maximum depth of a binary tree.
 */
class Solution {
    /**
     * Computes the maximum depth of a binary tree.
     * 
     * Time Complexity: O(n), where 'n' is the number of nodes in the binary tree.
     * This is because the function visits each node exactly once, and the work done
     * at each node is constant.
     * 
     * Space Complexity: O(h), where 'h' is the height of the binary tree. The space
     * complexity is determined by the recursive calls on the call stack during the
     * traversal. Each recursive call consumes space on the call stack.In the worst
     * case, where the tree is completely unbalanced (skewed), the height 'h' could
     * be equal to 'n', resulting in O(n) space complexity. In a balanced tree, the
     * height 'h' is O(log n).
     * 
     * @param root The root node of the binary tree.
     * @return The maximum depth of the binary tree.
     */
    public int maxDepth(TreeNode root) {
        // Base case: If the root is null, the depth is 0.
        if (root == null) {
            return 0;
        }

        // Calculate the depth of the left subtree.
        int leftSubtreeDepth = maxDepth(root.left) + 1;

        // Calculate the depth of the right subtree.
        int rightSubtreeDepth = maxDepth(root.right) + 1;

        // Return the maximum depth between the left and right subtrees.
        return Math.max(leftSubtreeDepth, rightSubtreeDepth);
    }
}
