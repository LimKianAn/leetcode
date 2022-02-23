package main

func findRedundantConnection_u(edges [][]int) []int {
	n := len(edges)
	ancestors := make([]int, n+1) // id in the range [1,n]
	for i := range ancestors {
		ancestors[i] = i // loops back, i.e. root
	}

	var getRoot func(id int) int
	getRoot = func(id int) int {
		if ancestorID := ancestors[id]; ancestorID != id {
			ancestors[id] = getRoot(ancestorID) // through 0 or 1 middle stop to the root, then path compression
		}
		return ancestors[id]
	}

	ranks := make([]int, n+1)
	var needsUnification func(i, j int) bool
	needsUnification = func(i, j int) bool {
		rootI := getRoot(i)
		rootJ := getRoot(j)
		if rootI == rootJ {
			return false
		}

		if ranks[rootI] > ranks[rootJ] {
			ancestors[rootJ] = rootI
		} else if ranks[rootI] < ranks[rootJ] {
			ancestors[rootI] = rootJ
		} else {
			ancestors[rootJ] = rootI
			ranks[rootI]++ // increments the rank of the new root
		}

		return true
	}

	for _, edge := range edges {
		if !needsUnification(edge[0], edge[1]) {
			return edge
		}
	}
	return []int{}
}
