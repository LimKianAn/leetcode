package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func levelOrder(root *TreeNode) (a [][]int) { // a := answer
	if root == nil {
		return
	}

	for queue := []*TreeNode{root}; len(queue) > 0; {
		tmp := []int{}
		l := len(queue)
		for i := 0; i < l; i++ {
			if n := queue[i].Left; n != nil {
				queue = append(queue, n)
			}
			if n := queue[i].Right; n != nil {
				queue = append(queue, n)
			}
			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[l:]
		a = append(a, tmp)
	}
	return
}

// // dfs
// func levelOrder(root *TreeNode) (a [][]int) { // a := answer
// 	f := func(n *TreeNode, level int) {}
// 	f = func(n *TreeNode, level int) { // n := node
// 		if n == nil {
// 			return
// 		}

// 		if level == len(a) { // first one on this level
// 			a = append(a, []int{n.Val})
// 		} else {
// 			a[level] = append(a[level], n.Val)
// 		}

// 		f(n.Left, level+1)
// 		f(n.Right, level+1)
// 	}
// 	f(root, 0)
// 	return
// }
