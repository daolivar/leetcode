/**
 * Solution for performing level order traversal on a binary tree.
 */
public class Solution {
    /**
     * Performs level order traversal on a binary tree and returns the result.
     * 
     * Time Complexity: O(n), where 'n' is the number of nodes in the binary tree.
     * The algorithm traverses each node once.
     * 
     * Space Complexity: O(m), where 'm' is the maximum number of nodes at any level
     * in the binary tree. In the worst case, all nodes at the last level would be
     * in the queue at the same time.
     * 
     * @param root The root of the binary tree.
     * @return List of lists representing the level order traversal.
     */
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> levelorder = new ArrayList<>();
        levelOrderTraversal(root, levelorder);
        return levelorder;
    }

    /**
     * Helper method for performing level order traversal on a binary tree.
     *
     * @param root       The root of the binary tree.
     * @param levelorder List to store the level order traversal.
     */
    public void levelOrderTraversal(TreeNode root, List<List<Integer>> levelorder) {
        // Check if the tree is not empty
        if (root != null) {
            // Create a queue for BFS traversal
            Queue<TreeNode> queue = new LinkedList<>();
            queue.add(root);

            // Iterate through the tree levels
            while (!queue.isEmpty()) {
                // Get the number of nodes at the current level
                int nodesAtCurrLevel = queue.size();

                // List to store values at the current level
                List<Integer> level = new ArrayList<>();

                // Process each node at the current level
                for (int i = 0; i < nodesAtCurrLevel; i++) {
                    // Dequeue the front node
                    TreeNode top = queue.poll();

                    // Add the value of the current node to the level list
                    level.add(top.val);

                    // Enqueue the left child if it exists
                    if (top.left != null) {
                        queue.add(top.left);
                    }

                    // Enqueue the right child if it exists
                    if (top.right != null) {
                        queue.add(top.right);
                    }
                }

                // Add the current level list to the final result
                levelorder.add(level);
            }
        }
    }
}
