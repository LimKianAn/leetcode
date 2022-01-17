// 2022.01.16

package main

// Compare with 0077
func permute(nums []int) [][]int {
	n := len(nums)
	used := make([]bool, n)

	ansLen := 1
	var factorial func(i int)
	factorial = func(i int) {
		if i == 1 {
			return
		}
		ansLen *= i
		factorial(i - 1)
	}
	factorial(n)

	ans := make([][]int, 0, ansLen)
	var dfs func(permu []int)
	dfs = func(permu []int) {
		if len(permu) == n { // len(tmp)==k, k:=number of selected
			cp := make([]int, n)
			copy(cp, permu)
			ans = append(ans, cp)
			return
		}

		for i, num := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			dfs(append(permu, num))
			used[i] = false
		}
	}
	dfs(make([]int, 0, n))
	return ans
}
