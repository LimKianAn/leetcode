// 2022.01.17

package main

func calcEquation_d(equations [][]string, values []float64, queries [][]string) []float64 {
	ratioMap := map[string]map[string]float64{}
	for i := range equations {
		a := equations[i][0]
		b := equations[i][1]
		if ratioMap[a] == nil {
			ratioMap[a] = map[string]float64{}
		}
		if ratioMap[b] == nil {
			ratioMap[b] = map[string]float64{}
		}
		ratioMap[a][b] = values[i]
		ratioMap[b][a] = 1 / values[i]
	}

	var dfs func(x, z string, visited map[string]bool) float64
	dfs = func(x, z string, visited map[string]bool) float64 {
		if x == z {
			return 1
		}

		visited[x] = true // sets the new entry point to true

		for y, ratioXY := range ratioMap[x] {
			if visited[y] {
				continue // avoids loops
			}

			ratioYZ := dfs(y, z, visited)
			if ratioYZ == -1 {
				continue // not viable
			}
			return ratioXY * ratioYZ
		}
		return -1
	}

	ans := []float64{}
	for _, query := range queries {
		a := query[0]
		b := query[1]
		if ratioMap[a] == nil || ratioMap[b] == nil {
			ans = append(ans, -1)
			continue // done with the current query
		}

		ans = append(ans, dfs(a, b, map[string]bool{}))
	}
	return ans
}
