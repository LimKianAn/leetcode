package main

import "log"

func subsets(nums []int) [][]int {
	a := [][]int{} // a := answer
	sub := []int{}

	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			log.Print(i, " ", sub)
			a = append(a, copied(sub))
			return
		}

		log.Print(i, "+", sub)
		sub = append(sub, nums[i])
		dfs(i + 1)

		log.Print(i, "-", sub)
		sub = sub[:len(sub)-1]
		dfs(i + 1)
	}

	dfs(0)
	return a
}

func copied(ii []int) []int {
	tmp := make([]int, len(ii))
	copy(tmp, ii)
	return tmp
}

func main() {
	subsets([]int{1, 2, 3})
}
