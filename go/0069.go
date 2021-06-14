package main

func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) >> 1
	}
	return r
}

// func mySqrt(x int) int {
// 	r := 0
// 	up := x // upper bound
// 	for r < up {
// 		mid := (r + up + 1) >> 1
// 		if mid*mid > x {
// 			up = mid - 1
// 		} else {
// 			r = mid
// 		}
// 	}
// 	return r
// }
