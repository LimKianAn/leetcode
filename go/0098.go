package main

import "math"

func isValidBST(root *TreeNode) bool {
	return isValidBSTAgainstBounds(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTAgainstBounds(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper { // equality is also false
		return false
	}
	return isValidBSTAgainstBounds(root.Left, lower, root.Val) && isValidBSTAgainstBounds(root.Right, root.Val, upper)
}
