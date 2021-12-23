package main

func postorderTraversal(root *TreeNode) (a []int) {
	var f func(*TreeNode)
	f = func(root *TreeNode) {
		if root != nil {
			f(root.Left)
			f(root.Right)
			a = append(a, root.Val)
		}
	}

	f(root)
	return
}
