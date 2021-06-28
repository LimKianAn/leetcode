package main

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]struct{}{}
	ii := []int{}
	for _, v := range nums1 {
		m[v] = struct{}{} // unique keys
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			ii = append(ii, v)
			delete(m, v) // already added
		}
	}
	return ii
}
