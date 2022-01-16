// 2022.01.16

package main

func subsets(nums []int) [][]int {
	n := len(nums)

	sub := []int{}

	ans := make([][]int, 0, 2^n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) { // exceeding the upper bound
			cp := make([]int, len(sub))
			copy(cp, sub)
			ans = append(ans, cp)
			return
		}

		sub = append(sub, nums[i]) // with the current integer
		dfs(i + 1)
		sub = sub[:len(sub)-1] // without the current integer

		dfs(i + 1)
	}
	dfs(0)
	return ans
}
