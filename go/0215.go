package main

// func findKthLargest(nums []int, k int) int {
// 	sort.Ints(nums)
// 	return nums[len(nums)-k]
// }

// O(n log n)
func findKthLargest(nums []int, k int) int {
	// heapify
	for i := (len(nums) - 1) / 2; i >= 0; i-- {
		siftDown(nums, i, len(nums))
	}
	for i := len(nums) - 1; i >= len(nums)-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		siftDown(nums, 0, i)
	}
	return nums[len(nums)-k]
}

// O(log n)
func siftDown(nums []int, i, max int) {
	// As long as the left child is less than max
	for 2*i+1 < max {
		j := 2*i + 1
		if j+1 < max && nums[j+1] > nums[j] {
			j++
		}
		if nums[i] > nums[j] {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
	}
}

func main() {
	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
}
