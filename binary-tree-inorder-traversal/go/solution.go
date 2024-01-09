// inorderTraversal performs an in-order traversal of a binary tree.
//
// Time Complexity:  O(n) where n is the number of nodes in the tree.
//
// Space Complexity: O(h) where h is the height of the tree, representing the
// maximum depth of the function call stack.
//
// Parameters:
//   - root: The root of the binary tree to traverse.
//
// Returns:
//   - An array containing the values of the tree nodes in in-order sequence.
func inorderTraversal(root *TreeNode) []int {
	var inorder []int
	inorderTraversalHelper(root, &inorder)
	return inorder
}

// inorderTraversalHelper is a helper function for in-order traversal.
//
// Time Complexity:  O(n) where n is the number of nodes in the subtree rooted at 'root'.
//
// Space Complexity: O(h) where h is the height of the subtree, representing the
// maximum depth of the function call stack.
//
// Parameters:
//   - root: The root of the subtree to traverse.
//   - inorder: A pointer to the array storing the in-order traversal result.
//
// Returns:
//   - None (modifies the 'inorder' slice in place).
func inorderTraversalHelper(root *TreeNode, inorder *[]int) {
	// Check if the tree is not empty.
	if root != nil {
		// Recursively traverse the left subtree.
		inorderTraversalHelper(root.Left, inorder)

		// Add the value of the current node to the result list
		*inorder = append(*inorder, root.Val)

		// Recursively traverse the left subtree.
		inorderTraversalHelper(root.Right, inorder)
	}
}
