package main

func missingNumber(nums []int) int {
	xor := 0
	for i, v := range nums {
		xor ^= i ^ v
	}
	return xor ^ len(nums)
}
