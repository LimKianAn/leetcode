package main

// Compare with 0046
func combine(n int, k int) [][]int { // from huahua
	ans := [][]int{}
	tmp := make([]int, 0, k)
	var dfs func(curr int)
	dfs = func(curr int) {
		if len(tmp) == k {
			copied := make([]int, len(tmp))
			copy(copied, tmp)
			ans = append(ans, copied)
			return
		}

		for i := curr; i <= n; i++ {
			tmp = append(tmp, i)
			dfs(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(1)
	return ans
}
