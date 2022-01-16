// 2022.01.16

package main

// Compare with 0046
func combine(n int, k int) [][]int { // from huahua
	tmp := make([]int, 0, k)

	ans := [][]int{}
	var dfs func(start int)
	dfs = func(start int) {
		if len(tmp) == k {
			cp := make([]int, k)
			copy(cp, tmp)
			ans = append(ans, cp)
			return
		}

		for i := start; i <= n; i++ {
			tmp = append(tmp, i)
			dfs(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(1)
	return ans
}
