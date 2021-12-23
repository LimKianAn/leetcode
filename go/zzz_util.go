// https://leetcode.com/discuss/interview-question/algorithms/233608/Uber-or-Maximum-drivers-online/320185

package main

const (
	start = 0
	end   = 1
)

var directions = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

type (
	ListNode struct {
		Val  int
		Next *ListNode
	}

	TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
