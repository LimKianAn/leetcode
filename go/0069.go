package main

// Time: O(log x)
// func mySqrt(x int) int {
// 	r := x
// 	for r*r > x {
// 		r = (r + x/r) >> 1
// 	}
// 	return r
// }

func mySqrt(x int) int {
	l, r := 0, x+1
	for l < r {
		m := l + (r-l)>>1
		if x < m*m {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}
