package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Sort(sortableIntervalsByEnd(intervals))

	n := len(intervals)
	start := 0
	overlappingGroupNum := 1 // only sorting by inerval-end guarantees all members in a group overlap with each other
	for i := 1; i < n; i++ {
		if intervals[start][1] > intervals[i][0] {
			continue
		}
		start = i
		overlappingGroupNum++
	}
	return n - overlappingGroupNum
}

type sortableIntervalsByEnd [][]int

func (s sortableIntervalsByEnd) Len() int {
	return len(s)
}

func (s sortableIntervalsByEnd) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s sortableIntervalsByEnd) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
