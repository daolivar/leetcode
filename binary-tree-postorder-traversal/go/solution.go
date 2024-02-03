// postorderTraversal performs a post-order traversal of a binary tree.
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
//   - An array containing the values of the tree nodes in post-order sequence.
func postorderTraversal(root *TreeNode) []int {
	var postorder []int
	postorderTraversalHelper(root, &postorder)
	return postorder
}

// postorderTraversalHelper is a helper function for post-order traversal.
//
// Time Complexity:  O(n) where n is the number of nodes in the subtree rooted at 'root'.
//
// Space Complexity: O(h) where h is the height of the subtree, representing the
// maximum depth of the function call stack.
//
// Parameters:
//   - root: The root of the subtree to traverse.
//   - postorder: A pointer to the array storing the post-order traversal result.
//
// Returns:
//   - None (modifies the 'postorder' slice in place).
func postorderTraversalHelper(root *TreeNode, postorder *[]int) {
	// Check if the tree is not empty.
	if root == nil {
		return
	}

	// Recursively traverse the left subtree.
	postorderTraversalHelper(root.Left, postorder)

	// Recursively traverse the right subtree.
	postorderTraversalHelper(root.Right, postorder)

	// Add the value of the current node to the result list.
	*postorder = append(*postorder, root.Val)
}
