package main

func maxProduct(nums []int) (product int) {
	maximum, minimum := 1, 1
	product = nums[0]

	for i := range nums {
		if nums[i] < 0 {
			maximum, minimum = minimum, maximum // prepares for the comparison below
		}

		maximum = max(nums[i], nums[i]*maximum)
		minimum = min(nums[i], nums[i]*minimum)

		if product < maximum {
			product = maximum
		}
	}

	return
}
