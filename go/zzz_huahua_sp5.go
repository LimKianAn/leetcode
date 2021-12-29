package main

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums) // [l,r) !
	for l < r {
		m := l + (r-l)>>1
		if target <= nums[m] { // at equality, the range keeps shifting leftwards
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(nums []int, target int) int {
	l, r := 0, len(nums) // [l,r) !
	for l < r {
		m := l + (r-l)>>1
		if target < nums[m] {
			r = m
		} else { // at equality, the range keeps shifting rightwards
			l = m + 1
		}
	}
	return l
}

// func main() {
// 	nums := []int{1, 2, 2, 2, 4, 4, 5}
// 	print(lowerBound(nums, 0) == 0)
// 	print(lowerBound(nums, 2) == 1)
// 	print(lowerBound(nums, 3) == 4)
// 	print(upperBound(nums, 2) == 4)
// 	print(upperBound(nums, 4) == 6)
// 	print(upperBound(nums, 5) == 7)
// }
