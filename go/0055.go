package main

func canJump(nums []int) bool {
	rightmostIndex := 0
	for i, jump := range nums {
		if rightmostIndex < i {
			return false
		}
		rightmostIndex = max(rightmostIndex, i+jump)
	}
	return true
}
