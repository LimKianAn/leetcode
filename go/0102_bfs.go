package main

// bfs
func levelOrder0102(root *TreeNode) (allVals [][]int) {
	if root == nil {
		return
	}

	for queue := []*TreeNode{root}; len(queue) > 0; {
		size := len(queue)
		vals := []int{}
		for i := 0; i < size; i++ {
			if node := queue[i].Left; node != nil {
				queue = append(queue, node) // enqueues the left child
			}

			if node := queue[i].Right; node != nil {
				queue = append(queue, node) // enqueues the right child
			}

			vals = append(vals, queue[i].Val)
		}
		allVals = append(allVals, vals)
		queue = queue[size:] // dequeues the processed one
	}

	return
}
