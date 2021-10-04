package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTAgainstBounds(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTAgainstBounds(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return isValidBSTAgainstBounds(root.Left, lower, root.Val) && isValidBSTAgainstBounds(root.Right, root.Val, upper)
}