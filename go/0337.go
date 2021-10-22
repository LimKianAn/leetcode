package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var f func(*TreeNode) (int, int)
	f = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}

		l0, l1 := f(root.Left)
		r0, r1 := f(root.Right)

		now0 := max(l0, l1) + max(r0, r1) // the current node isn't robbed
		now1 := root.Val + l0 + r0        // the current node is robbed
		return now0, now1
	}

	max0, max1 := f(root)
	return max(max0, max1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
