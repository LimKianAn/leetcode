package main

func integerBreak(n int) int {
	memo := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			memo[i] = max(memo[i], max(j*(i-j), j*memo[i-j]))
		}
	}
	return memo[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
