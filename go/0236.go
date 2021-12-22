package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// root == p || root == q. We need a way to signal that one of the nodes is on this sub-tree identified by root.
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil {
		if r != nil {
			// one on the left, the other one on the right
			return root
		}
		return l
	}
	return r
}
