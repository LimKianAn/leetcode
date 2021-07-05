package main

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	if target < nums[0] || target > nums[j] {
		return -1
	}
	for i <= j {
		mid := i + (j-i)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}
