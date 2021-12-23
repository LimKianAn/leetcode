package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lHeight := height(root.Left)
	rHeight := height(root.Right)
	return isBalanced(root.Left) && isBalanced(root.Right) && abs(lHeight-rHeight) <= 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1 // plus the height of the root, which is one
}
