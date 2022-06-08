package leetcode

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	postorder(root, &result)
	return result
}

func postorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, res)
	postorder(root.Right, res)
	*res = append(*res, root.Val)
}
