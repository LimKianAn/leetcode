package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	val := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{val}
	}
	return append(paths(val, root.Left), paths(val, root.Right)...)
}

func paths(val string, child *TreeNode) []string {
	paths := []string{}
	if child != nil {
		sub := binaryTreePaths(child)
		for i := range sub {
			paths = append(paths, val+"->"+sub[i])
		}
	}
	return paths
}
