// 2021.01.15

package main

func findCheapestPrice_b(n int, flights [][]int, src int, dst int, k int) int {
	if src == dst {
		return 0
	}

	priceMap := map[int]map[int]int{}
	for _, f := range flights {
		from, to, price := f[0], f[1], f[2]
		if priceMap[from] == nil {
			priceMap[from] = map[int]int{}
		}
		priceMap[from][to] = price
	}

	const maxInt = 1<<63 - 1 // lc doesn't like math.MaxInt
	cityToCost := make([]int, n)
	for i := range cityToCost {
		cityToCost[i] = maxInt
	}
	cityToCost[src] = 0

	stop := 0
	q := [][2]int{{src, 0}} // city, cost
	for len(q) > 0 && stop <= k {
		for qLen := len(q); qLen > 0; qLen-- {
			front := q[0]
			q = q[1:]

			cur, curCost := front[0], front[1]
			for next, price := range priceMap[cur] {
				nextCost := curCost + price
				if nextCost >= cityToCost[next] {
					continue
				}
				cityToCost[next] = nextCost

				if cur == dst {
					continue
				}
				q = append(q, [2]int{next, nextCost})
			}
		}
		stop++
	}

	if cityToCost[dst] == maxInt {
		return -1
	}
	return cityToCost[dst]
}
