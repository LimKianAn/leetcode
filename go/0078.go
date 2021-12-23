package main

import "log"

func subsets(nums []int) [][]int {
	a := [][]int{} // a := answer
	sub := []int{}

	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) { // exceeding the upper bound
			log.Print(i, " ", sub)
			a = append(a, copiedInts(sub))
			return
		}

		log.Print(i, "+", sub)
		sub = append(sub, nums[i]) // with the current integer
		dfs(i + 1)

		log.Print(i, "-", sub)
		sub = sub[:len(sub)-1] // without the current integer
		dfs(i + 1)
	}

	dfs(0)
	return a
}

// func main() {
// 	subsets([]int{1, 2, 3})
// }
