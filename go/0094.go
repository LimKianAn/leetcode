package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func inorderTraversal(root *TreeNode) []int {
// 	v := []int{} // values of nodes
// 	addValue(root, &v)
// 	return v
// }

// func addValue(root *TreeNode, v *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	addValue(root.Left, v)
// 	*v = append(*v, root.Val)
// 	addValue(root.Right, v)
// }

func inorderTraversal(root *TreeNode) (ii []int) {
	if root == nil {
		return
	}

	now := root
	memo := []*TreeNode{}
	for now != nil || len(memo) > 0 {
		if now != nil {
			memo = append(memo, now)
			now = now.Left
		} else {
			// Pop the last one
			last := len(memo) - 1
			now = memo[last]
			memo = memo[:last]

			ii = append(ii, now.Val)
			now = now.Right
		}
	}
	log.Print(ii)
	return ii
}

func main() {
	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{20, nil, nil}
	root.Right = &TreeNode{30, nil, nil}
	root.Left.Left = &TreeNode{40, nil, nil}
	root.Left.Right = &TreeNode{50, nil, nil}
	root.Right.Right = &TreeNode{60, nil, nil}
	root.Left.Left.Left = &TreeNode{70, nil, nil}
	inorderTraversal(root)
}
