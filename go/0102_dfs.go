package main

// dfs
func levelOrder(root *TreeNode) (a [][]int) { // a := answer
	var f func(n *TreeNode, level int)
	f = func(n *TreeNode, level int) { // n := node
		if n == nil {
			return
		}

		if level == len(a) { // first one on this level
			a = append(a, []int{n.Val})
		} else {
			a[level] = append(a[level], n.Val)
		}

		f(n.Left, level+1)
		f(n.Right, level+1)
	}
	f(root, 0)
	return
}
