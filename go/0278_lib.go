package main

import "sort"

var isBadVersion func(v int) bool

func firstBadVersionLib(n int) int { // 1-based array works as well
	return sort.Search(n, isBadVersion)
}
