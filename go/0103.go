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
			return // done
		}

		if level >= len(answer) {
			answer = append(answer, []int{}) // creating the slice for this level
		}

		if level%2 == 0 {
			answer[level] = append(answer[level], tree.Val) // left to right
		} else {
			answer[level] = append([]int{tree.Val}, answer[level]...) // right to left
		}

		dfs(tree.Left, level+1)
		dfs(tree.Right, level+1)
	}

	dfs(root, 0)
	return
}
