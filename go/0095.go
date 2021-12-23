package main

func generateTrees(n int) []*TreeNode {
	return generateTreesBetween(1, n)
}

func generateTreesBetween(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil} // nil is required here for the null node
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
