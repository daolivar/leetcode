package leetcode

func isSymmetric(root *TreeNode) bool {
	return mirrorNodes(root.Left, root.Right)
}

func mirrorNodes(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return mirrorNodes(left.Left, right.Right) && mirrorNodes(left.Right, right.Left)
}
