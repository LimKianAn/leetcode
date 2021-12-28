package main

func findRedundantConnectionDFS(edges [][]int) []int {
	graph := map[int][]int{}
	var connected func(u, v int, visited map[int]bool) bool
	connected = func(u, v int, visited map[int]bool) bool {
		if u == v {
			return true
		}

		if graph[u] == nil || graph[v] == nil {
			return false
		}

		visited[u] = true
		for _, neighbor := range graph[u] { // don't forget the slient index!
			if visited[neighbor] {
				continue
			}

			if connected(neighbor, v, visited) {
				return true
			}
		}

		return false
	}

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		if connected(u, v, map[int]bool{}) {
			return edge
		}

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	return []int{}
}
