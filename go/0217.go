package main

func containsDuplicate(nums []int) bool {
	record := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := record[v]; ok {
			return true
		}
		record[v] = struct{}{}
	}
	return false
}
