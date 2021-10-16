package main

import "sort"

func firstBadVersion(n int) int {
	return sort.Search(n, func(v int) bool { return isBadVersion(v) })
}

func isBadVersion(v int) bool {
	return true
}
