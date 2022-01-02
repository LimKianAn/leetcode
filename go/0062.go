package main

// func uniquePaths(m int, n int) int { // recursion
// 	if m <= 0 || n <= 0 {
// 		return 0
// 	}
// 	if m == 1 && n == 1 {
// 		return 1
// 	}
// 	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
// }

// func uniquePaths(m int, n int) int { // recursion with memo array
// 	memo := new2DArray(m+1, n+1)
// 	var f func(m, n int) int
// 	f = func(m, n int) int {
// 		if m <= 0 || n <= 0 {
// 			return 0
// 		}
// 		if m == 1 && n == 1 {
// 			return 1
// 		}
// 		if memo[m][n] == 0 {
// 			memo[m][n] = f(m-1, n) + f(m, n-1)
// 		}
// 		return memo[m][n]
// 	}
// 	return f(m, n)
// }

func uniquePaths(m int, n int) int { // dp
	dp := new2DArray(m+1, n+1)
	dp[1][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m][n]
}
