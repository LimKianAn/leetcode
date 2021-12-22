package main

import (
	"log"
	"strconv"
)

func summaryRanges(nums []int) (ss []string) {
	for start := 0; start < len(nums); {
		s := strconv.Itoa(nums[start])
		last := lastOfRange(nums, start)
		if last != start {
			s += "->" + strconv.Itoa(nums[last])
		}
		ss = append(ss, s)
		start = last + 1 // after the gap
	}
	return
}

func lastOfRange(nums []int, start int) int {
	last := start
	for i := start + 1; i < len(nums) && nums[i] == nums[i-1]+1; i++ {
		last = i
	}
	return last
}

func main() {
	log.Print(lastOfRange([]int{0, 1, 3, 6, 7, 8}, 3))
}
