// 2022.01.16

package main

// Compare with 0046
func combine(n int, k int) [][]int {
	ans := [][]int{}
	var dfs func(start int, combi []int)
	dfs = func(start int, combi []int) {
		if len(combi) == k {
			cp := make([]int, k)
			copy(cp, combi)
			ans = append(ans, cp)
			return
		}

		for i := start; i <= n; i++ {
			dfs(i+1, append(combi, i))
		}
	}
	dfs(1, make([]int, 0, k))
	return ans
}
