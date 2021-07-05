package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	ll1 := leafs(root1)
	ll2 := leafs(root2)
	if len(ll1) != len(ll2) {
		return false
	}
	for i := range ll1 {
		if ll1[i] != ll2[i] {
			return false
		}
	}
	return true
}

func leafs(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	return append(leafs(root.Left), leafs(root.Right)...)
}
