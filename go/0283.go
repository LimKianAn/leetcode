package main

func moveZeroes(nums []int) {
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}
