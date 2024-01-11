// isSameTree checks if two binary trees are identical.
//
// Time Complexity: O(n) where n is the number of nodes in the tree, as each node is visited once.
//
// Space Complexity: O(h) where h is the height of the tree, representing the maximum depth of the function call stack.
//
// Parameters:
//   - p: The root of the first binary tree.
//   - q: The root of the second binary tree.
//
// Returns:
//   - true if the trees are identical, false otherwise.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Check if both trees are nil, indicating they are identical.
	if p == nil && q == nil {
		return true
	}

	// If one of the trees is empty while the other is not, or their values are
	// different, return false.
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	// Recursively check the left and right subtrees for equality.
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
