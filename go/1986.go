// 2022.02.25

package main

import "math"

func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)

	initSlice := make([]int, sessionTime+1) // i in the range of 0,1,2,...,sessionTime
	for i := range initSlice {
		initSlice[i] = -1
	}
	dp := make([][]int, 1<<n) // number of possible masks
	for i := range dp {
		copied := make([]int, len(initSlice))
		copy(copied, initSlice)
		dp[i] = copied
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var dfs func(mask, elapsedSessionTime int) int
	dfs = func(mask, elapsedSessionTime int) int {
		if mask == 1<<n-1 { // all one
			return 1 // concludes a session
		}

		if dp[mask][elapsedSessionTime] != -1 {
			return dp[mask][elapsedSessionTime] // pruning
		}

		ans := math.MaxInt
		for i := range tasks {
			if mask&(1<<i) != 0 { // already set
				continue
			}
			newMask := mask | 1<<i
			if newEST := elapsedSessionTime + tasks[i]; newEST <= sessionTime {
				ans = min(ans, dfs(newMask, newEST))
			} else {
				ans = min(ans, 1+dfs(newMask, tasks[i])) // concludes a session and starts a new one
			}
		}
		dp[mask][elapsedSessionTime] = ans
		return ans
	}
	return dfs(0, 0)
}
