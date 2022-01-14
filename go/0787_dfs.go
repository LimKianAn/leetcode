package main

func findCheapestPrice_d(n int, flights [][]int, src int, dst int, k int) int {
	fPrice := map[int]map[int]int{}
	for _, f := range flights {
		from, to, price := f[0], f[1], f[2]
		if fPrice[from] == nil {
			fPrice[from] = map[int]int{}
		}
		fPrice[from][to] = price
	}

	visited := make([]bool, n)
	visited[src] = true

	maxInt := 1<<63 - 1
	ans := maxInt

	var dfs func(cur, cost, stop int)
	dfs = func(cur, cost, stop int) {
		if cur == dst {
			ans = cost
			return
		}

		if stop > k {
			return
		}

		for to, price := range fPrice[cur] {
			if visited[to] {
				continue
			}

			newCost := cost + price
			if newCost >= ans {
				continue
			}

			visited[to] = true
			dfs(to, newCost, stop+1)
			visited[to] = false
		}
	}
	dfs(src, 0, 0)

	if ans == maxInt {
		return -1
	}
	return ans
}
