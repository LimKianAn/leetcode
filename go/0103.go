package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (answer [][]int) {
	var dfs func(*TreeNode, int)
	dfs = func(tree *TreeNode, level int) {
		if tree == nil {
			return
		}

		if level >= len(answer) {
			answer = append(answer, []int{})
		}

		if level%2 == 0 {
			answer[level] = append(answer[level], tree.Val)
		} else {
			answer[level] = append([]int{tree.Val}, answer[level]...)
		}

		dfs(tree.Left, level+1)
		dfs(tree.Right, level+1)
	}

	dfs(root, 0)
	return
}
