package main

const duplicates = 2

func removeDuplicates(nums []int) int {
	slow := 0
	for fast, v := range nums {
		if fast < duplicates || v != nums[slow-duplicates] {
			nums[slow] = v
			slow++
		}
	}
	return slow
}
