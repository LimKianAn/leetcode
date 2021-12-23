package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 { // odd number
		return false
	}

	sum = sum >> 1                   // divided by 2
	largestSum := make([]int, sum+1) // easier to reason about
	for _, num := range nums {
		for sm := sum; sm >= num; sm-- {
			// largestSum[sm] is the largest possible sum which is less than or equal to sm
			// and consists of a subarray.
			// Either num can't be added to the previous largestSum[sm] (it would exceed sm)
			// or it can only be added to the largestSum[sm-num]
			largestSum[sm] = max(largestSum[sm], largestSum[sm-num]+num)
		}
	}
	return largestSum[sum] == sum
}

// func main() {
// 	log.Println(canPartition([]int{8, 3, 5}))
// }
