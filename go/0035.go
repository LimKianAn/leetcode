package main

func searchInsert(nums []int, target int) int {
	i := 0
	up := len(nums) - 1

	for i <= up {
		mid := i + (up-i)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			i = mid + 1
		} else {
			up = mid - 1
		}
	}
	return i
}
