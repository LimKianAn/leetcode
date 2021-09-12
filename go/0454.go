package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	memo := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			memo[v1+v2]++
		}
	}

	count := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += memo[0-v3-v4]
		}
	}
	return count
}
