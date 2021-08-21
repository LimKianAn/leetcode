package main

import (
	"sort"
	"strconv"
)

type intSlice []int

func (ii intSlice) Len() int { return len(ii) }
func (ii intSlice) Less(i, j int) bool {
	return strconv.Itoa(ii[i])+strconv.Itoa(ii[j]) > strconv.Itoa(ii[j])+strconv.Itoa(ii[i])
}
func (ii intSlice) Swap(i, j int) { ii[i], ii[j] = ii[j], ii[i] }

func largestNumber(nums []int) string {
	sort.Sort(intSlice(nums))
	if nums[0] == 0 {
		return "0"
	}

	s := ""
	for _, v := range nums {
		s += strconv.Itoa(v)
	}
	return s
}

func main() {
	ii := []int{0, 9}
	print(largestNumber(ii))
}
