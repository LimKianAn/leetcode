package main

func search(nums []int, target int) int {
	for l, r := 0, len(nums)-1; l <= r; { // l := left, r := right
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[l] { // "must be >="
			if target >= nums[l] && target < nums[mid] {
				return searchSorted(nums, l, mid, target)
			}
			l = mid + 1
		} else {
			if target > nums[mid] && target <= nums[r] {
				return searchSorted(nums, mid, r, target)
			}
			r = mid - 1

		}
	}
	return -1
}

func searchSorted(nums []int, l, r, target int) int { // l := left index, r := right index, target := target number
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
