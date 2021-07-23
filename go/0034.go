package main

func searchRange(nums []int, target int) []int {
	first := first(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}
	return []int{first, last(nums, target)}
}

func first(nums []int, target int) int {
	for l, r := 0, len(nums)-1; l <= r; { // l := left, r := right
		if mid := l + (r-l)>>1; nums[mid] == target {
			if mid == 0 {
				return 0
			}
			if nums[mid-1] != target {
				return mid
			}
			r = mid - 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func last(nums []int, target int) int {
	max := len(nums) - 1
	for l, r := 0, max; l <= r; { // l := left, r := right
		if mid := l + (r-l)>>1; nums[mid] == target {
			if mid == max {
				return max
			}
			if nums[mid+1] != target {
				return mid
			}
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
