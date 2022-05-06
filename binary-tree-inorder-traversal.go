package leetcode

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}
