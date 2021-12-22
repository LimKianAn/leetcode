package main

func searchInsert(nums []int, target int) int {
	i := 0
	upper := len(nums) - 1

	for i <= upper {
		mid := i + (upper-i)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			upper = mid - 1
		}
	}
	return i
}
