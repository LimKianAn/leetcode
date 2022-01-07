package main

func maxArea(height []int) int {
	ans := 0
	for l, r := 0, len(height)-1; l < r; { // left, right
		h := 0     // height
		w := r - l // width
		if height[l] < height[r] {
			h = height[l]
			l++
		} else {
			h = height[r]
			r--
		}
		ans = max(ans, w*h)
	}
	return ans
}
