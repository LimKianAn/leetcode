package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) (val int) {
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrder(root.Left)
		if k--; k == 0 {
			val = root.Val
			return
		}
		inOrder(root.Right)
	}

	inOrder(root)
	return
}
