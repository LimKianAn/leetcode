package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	v := []int{} // values of nodes
	addValue(root, &v)
	return v
}

func addValue(root *TreeNode, v *[]int) {
	if root == nil {
		return
	}
	addValue(root.Left, v)
	*v = append(*v, root.Val)
	addValue(root.Right, v)
}
