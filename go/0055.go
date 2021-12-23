package main

func canJump(nums []int) bool {
	rightmost := 0
	for i, v := range nums {
		if rightmost < i {
			return false
		}

		rightmost = max(rightmost, i+v)
	}
	return true
}
