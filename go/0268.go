package main

func missingNumber(nums []int) int { // len(nums) == n
	xor := 0
	for i, v := range nums { // [0,n-1]
		xor ^= i ^ v
	}
	return xor ^ len(nums) // n
}
