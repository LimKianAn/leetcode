// https://zxi.mytechroad.com/blog/searching/leetcode-815-bus-routes/
package main

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	stopToRouteIDs := map[int][]int{}
	for routeID, route := range routes {
		for _, busStop := range route {
			stopToRouteIDs[busStop] = append(stopToRouteIDs[busStop], routeID)
		}
	}

	taken := map[int]bool{} // taken[routeID]
	q := []int{source}      // stops
	busNum := 1
	for len(q) > 0 {
		for qLen := len(q); qLen > 0; qLen-- {
			busStop := q[0] // front
			q = q[1:]       // pop

			for _, routeID := range stopToRouteIDs[busStop] {
				if taken[routeID] {
					continue
				}
				taken[routeID] = true

				for _, busStop := range routes[routeID] {
					if busStop == target {
						return busNum
					}
					q = append(q, busStop)
				}
			}
		}
		busNum++
	}
	return -1
}
