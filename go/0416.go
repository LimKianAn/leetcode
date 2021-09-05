package main

import "log"

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 {
		return false
	}

	sum = sum >> 1
	memo := make([]int, sum+1) // easier to reason about
	for _, num := range nums {
		for sm := sum; sm >= num; sm-- {
			// memo[sm] := current max sum which is less than or equal to sm
			// and consists of the current subarray.
			// Either num can't be added to the previous memo[sm]
			// or it can only be added to the memo[sm-num]
			memo[sm] = max(memo[sm], memo[sm-num]+num)
		}
	}
	return memo[sum] == sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	log.Println(canPartition([]int{8, 3, 5}))
}
