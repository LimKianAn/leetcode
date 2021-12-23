package main

func preorderTraversal(root *TreeNode) []int {
	ii := &[]int{}
	f(root, ii)
	return *ii
}

func f(root *TreeNode, ii *[]int) {
	if root != nil {
		*ii = append(*ii, root.Val)
		f(root.Left, ii)
		f(root.Right, ii)
	}
}
