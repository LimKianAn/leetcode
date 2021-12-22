package main

import "sort"

func threeSum(nums []int) (a [][]int) { // a := answer
	sort.Ints(nums)
	N := len(nums)
outer:
	for i := 1; i < N-1; i++ { // in the range (first, last)
		l, r := 0, N-1 // left, right
		if nums[i] == nums[i-1] {
			l = i - 1 // the possible solution with nums[i-1] has been evaluated
		}

		for l < i && i < r {
			if nums[l] > 0 {
				continue outer // sum of positive ints != 0
			}
			if l > 0 && nums[l-1] == nums[l] { // l might be zero
				l++
				continue
			}
			if r < N-1 && nums[r] == nums[r+1] { // r might be N-1
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
