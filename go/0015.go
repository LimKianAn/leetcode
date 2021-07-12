package main

import "sort"

func threeSum(nums []int) (a [][]int) { // a := answer
	sort.Ints(nums)
	len := len(nums)
outer:
	for i := 1; i < len-1; i++ { // in the range (first, last)
		l, r := 0, len-1 // left, right
		if i > 1 && nums[i] == nums[i-1] {
			l = i - 1
		}
		for l < i && r > i {
			if nums[l] > 0 {
				continue outer // sum of positive ints != 0
			}
			if l > 0 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < len-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			if sum := nums[l] + nums[i] + nums[r]; sum == 0 {
				a = append(a, []int{nums[l], nums[i], nums[r]})
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return
}
