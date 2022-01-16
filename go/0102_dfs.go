// 2022.01.16

package main

func levelOrder_d(root *TreeNode) [][]int {
	ans := [][]int{}
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level == len(ans) { // first one on this level
			ans = append(ans, []int{root.Val})
		} else {
			ans[level] = append(ans[level], root.Val)
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return ans
}
