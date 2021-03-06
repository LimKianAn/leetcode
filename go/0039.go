package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	a := [][]int{} // a := answer

	var f func([]int, int, int)
	f = func(combi []int, subT int, index int) {
		if subT < 0 {
			return
		}

		if subT == 0 {
			a = append(a, copiedInts(combi))
			return
		}

		for i := index; i < len(candidates); i++ {
			f(append(combi, candidates[i]), subT-candidates[i], i) // still i, since it can be reused
		}
	}

	f([]int{}, target, 0)

	return a
}

func copiedInts(a []int) []int {
	tmp := make([]int, len(a)) // size must be given
	copy(tmp, a)
	return tmp
}
