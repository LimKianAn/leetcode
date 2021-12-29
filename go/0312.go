package main

func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, append(nums, 1)...) // adds the starat and the end

	localMax := make([][]int, n+2)
	for i := range localMax {
		localMax[i] = make([]int, n+2)
	}

	for length := 1; length <= n; length++ { // starts from the end game, i.e. 1 balloon, back to all ballons, i.e. n
		for i := 1; i <= n-length+1; i++ { // finds the upper bound by subtracting the length from n
			j := i + length - 1
			for middle := i; middle <= j; middle++ {
				localMax[i][j] = max(localMax[i][j], localMax[i][middle-1]+nums[i-1]*nums[middle]*nums[j+1]+localMax[middle+1][j])
			}
		}
	}
	return localMax[1][n]
}
