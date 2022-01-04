package main

// func minFlipsMonoIncr(s string) int {
// 	n := len(s)
// 	l := make([]int, n)
// 	l[0] = int(s[0] - '0')
// 	for i := 1; i < n; i++ {
// 		l[i] = l[i-1] + int(s[i]-'0') // prefix sum
// 	}

// 	r := make([]int, n+1)
// 	r[n-1] = int('1' - s[n-1])
// 	for i := n - 2; i >= 0; i-- {
// 		r[i] = r[i+1] + int('1'-s[i]) // suffix sum
// 	}

// 	ans := r[0] // flips needed to make all 1s
// 	for i := 0; i < n; i++ {
// 		ans = min(ans, l[i]+r[i+1]) // r[n] == 0
// 	}
// 	return ans
// }

// func minFlipsMonoIncr(s string) int {
// 	n := len(s)
// 	dp := make([][]int, n)
// 	for i := range dp {
// 		dp[i] = make([]int, 2)
// 	}

// 	dp[0][0] = int(s[0] - '0') // dp[i][0] := flips so that s[i] == 0 && s[0:i+1] legit
// 	dp[0][1] = int('1' - s[0]) // dp[i][1] := flips so that s[i] == 1 && s[0:i+1] legit

// 	for i := 1; i < n; i++ {
// 		dp[i][0] = dp[i-1][0] + int(s[i]-'0')
// 		dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + int('1'-s[i])
// 	}
// 	return min(dp[n-1][0], dp[n-1][1])
// }

func minFlipsMonoIncr(s string) int { // evolved from the previous solution
	dp0, dp1 := 0, 0
	for i := range s {
		dp0, dp1 = dp0+int(s[i]-'0'), min(dp0, dp1)+int('1'-s[i])
	}
	return min(dp0, dp1)
}
