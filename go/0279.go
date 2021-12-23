package main

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i // consists of 1,1,...
		for j := 2; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1) // +1 because of j itself
		}
	}
	return dp[n]
}
