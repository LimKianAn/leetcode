// 2022.01.16

package main

func zigzagLevelOrder_b(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}
	evenQ, oddQ := []*TreeNode{root}, []*TreeNode{}
	level := 0
	for len(evenQ) > 0 || len(oddQ) > 0 {
		if len(ans) == level {
			ans = append(ans, []int{})
		}

		if level%2 == 0 {
			for len(evenQ) > 0 {
				front := evenQ[0]
				evenQ = evenQ[1:]

				ans[level] = append(ans[level], front.Val)

				if l := front.Left; l != nil {
					oddQ = append(oddQ, l)
				}
				if r := front.Right; r != nil {
					oddQ = append(oddQ, r)
				}
			}
		} else {
			for len(oddQ) > 0 {
				back := oddQ[len(oddQ)-1]
				oddQ = oddQ[:len(oddQ)-1]

				ans[level] = append(ans[level], back.Val)

				if r := back.Right; r != nil {
					evenQ = append(evenQ, r)
				}
				if l := back.Left; l != nil {
					evenQ = append(evenQ, l)
				}
			}

			for l, r := 0, len(evenQ)-1; l < r; {
				evenQ[l], evenQ[r] = evenQ[r], evenQ[l]
				l++
				r--
			}
		}
		level++
	}
	return ans
}
