package main

func distanceK_recursion(root *TreeNode, target *TreeNode, k int) []int {
	ans := []int{}

	var addByDepth func(cur *TreeNode, k int) // adds nodes k steps from root to ans
	addByDepth = func(cur *TreeNode, k int) {
		if cur == nil {
			return
		}

		if k < 0 {
		} else if k == 0 {
			ans = append(ans, cur.Val)
		} else {
			addByDepth(cur.Left, k-1)
			addByDepth(cur.Right, k-1)
		}
	}

	var traversalAdd func(cur *TreeNode) (depth int)
	traversalAdd = func(cur *TreeNode) (depth int) {
		if cur == nil {
			return -1
		}

		if cur == target {
			addByDepth(target, k)
			return 0
		}

		l, r := traversalAdd(cur.Left), traversalAdd(cur.Right)
		if l >= 0 {
			if l == k-1 {
				ans = append(ans, cur.Val)
			}
			addByDepth(cur.Right, k-l-2) // l+2+?=k
			return l + 1
		}
		if r >= 0 {
			if r == k-1 {
				ans = append(ans, cur.Val)
			}
			addByDepth(cur.Left, k-r-2)
			return r + 1
		}

		return -1
	}
	traversalAdd(root)
	return ans
}
