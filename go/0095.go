package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return generateTreesBetween(1, n)
}

func generateTreesBetween(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	trees := []*TreeNode{}
	for i := start; i <= end; i++ {
		lTrees := generateTreesBetween(start, i-1)
		rTrees := generateTreesBetween(i+1, end)
		for _, lt := range lTrees {
			for _, rt := range rTrees {
				tree := &TreeNode{Val: i, Left: lt, Right: rt}
				trees = append(trees, tree)
			}
		}
	}
	return trees
}
