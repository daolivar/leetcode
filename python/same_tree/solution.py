class Solution(object):
    def isSameTree(self, p, q):
        """
        Determine if two binary trees are the same.
        Two binary trees are considered the same if they are structurally identical
        and the nodes have the same value.

        :type p: Optional[TreeNode]  # Root of the first binary tree
        :type q: Optional[TreeNode]  # Root of the second binary tree
        :rtype: bool  # The function returns a boolean (True if the trees are identical, False otherwise)
        """

        # Base case: If both trees are empty, they are the same
        if not p and not q:
            return True

        # If one tree is empty and the other is not, or their values differ, they are not the same
        if not p or not q or p.val != q.val:
            return False

        # Recursively check if the left subtrees and right subtrees are the same
        return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
