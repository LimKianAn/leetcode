package main

func findCheapestPrice_b(n int, flights [][]int, src int, dst int, k int) int {
	const maxInt = 1<<63 - 1 // lc doesn't like math.MaxInt

	flightPrice := map[int]map[int]int{}
	for _, f := range flights {
		from, to, price := f[0], f[1], f[2]
		if flightPrice[from] == nil {
			flightPrice[from] = map[int]int{}
		}
		flightPrice[from][to] = price
	}

	q := [][2]int{{src, 0}} // city, cost

	cheapest := make([]int, n)
	for i := range cheapest {
		cheapest[i] = maxInt
	}
	cheapest[src] = 0

	stop := 0
	for len(q) > 0 && stop <= k {
		for qLen := len(q); qLen > 0; qLen-- {
			front := q[0]
			q = q[1:]

			city, cost := front[0], front[1]
			for to, price := range flightPrice[city] {
				newCost := cost + price
				if newCost >= cheapest[to] {
					continue
				}
				cheapest[to] = newCost

				if city == dst {
					continue
				}

				q = append(q, [2]int{to, newCost})
			}
		}
		stop++
	}

	if cheapest[dst] == maxInt {
		return -1
	}
	return cheapest[dst]
}
