// https://www.youtube.com/watch?v=3mY5W0yojtA
package main

// func climbStairs(n int) int { // recursion
// 	if n <= 1 {
// 		return 1
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

// func climbStairs(n int) int { // recursion with memo
// 	memo := make([]int, n+1)
// 	var f func(n int) int
// 	f = func(n int) int {
// 		if n <= 1 {
// 			return 1
// 		}
// 		if memo[n] == 0 {
// 			memo[n] = f(n-1) + f(n-2)
// 		}
// 		return memo[n]
// 	}
// 	return f(n)
// }

// func climbStairs(n int) int { // dp with memo array
// 	dp := make([]int, n+1)
// 	dp[0], dp[1] = 1, 1
// 	for i := 2; i <= n; i++ {
// 		dp[i] = dp[i-1] + dp[i-2]
// 	}
// 	return dp[n] // last element
// }

func climbStairs(n int) int { // dp with 2 var
	dp0, dp1 := 1, 1
	for i := 2; i <= n; i++ {
		dp0, dp1 = dp1, dp0+dp1
	}
	return dp1
}
