package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	memo := make([]int, n)
	memo[0] = 1
	memo[1] = 2
	for i := 2; i < n; i++ {
		memo[i] = memo[i-2] + memo[i-1]
	}
	return memo[n-1] // last element
}
