package main

func singleNumber(nums []int) int {
	x := 0
	for i := range nums {
		x ^= nums[i]
	}
	return x
}
