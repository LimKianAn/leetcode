package main

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}

	var getRoot func(id int) int
	getRoot = func(id int) int {
		if parentID := parents[id]; parentID != id {
			parents[id] = getRoot(parentID)
		}
		return parents[id]
	}

	ranks := make([]int, n)
	var unify func(i, j int)
	unify = func(i, j int) {
		rootI := getRoot(i)
		rootJ := getRoot(j)
		if rootI == rootJ {
			return
		}

		if ranks[rootI] > ranks[rootJ] {
			parents[rootJ] = rootI
		} else if ranks[rootI] < ranks[rootJ] {
			parents[rootI] = rootJ
		} else {
			parents[rootJ] = rootI
			ranks[rootI]++ // increments whenever a root grows (possibly more than one child each time)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				unify(i, j)
			}
		}
	}

	roots := map[int]struct{}{}
	for i := 0; i < n; i++ {
		roots[getRoot(i)] = struct{}{}
	}
	return len(roots)
}
