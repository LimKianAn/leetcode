package main

type treeNode struct {
	val         int
	left, right *treeNode
}

func newBST(nums []int) *treeNode {
	var root *treeNode // root := (*treeNode)(nil) would also do, but not &treeNode{}, i.e. a node with val 0
	for _, num := range nums {
		root = insert(root, num)
	}
	return root
}

func insert(root *treeNode, val int) *treeNode {
	if root == nil {
		return &treeNode{val: val}
	}

	if val <= root.val {
		root.left = insert(root.left, val) // new pointer in root.left
	} else {
		root.right = insert(root.right, val) // new pointer in root.right
	}

	return root
}

func inorderPrint(root *treeNode) {
	if root == nil {
		return
	}

	inorderPrint(root.left)
	print(root.val)
	inorderPrint(root.right)
}

// func main() {
// 	inorderPrint(newBST([]int{5, 3, 1, 4, 7, 6}))
// }
