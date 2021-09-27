package main

import (
	"sort"
)

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			memo := i // index of the smallest number in nums[i:] so that nums[i-1] < the number < nums[i]
			for j := i; j < len(nums); j++ {
				if nums[i-1] < nums[j] && nums[j] < nums[memo] {
					memo = j
				}
			}

			swap(nums, i-1, memo)
			sort.Ints(nums[i:])
			return
		}
	}

	sort.Ints(nums)
	return
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
