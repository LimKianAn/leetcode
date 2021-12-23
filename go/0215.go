package main

// func findKthLargest(nums []int, k int) int {
// 	sort.Ints(nums)
// 	return nums[len(nums)-k]
// }

// O(n log n)
func findKthLargest(nums []int, k int) int {
	// heapify
	for i := len(nums)/2 - 1; i >= 0; i-- { // len(nums)/2 -1, the last parent
		siftDown(nums, i, len(nums))
	}

	for i := len(nums) - 1; i >= len(nums)-k; i-- {
		// Insert the number at the end (bottom right), e.g. 3rd, 2nd, 1st
		nums[0], nums[i] = nums[i], nums[0]
		// Sift down with the greatest numbers at the end intact
		siftDown(nums, 0, i)
	}
	return nums[len(nums)-k]
}

// O(log n)
// Assume the children form a valid heap
// upperBound is exclusive ")"
func siftDown(nums []int, parent, upperBound int) {
	for child := 2*parent + 1; child < upperBound; child = 2*parent + 1 { // as long as the left child exists
		if child+1 < upperBound && nums[child] < nums[child+1] {
			child++ // the greater among the two children
		}
		if nums[parent] > nums[child] {
			break // the children downward assumed to be valid
		}
		nums[parent], nums[child] = nums[child], nums[parent]
		parent = child // child becomes the new parent
	}
}

// func main() {
// 	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
// }
