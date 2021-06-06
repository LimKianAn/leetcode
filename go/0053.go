package main

func maxSubArray(nums []int) int {
	g := nums[0] // global max
	l := 0       // local max
	for i := range nums {
		l += nums[i]
		if l > g {
			g = l
		}
		if l < 0 {
			l = 0
		}
	}
	return g
}

// func maxSubArray(nums []int) int {
// 	g := nums[0]                // global max
// 	l := make([]int, len(nums)) // local max array
// 	l[0] = g
// 	for i := 1; i < len(nums); i++ {
// 		if l[i-1] > 0 {
// 			l[i] = l[i-1] + nums[i]
// 		} else {
// 			l[i] = nums[i]
// 		}
// 		g = max(g, l[i])
// 	}
// 	return g
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
