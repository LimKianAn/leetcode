package main

func containsNearbyDuplicate(nums []int, k int) bool {
	record := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := record[v]; ok {
			if i-j <= k {
				return true
			}
		}
		record[v] = i // always
	}
	return false
}
