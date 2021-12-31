package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	if p == nil {
// 		return q == nil
// 	}
// 	if q == nil {
// 		return p == nil
// 	}
// 	if p.Val != q.Val {
// 		return false
// 	}
// 	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
// }
