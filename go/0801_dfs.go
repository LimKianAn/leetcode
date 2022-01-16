// 2022.01.16

package main

func minSwap_df(a []int, b []int) int {
	ans := 1<<63 - 1
	var dfs func(i, count int)
	dfs = func(i, count int) {
		if count >= ans {
			return
		}

		if i == len(a) {
			ans = count
			return
		}

		if i == 0 || a[i-1] < a[i] && b[i-1] < b[i] { // always branches when i==0
			dfs(i+1, count)
		}

		if i == 0 || a[i-1] < b[i] && b[i-1] < a[i] { // always branches when i==0
			a[i], b[i] = b[i], a[i]
			dfs(i+1, count+1) // only swaps a and b in this branch
			a[i], b[i] = b[i], a[i]
		}
	}
	dfs(0, 0)
	return ans
}
