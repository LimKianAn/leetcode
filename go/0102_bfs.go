// 2022.01.16

package main

func levelOrder_b(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}
	for q := []*TreeNode{root}; len(q) > 0; {
		qLen := len(q)
		vals := []int{}
		for i := 0; i < qLen; i++ {
			vals = append(vals, q[i].Val)

			if l := q[i].Left; l != nil {
				q = append(q, l) // enqueues the left child
			}
			if r := q[i].Right; r != nil {
				q = append(q, r) // enqueues the right child
			}
		}
		ans = append(ans, vals)
		q = q[qLen:] // dequeues the processed one
	}
	return ans
}
