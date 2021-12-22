// https://leetcode.com/discuss/interview-question/algorithms/233608/Uber-or-Maximum-drivers-online/320185

package main

import (
	"log"
	"sort"
)

const (
	start = 0
	end   = 1
)

type sortable2DInts [][]int

func (x sortable2DInts) Len() int {
	return len(x)
}

func (x sortable2DInts) Less(i, j int) bool {
	return x[i][start] < x[j][start]
}

func (x sortable2DInts) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxNumOverlapped(schedules [][]int) int {
	if len(schedules) == 0 {
		return 0
	}

	sort.Sort(sortable2DInts(schedules))

	globalMax := 1
	for i := 0; i < len(schedules); i++ {
		localMax := 1
		for j := i + 1; j < len(schedules); j++ {
			if schedules[i][end] < schedules[j][start] {
				break
			}
			localMax++
		}
		globalMax = max(globalMax, localMax)
	}
	return globalMax
}

func main() {
	schedulesA := [][]int{{1, 8}, {5, 10}, {14, 20}}
	log.Println(maxNumOverlapped(schedulesA))

	schedulesB := [][]int{{1, 8}, {9, 10}, {14, 20}}
	log.Println(maxNumOverlapped(schedulesB))

	schedulesC := [][]int{{5, 10}, {14, 20}, {18, 28}, {18, 28}, {21, 23}}
	log.Println(maxNumOverlapped(schedulesC))

	schedulesD := [][]int{}
	log.Println(maxNumOverlapped(schedulesD))
}
