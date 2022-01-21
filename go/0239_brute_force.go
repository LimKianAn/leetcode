// 2022.01.21

package main

func maxSlidingWindow_b(nums []int, k int) []int {
	ans := make([]int, 0, len(nums))
	for i := 0; i <= len(nums)-k; i++ {
		max := -1 << 63
		for _, num := range nums[i : i+k] {
			if num > max {
				max = num
			}
		}
		ans = append(ans, max)
	}
	return ans
}
