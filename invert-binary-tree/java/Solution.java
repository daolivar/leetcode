/**
 * Solution to invert a binary tree. To invert a binary tree means to swap the
 * left and right subtrees of every node in the tree.
 */
class Solution {
    /**
     * Inverts a binary tree.
     * 
     * This method recursively traverses the binary tree and swaps the left and
     * right subtrees at each node to create the inverted tree.
     * 
     * Time Complexity: O(n), where 'n' is the number of nodes in the binary tree.
     * This is because the function visits each node exactly once.
     * 
     * Space Complexity: O(h), where 'h' is the height of the binary tree. It's
     * important to note that the space complexity is determined by the maximum
     * depth of the recursive call stack, which corresponds to the height of the
     * tree. The number of recursive calls on the stack are proportional to the
     * height of the tree. In the worst case, where the tree is completely
     * unbalanced (skewed), the height 'h' could be equal to 'n', resulting in O(n)
     * space complexity. In a balanced tree, the height 'h' is O(log n).
     * 
     * @param root The root of the binary tree to be inverted.
     * @return The root of the inverted binary tree.
     */
    public TreeNode invertTree(TreeNode root) {
        // Base case: If the root is null, return null.
        if (root == null) {
            return null;
        }

        // Swap the left and right subtrees of the current node.
        TreeNode temp = root.left;
        root.left = root.right;
        root.right = temp;

        // Recursively invert the left and right subtrees.
        invertTree(root.left);
        invertTree(root.right);

        // Return the root of the inverted binary tree.
        return root;
    }
}
