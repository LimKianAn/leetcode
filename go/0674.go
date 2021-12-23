package main

func findLengthOfLCIS(nums []int) int {
	m := 1 // max
	l := 1 // current length
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			l++
			m = max(m, l)
		} else {
			l = 1 // reset
		}
	}
	return m
}
