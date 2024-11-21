class Solution(object):
    def maxDepth(self, root):
        """
        Calculate the maximum depth of a binary tree.

        :type root: Optional[TreeNode]  # The root node of the binary tree
        :rtype: int  # The function returns an integer representing the maximum depth of the tree
        """

        # Base case: if the current node is None (empty tree), the depth is 0
        if not root:
            return 0

        # Recursively calculate the maximum depth of the left and right subtrees
        # Take the maximum of both depths and add 1 (for the current node/root)
        return max(self.maxDepth(root.left), self.maxDepth(root.right)) + 1
