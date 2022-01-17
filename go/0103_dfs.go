// 2022.01.16

package main

func zigzagLevelOrder_d(root *TreeNode) [][]int {
	ans := [][]int{}
	var dfs func(tree *TreeNode, level int)
	dfs = func(tree *TreeNode, level int) {
		if tree == nil {
			return // done
		}

		if level == len(ans) {
			ans = append(ans, []int{}) // creates the slice for this level
		}

		if level%2 == 0 {
			ans[level] = append(ans[level], tree.Val) // left to right
		} else {
			ans[level] = append([]int{tree.Val}, ans[level]...) // right to left
		}

		dfs(tree.Left, level+1)
		dfs(tree.Right, level+1)
	}
	dfs(root, 0)
	return ans
}
