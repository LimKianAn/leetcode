package main

// func numTilings(n int) int {
// 	dp := make([][]uint64, n+1)
// 	for i := range dp {
// 		dp[i] = make([]uint64, 3)
// 	}
// 	dp[0][0] = 1
// 	dp[1][0] = 1
// 	mod := uint64(1000000007)
// 	for i := 2; i <= n; i++ {
// 		dp[i][0] = (dp[i-2][0] + dp[i-1][0] + dp[i-1][1] + dp[i-1][2]) % mod
// 		dp[i][1] = (dp[i-2][0] + dp[i-1][2]) % mod
// 		dp[i][2] = (dp[i-2][0] + dp[i-1][1]) % mod
// 	}
// 	return int(dp[n][0])
// }

func numTilings(n int) int {
	dp := make([][]uint64, n+1)
	for i := range dp {
		dp[i] = make([]uint64, 2)
	}
	dp[0][0] = 1
	dp[1][0] = 1
	mod := uint64(1000000007)
	for i := 2; i <= n; i++ {
		dp[i][0] = (dp[i-2][0] + dp[i-1][0] + 2*dp[i-1][1]) % mod // dp[i][1] and dp[i][2] are the same
		dp[i][1] = (dp[i-2][0] + dp[i-1][1]) % mod
	}
	return int(dp[n][0])
}
