package main

func search0033(nums []int, target int) int {
	var f func(l, r int) int
	f = func(l, r int) int { // l := left, r := right
		for l <= r {
			mid := l + (r-l)>>1
			if nums[mid] == target {
				return mid
			} else if target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return -1
	}

	for l, r := 0, len(nums)-1; l <= r; {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] { // "must be <="
			if nums[l] <= target && target < nums[mid] {
				return f(l, mid)
			}
			l = mid + 1
		} else {
			if nums[mid] < target && target <= nums[r] {
				return f(mid, r)
			}
			r = mid - 1
		}
	}
	return -1
}
