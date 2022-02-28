// 2022.02.24

package main

import (
	"math"
	"sort"
)

func minimumTimeRequired(jobs []int, k int) int {
	sort.Slice(jobs, func(i, j int) bool { // descending
		return jobs[i] > jobs[j]
	})

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	workingTime := make([]int, k)

	ans := math.MaxInt
	var dfs func(i, curMax int)
	dfs = func(i, curMax int) {
		if curMax >= ans {
			return
		}

		if i == len(jobs) {
			ans = curMax
			return
		}

		seen := map[int]bool{}
		for j := 0; j < k; j++ { // adds jobs[i] to worker 0 or 1 or 2 or ...
			if seen[workingTime[j]] { // pruning
				continue
			}
			seen[workingTime[j]] = true

			workingTime[j] += jobs[i]
			dfs(i+1, max(curMax, workingTime[j]))
			workingTime[j] -= jobs[i]
		}
	}
	dfs(0, 0)
	return ans
}
