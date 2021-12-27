package main

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parents := make([]int, n+1) // id in the range [1,n]
	for i := range parents {
		parents[i] = i // loops back, i.e. root
	}

	var getRoot func(id int) int
	getRoot = func(id int) int {
		if parentID := parents[id]; parentID != id {
			parents[id] = getRoot(parentID) // through 2 or 3 layers to the root
		}
		return parents[id]
	}

	ranks := make([]int, n+1)
	var needsMerge func(i, j int) bool
	needsMerge = func(i, j int) bool {
		rootI := getRoot(i)
		rootJ := getRoot(j)
		if rootI == rootJ {
			return false
		}

		if ranks[rootI] > ranks[rootJ] {
			parents[rootJ] = rootI
		} else if ranks[rootI] < ranks[rootJ] {
			parents[rootI] = rootJ
		} else {
			parents[rootJ] = rootI
			ranks[rootI]++ // increments the rank of the new root
		}

		return true
	}

	for _, edge := range edges {
		if !needsMerge(edge[0], edge[1]) {
			return edge
		}
	}
	return []int{}
}
