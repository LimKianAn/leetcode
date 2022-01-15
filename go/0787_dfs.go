// 2022.01.15

package main

func findCheapestPrice_d(n int, flights [][]int, src int, dst int, k int) int {
	if src == dst {
		return 0
	}

	priceMap := map[int]map[int]int{}
	for _, f := range flights {
		cur, next, price := f[0], f[1], f[2]
		if priceMap[cur] == nil {
			priceMap[cur] = map[int]int{}
		}
		priceMap[cur][next] = price
	}

	visited := make([]bool, n)
	visited[src] = true

	maxInt := 1<<63 - 1
	ans := maxInt

	var dfs func(cur, stop, cost int)
	dfs = func(cur, stop, cost int) {
		if cur == dst {
			ans = cost
			return
		}

		if stop > k {
			return
		}

		for next, price := range priceMap[cur] {
			if visited[next] {
				continue
			}

			nextCost := cost + price
			if nextCost >= ans {
				continue
			}

			visited[next] = true
			dfs(next, stop+1, nextCost)
			visited[next] = false
		}
	}
	dfs(src, 0, 0)

	if ans == maxInt {
		return -1
	}
	return ans
}
