package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	tree := &TreeNode{Val: preorder[0]}
	for i := range inorder {
		if inorder[i] == preorder[0] {
			tree.Left = buildTree(preorder[1:i+1], inorder[:i]) // translating preorder's index to inorder's
			tree.Right = buildTree(preorder[i+1:], inorder[i+1:])
		}
	}
	return tree
}

/*
preorder root left right
   0
 1   2
inorder left root right
   1
 0   2
 */