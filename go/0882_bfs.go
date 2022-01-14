package main

func reachableNodes_b(edges [][]int, maxMoves int, n int) int {
	graph := map[int]map[int]int{}
	for _, e := range edges {
		u, v, nodeNum := e[0], e[1], e[2]
		if graph[u] == nil {
			graph[u] = map[int]int{}
		}
		if graph[v] == nil {
			graph[v] = map[int]int{}
		}
		graph[u][v] = nodeNum
		graph[v][u] = nodeNum
	}

	nodeToQuota := map[int]int{0: maxMoves} // nodeID to quotaOfMoves

	q := [][2]int{{0, maxMoves}} // [2]int{nodeID, quotaOfMoves};
	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		cur, curQuota := front[0], front[1]
		for next, nodeNum := range graph[cur] {
			nextQuota := curQuota - nodeNum - 1
			if nextQuota < 0 {
				continue
			}
			if quota, ok := nodeToQuota[next]; ok && quota >= nextQuota { // pruning equality!
				continue
			}
			nodeToQuota[next] = nextQuota
			q = append(q, [2]int{next, nextQuota})
		}
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	ans := len(nodeToQuota)
	for _, e := range edges {
		u, v, nodeNum := e[0], e[1], e[2]
		uv := nodeToQuota[u]
		vu := nodeToQuota[v]
		ans += min(uv+vu, nodeNum)
	}
	return ans
}
