package main

func minSubArrayLen(target int, nums []int) int {
	virtualL := len(nums) + 1                      // virtualL := virtual length
	minL := virtualL                               // minL := minimum length
	for l, r, sum := 0, 0, 0; r < len(nums); r++ { // l, r := left, right
		sum += nums[r]
		for ; sum >= target && l < len(nums); l++ {
			minL = min(minL, r-l+1)
			sum -= nums[l]
		}
	}
	if minL == virtualL {
		return 0
	}
	return minL
}
