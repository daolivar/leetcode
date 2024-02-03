/**
 * Solution to check if a binary tree is symmetric.
 */
public class Solution {
    /**
     * Checks if a given binary tree is symmetric.
     * 
     * Time Complexity: O(n) where n is the number of nodes in the binary tree.
     * This is because each node is visited once.
     * 
     * Space Complexity: O(h) where h is the height of the binary tree. In the
     * worst case, the recursion depth is equal to the height of the tree.
     * 
     * @param root The root of the binary tree.
     * @return True if the tree is symmetric, false otherwise.
     */
    public boolean isSymmetric(TreeNode root) {
        if (root == null) {
            return true;
        }

        return checkSymmetry(root.left, root.right);
    }

    /**
     * Recursively checks if two subtrees are symmetric.
     * 
     * @param left  The root of the left subtree.
     * @param right The root of the right subtree.
     * @return True if the subtrees are symmetric, false otherwise.
     */
    public boolean checkSymmetry(TreeNode left, TreeNode right) {
        // If both nodes are null, they are symmetric.
        if (left == null && right == null) {
            return true;
        }

        // If one node is null and the other is not, they are not symmetric.
        if (left == null || right == null) {
            return false;
        }

        // Check if the values are equal and the inner subtrees are symmetric.
        return (left.val == right.val) && checkSymmetry(left.left, right.right)
                && checkSymmetry(left.right, right.left);
    }
}
