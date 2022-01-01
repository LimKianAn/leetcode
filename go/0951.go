package main

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	ll := flipEquiv(root1.Left, root2.Left)
	rr := flipEquiv(root1.Right, root2.Right)
	lr := flipEquiv(root1.Left, root2.Right)
	rl := flipEquiv(root1.Right, root2.Left)
	return ll && rr || lr && rl
}
