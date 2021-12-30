package main

func firstBadVersionBinarySearch(n int) int { // 1-based array works as well
	l, r := 1, n // [l,r) r should be n+1, but to avoid overflow we use n. If not found in [1,n-1], return n, which must be the answer.
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
