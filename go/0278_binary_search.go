package main

func firstBadVersionBinarySearch(n int) int { // 1-based array works as well
	l, r := 1, n+1
	for l < r {
		m := l + (r-l)>>1
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
