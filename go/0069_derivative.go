package main

// Time: O(log x)
func mySqrtDerivative(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) >> 1
	}
	return r
}
