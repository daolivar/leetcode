/**
 * The Solution class represents a solution for performing inorder traversal on
 * a binary tree.
 */
public class Solution {
    /**
     * Performs inorder traversal on a binary tree.
     * 
     * Time Complexity: O(n) where n is the number of nodes in the binary tree.
     * This is because each node is visited exactly once.
     * 
     * Space Complexity: O(h) where h is the height of the binary tree. This is
     * because the space required for the recursion call stack is proportional to
     * the height of the tree. In a balanced binary tree, the height is log(n),
     * where n is the number of nodes. In the worst case (skewed tree), the height
     * can be n, resulting in a space complexity of O(n).
     * 
     * @param root The root of the binary tree.
     * @return A list containing the inorder traversal of the binary tree.
     */
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> inorder = new ArrayList<>();
        inorderTraversal(root, inorder);
        return inorder;
    }

    /**
     * Helper method for inorder traversal on a binary tree.
     * 
     * @param root   The current node in the traversal.
     * @param result The list to store the inorder traversal result.
     */
    private void inorderTraversal(TreeNode root, List<Integer> inorder) {
        // Check if the tree is not empty.
        if (root != null) {
            // Recursively traverse the left subtree.
            inorderTraversal(root.left, inorder);

            // Add the value of the current node to the result list.
            inorder.add(root.val);

            // Recursively traverse the right subtree.
            inorderTraversal(root.right, inorder);
        }
    }
}
