package main

func sortColors(nums []int) {
	i0, i1 := 0, 0 // i0 := index for zero
	for i, v := range nums {
		nums[i] = 2
		if v <= 1 {
			nums[i1] = 1
			i1++
		}
		if v == 0 {
			nums[i0] = 0
			i0++
		}
	}
}
