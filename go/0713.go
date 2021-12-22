package main

func numSubarrayProductLessThanK(nums []int, k int) (n int) {
	if k <= 1 {
		return 0
	}

	prod := 1
	start := 0
	for end, val := range nums {
		for prod *= val; prod >= k; {
			prod /= nums[start]
			start++
		}
		n += end - start + 1 // e.g. start = 0, end = 2, the subarrays would be [0, 1, 2], [1,2], [2]
	}
	return
}
