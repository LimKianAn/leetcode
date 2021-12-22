package main

func sortColors(nums []int) {
	zero, zeroAndOne := 0, 0 // zero := counter for zero, zeroAndOne := counter for zero and one
	for i, v := range nums {
		nums[i] = 2

		if v <= 1 {
			nums[zeroAndOne] = 1
			zeroAndOne++
		}

		if v == 0 {
			nums[zero] = 0
			zero++
		}
	}
}
