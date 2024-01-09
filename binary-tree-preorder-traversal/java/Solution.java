/**
 * The Solution for performing preorder traversal on a binary tree.
 */
public class Solution {
    /**
     * Performs preorder traversal on a binary tree.
     * 
     * Time Complexity: O(n), where 'n' is the number of nodes in the binary tree.
     * This is because each node is visited exactly once.
     * 
     * Space Complexity: O(h), where 'h' is the height of the binary tree. This is
     * because the space required for the recursion call stack is proportional to
     * the height of the tree. In a balanced binary tree, the height is log(n),
     * where n is the number of nodes. In the worst case (skewed tree), the height
     * can be n, resulting in a space complexity of O(n).
     * 
     * @param root The root of the binary tree.
     * @return A list containing the preorder traversal of the binary tree.
     */
    public List<Integer> preorderTraversal(TreeNode root) {
        List<Integer> preorder = new ArrayList<>();
        preorderTraversal(root, preorder);
        return preorder;
    }

    /**
     * Helper method for preorder traversal on a binary tree.
     *
     * @param root     The current node in the traversal.
     * @param preorder The list to store the preorder traversal result.
     */
    private void preorderTraversal(TreeNode root, List<Integer> preorder) {
        // Check if the tree is not empty.
        if (root != null) {
            // Add the value of the current node to the result list.
            preorder.add(root.val);

            // Recursively traverse the left subtree.
            preorderTraversal(root.left, preorder);

            // Recursively traverse the right subtree.
            preorderTraversal(root.right, preorder);
        }
    }
}
