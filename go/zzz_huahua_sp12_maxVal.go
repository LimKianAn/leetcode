package main

import "math"

func maxVal(root *treeNode) int {
	if root == nil {
		return math.MinInt
	}

	return max(root.val, max(maxVal(root.left), maxVal(root.right)))
}

func main() {
	print(maxVal(newBST([]int{5, 3, 1, 4, 7, 6})))
}
