package main

func findCheapestPrice_b_s(n int, flights [][]int, src int, dst int, k int) int {
	if src == dst {
		return 0
	}

	flightPrice := map[int]map[int]int{}
	for _, f := range flights {
		from, to, price := f[0], f[1], f[2]
		if flightPrice[from] == nil {
			flightPrice[from] = map[int]int{}
		}
		flightPrice[from][to] = price
	}

	q := [][3]int{{src, 0, 0}} // city, cost, stop

	const maxInt = 1<<63 - 1 // lc doesn't like math.MaxInt
	cheapest := make([]int, n)
	for i := range cheapest {
		cheapest[i] = maxInt
	}
	cheapest[src] = 0

	for i := 0; i < len(q); i++ {
		city, cost, stop := q[i][0], q[i][1], q[i][2]
		for to, price := range flightPrice[city] {
			newCost := cost + price
			if newCost >= cheapest[to] {
				continue
			}
			cheapest[to] = newCost

			if stop == k {
				continue
			}

			if to == dst {
				continue
			}

			q = append(q, [3]int{to, newCost, stop + 1})
		}
	}

	if cheapest[dst] == maxInt {
		return -1
	}
	return cheapest[dst]
}
