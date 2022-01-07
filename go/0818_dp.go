package main

import . "math"

func racecar_d(target int) int { // from huahua
	dp := make([]int, target+1)
	var f func(target int) int // top-down
	f = func(target int) int {
		if dp[target] > 0 {
			return dp[target]
		}

		// m*A
		// least m s.t.
		// 2^0+2^1+...+2^(m-2)+2^(m-1) >= target
		// 2^m-1 >= target
		m := int(Ceil(Log2(float64(target) + 1)))
		if 1<<m-1 == target {
			dp[target] = m
			return m
		}

		// m*A
		// R
		// ?
		dp[target] = m + 1 + f(1<<m-1-target)

		// (m-1)*A
		// R
		// n*A, n=[0,m-1)
		// R
		// ?
		for n := 0; n < m-1; n++ {
			newStart := 1<<(m-1) - 1 - (1<<n - 1)
			dp[target] = min(dp[target], m-1+1+n+1+f(target-newStart))
		}
		return dp[target]
	}
	return f(target)
}