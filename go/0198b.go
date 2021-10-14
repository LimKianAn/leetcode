package main

func rob(nums []int) int {
	even, odd := 0, 0
	for i := range nums {
		if i%2 == 0 {
			even = max(even+nums[i], odd)
		} else {
			odd = max(even, odd+nums[i])
		}
	}
	return max(even, odd)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
