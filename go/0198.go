package main

func rob(nums []int) int {
	cur, pre := 0, 0
	for i := range nums {
		cur, pre = max(cur, pre+nums[i]), cur
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
