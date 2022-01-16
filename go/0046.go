// 2022.01.16

package main

// Compare with 0077
func permute(nums []int) [][]int { // from huahua
	n := len(nums)
	used := make([]bool, n)
	tmp := make([]int, 0, n)

	ansLen := 1
	var fatorial func(i int)
	fatorial = func(i int) {
		if i == 1 {
			return
		}
		ansLen *= i
		fatorial(i - 1)
	}
	fatorial(n)

	ans := make([][]int, 0, ansLen)
	var dfs func()
	dfs = func() {
		if len(tmp) == n { // len(tmp)==k, k:=number of selected
			copied := make([]int, n)
			copy(copied, tmp)
			ans = append(ans, copied)
			return
		}

		for i, num := range nums {
			if used[i] {
				continue
			}
			used[i] = true

			tmp = append(tmp, num)
			dfs()
			tmp = tmp[:len(tmp)-1]

			used[i] = false
		}
	}
	dfs()
	return ans
}
