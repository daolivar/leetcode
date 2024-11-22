class Solution(object):
    def invertTree(self, root):
        """
        Invert a binary tree by swapping the left and right children of every node.

        :type root: Optional[TreeNode]  # The root node of the binary tree
        :rtype: Optional[TreeNode]  # The function returns the root of the inverted binary tree
        """

        # Base case: If the tree is empty (root is None), return None
        if not root:
            return root

        # Swap the left and right children of the current node
        root.left, root.right = root.right, root.left

        # Recursively invert the left subtree
        self.invertTree(root.left)
        # Recursively invert the right subtree
        self.invertTree(root.right)

        # Return the root node, which now represents the inverted tree
        return root
