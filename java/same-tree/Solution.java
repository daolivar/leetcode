/**
 * Solution class for LeetCode problem 100. Same Tree
 */
public class Solution {

    /**
     * Determines if two binary trees are identical.
     * 
     * Two binary trees are considered the same if they are structurally identical,
     * and the nodes have the same value.
     * 
     * @param p the root node of the first tree
     * @param q the root node of the second tree
     * @return true if the trees are identical, false otherwise
     */
    public boolean isSameTree(TreeNode p, TreeNode q) {
        // If both nodes are null, the trees are identical up to this point
        if (p == null && q == null) {
            return true;
        }
        // If one of the nodes is null, the trees are not identical
        if (p == null || q == null) {
            return false;
        }
        // If the values of the nodes are different, the trees are not identical
        if (p.val != q.val) {
            return false;
        }
        // Recursively check the left and right subtrees
        return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
    }
}
