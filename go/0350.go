package main

func intersect(nums1 []int, nums2 []int) []int {
	ii := []int{}
	m := map[int]int{}
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if m[v] > 0 {
			ii = append(ii, v)
			m[v]--
		}
	}
	return ii
}
