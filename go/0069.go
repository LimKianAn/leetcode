package main

func mySqrt(x int) int {
	l, r := 0, x+1
	for l < r {
		m := l + (r-l)>>1
		if x/m < m { // avoids overflow
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}
