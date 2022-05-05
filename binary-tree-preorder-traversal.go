package leetcode

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	preOrder(root, &result)
	return result
}

func preOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}
