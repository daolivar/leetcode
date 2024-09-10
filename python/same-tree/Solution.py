class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        """
        Check if two binary trees are the same.

        Parameters:
        p (Optional[TreeNode]): The root of the first binary tree.
        q (Optional[TreeNode]): The root of the second binary tree.

        Returns:
        bool: True if both trees are identical, False otherwise.
        """
        if not p and not q:
            return True
        if not p or not q or p.val != q.val:
            return False

        return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
