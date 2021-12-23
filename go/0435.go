package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	l := len(intervals)
	if l == 0 {
		return 0
	}

	s := sortable(intervals)
	sort.Sort(&s)

	notOverlapped := 1
	for last, now := 0, 1; now < l; now++ {
		if s[last][end] <= s[now][start] {
			notOverlapped++
			last = now
		}
	}
	return l - notOverlapped
}

type sortable [][]int

func (s *sortable) Len() int {
	return len(*s)
}

func (s *sortable) Less(i, j int) bool {
	if (*s)[i][end] != (*s)[j][end] {
		return (*s)[i][end] < (*s)[j][end]
	}
	return false
}

func (s *sortable) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
