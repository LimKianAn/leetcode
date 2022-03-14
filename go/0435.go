package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	overlappingGroupNum := 1 // only sorting by inerval-end guarantees all members in a group overlap with each other
	start := 0
	for i := 1; i < n; i++ {
		if intervals[start][1] > intervals[i][0] {
			continue
		}
		start = i
		overlappingGroupNum++
	}
	return n - overlappingGroupNum
}
