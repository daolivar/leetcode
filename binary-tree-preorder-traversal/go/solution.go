// preorderTraversal performs a pre-order traversal of a binary tree.
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
//   - An array containing the values of the tree nodes in pre-order sequence.
func preorderTraversal(root *TreeNode) []int {
	var preorder []int
	preorderTraversalHelper(root, &preorder)
	return preorder
}

// preorderTraversalHelper is a helper function for pre-order traversal.
//
// Time Complexity:  O(n) where n is the number of nodes in the subtree rooted at 'root'.
//
// Space Complexity: O(h) where h is the height of the subtree, representing the
// maximum depth of the function call stack.
//
// Parameters:
//   - root: The root of the subtree to traverse.
//   - preorder: A pointer to the array storing the pre-order traversal result.
//
// Returns:
//   - None (modifies the 'preorder' slice in place).
func preorderTraversalHelper(root *TreeNode, preorder *[]int) {
	// Check if the tree is not empty.
	if root != nil {
		// Add the value of the current node to the result list
		*preorder = append(*preorder, root.Val)

		// Recursively traverse the left subtree.
		preorderTraversalHelper(root.Left, preorder)

		// Recursively traverse the right subtree.
		preorderTraversalHelper(root.Right, preorder)
	}
}
