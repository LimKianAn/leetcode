package main

func maxArea(height []int) int {
	max := 0
	for l, r := 0, len(height)-1; l < r; { // left, right
		h := 0 // height
		w := r - l // width
		if height[l] < height[r] {
			h = height[l]
			l++
		} else {
			h = height[r]
			r--
		}
		if tmp := w * h; tmp > max {
			max = tmp
		}
	}
	return max
}
