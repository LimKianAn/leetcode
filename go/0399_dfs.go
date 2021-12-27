package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := map[string]map[string]float64{}
	for i := range equations {
		a := equations[i][0]
		b := equations[i][1]
		if graph[a] == nil {
			graph[a] = map[string]float64{}
		}
		if graph[b] == nil {
			graph[b] = map[string]float64{}
		}

		quotient := values[i]
		graph[a][b] = quotient
		graph[b][a] = 1 / quotient
	}

	var dfs func(x, z string, isVisited map[string]bool) float64
	dfs = func(x, z string, isVisited map[string]bool) float64 {
		if x == z {
			return 1
		}
		isVisited[x] = true // sets the new entry point to true

		for y, quotientXY := range graph[x] {
			if isVisited[y] {
				continue // avoids loops
			}

			quotientYZ := dfs(y, z, isVisited)
			if quotientYZ == -1 {
				continue // not viable
			}
			return quotientXY * quotientYZ
		}
		return -1
	}

	ans := []float64{}
	for _, query := range queries {
		a := query[0]
		b := query[1]
		if graph[a] == nil || graph[b] == nil {
			ans = append(ans, -1)
			continue // done with the current query
		}

		ans = append(ans, dfs(a, b, map[string]bool{}))
	}
	return ans
}
