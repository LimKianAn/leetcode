// 2022.01.26

package main

func findRedundantConnection_d(edges [][]int) []int {
	connectedNodes := map[int][]int{}
	var isConnected func(u, v int, isVisited map[int]bool) bool
	isConnected = func(u, v int, isVisited map[int]bool) bool {
		if u == v {
			return true
		}
		if connectedNodes[u] == nil || connectedNodes[v] == nil {
			return false
		}
		isVisited[u] = true
		for _, neighbor := range connectedNodes[u] {
			if isVisited[neighbor] {
				continue
			}
			if isConnected(neighbor, v, isVisited) {
				return true
			}
		}
		return false
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if isConnected(u, v, map[int]bool{}) {
			return edge
		}
		connectedNodes[u] = append(connectedNodes[u], v)
		connectedNodes[v] = append(connectedNodes[v], u)
	}
	return []int{}
}
