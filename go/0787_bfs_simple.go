// 2022.01.15

package main

func findCheapestPrice_b_s(n int, flights [][]int, src int, dst int, k int) int {
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

	q := [][3]int{{src, 0, 0}} // city, stop, cost
	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		city, stop, cost := front[0], front[1], front[2]
		for to, price := range priceMap[city] {
			newCost := cost + price
			if cityToCost[to] <= newCost {
				continue
			}
			cityToCost[to] = newCost

			if stop == k {
				continue
			}

			if to == dst {
				continue
			}

			q = append(q, [3]int{to, stop + 1, newCost})
		}
	}

	if cityToCost[dst] == maxInt {
		return -1
	}
	return cityToCost[dst]
}
