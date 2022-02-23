// 2022.02.22

package main

import "sort"

func majorityElement_bm(nums []int) int { // Boyer-Moore, t O(n), sp O(1)
	count := 0
	candidate := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num // restarts
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func majorityElement_ht(nums []int) int { // hash table, t O(n), sp O(1)
	n := len(nums)
	count := map[int]int{}
	for _, num := range nums {
		count[num]++
		if count[num] > n/2 {
			return num
		}
	}
	return -1
}

func majorityElement_s(nums []int) int { // sort, t O(nlogn), sp O(1)
	sort.Ints(nums)
	return nums[len(nums)/2]
}
