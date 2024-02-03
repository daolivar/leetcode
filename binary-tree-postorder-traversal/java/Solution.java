/**
 * Solution for performing postorder traversal on a binary tree.
 * The postorder traversal visits the left subtree, then the right subtree, and
 * finally the root.
 */
public class Solution {
    /**
     * Performs postorder traversal on the given binary tree and returns the result
     * as a list of integers.
     * 
     * Time Complexity: O(n) where n is the number of nodes in the binary tree.
     * Each node is visited exactly once during the traversal.
     * 
     * Space Complexity: O(n) where n is the height of the binary tree. The space
     * is used for the recursive call stack, and in the worst case, the height of
     * the call stack is equal to the height of the tree.
     * 
     * @param root The root of the binary tree.
     * @return A list of integers representing the postorder traversal of the binary
     *         tree.
     */
    public List<Integer> postorderTraversal(TreeNode root) {
        List<Integer> postorder = new ArrayList<>();
        postorderTraversal(root, postorder);
        return postorder;
    }

    /**
     * Recursive helper function to perform postorder traversal on the binary tree.
     *
     * @param root      The current node in the traversal.
     * @param postorder The list to store the result of postorder traversal.
     */
    private void postorderTraversal(TreeNode root, List<Integer> postorder) {
        // Check if the tree is not empty.
        if (root != null) {
            // Recursively traverse the left subtree.
            postorderTraversal(root.left, postorder);

            // Recursively traverse the right subtree.
            postorderTraversal(root.right, postorder);

            // Add the value of the current node to the result list.
            postorder.add(root.val);
        }
    }
}
