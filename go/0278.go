package main

import "sort"

func firstBadVersion(n int) int { // 1-based array works as well
	return sort.Search(n, func(v int) bool { return isBadVersion(v) })
}

func isBadVersion(v int) bool { // dummy
	return true
}
