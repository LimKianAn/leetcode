package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		leafOfFlattened(root.Left).Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}

func leafOfFlattened(flattened *TreeNode) *TreeNode {
	for flattened.Right != nil {
		flattened = flattened.Right
	}
	return flattened
}
