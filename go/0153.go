package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1 // l := left, r := right
	for l < r {
		m := l + (r-l)>>1 // m := middle
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
