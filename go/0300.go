package main

import "sort"

// O(n log n)
func lengthOfLIS(nums []int) int {
	memo := []int{}
	for _, n := range nums {
		// O(log n)
		if i := sort.SearchInts(memo, n); i == len(memo) {
			memo = append(memo, n)
		} else {
			memo[i] = n
		}
	}
	return len(memo)
}
