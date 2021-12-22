package main

const duplicates = 2

func removeDuplicates(nums []int) int {
	slow := 0 // the index of the last modified number + 1
	for fast, v := range nums {
		// fast < duplicates, just let the first two in
		// v != nums[slow-duplicates], the same value can not appear more than twice consecutively
		if fast < duplicates || v != nums[slow-duplicates] {
			nums[slow] = v
			slow++
		}
	}
	return slow
}
