package main

func rob0198(nums []int) int {
	cur, pre := 0, 0
	for i := range nums {
		cur, pre = max(cur, pre+nums[i]), cur // the current at the end is the old current before max() is executed
	}
	return cur
}
