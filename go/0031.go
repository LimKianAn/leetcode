package main

import (
	"sort"
)

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			memo := i
			for j := i; j < len(nums); j++ {
				if nums[i-1] < nums[j] && nums[j] < nums[memo] {
					memo = j // making memo smaller and smaller
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

// 1,2,4,3
//     i
//       ^ memo (in the end, we found a smallest nums[memo])
