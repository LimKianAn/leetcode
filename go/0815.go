// https://zxi.mytechroad.com/blog/searching/leetcode-815-bus-routes/
package main

func numBusesToDestination(routes [][]int, source int, target int) (minBusNum int) {
	if source == target {
		return
	}

	busIDs := map[int][]int{}
	for busID := 0; busID < len(routes); busID++ {
		for _, busStop := range routes[busID] {
			busIDs[busStop] = append(busIDs[busStop], busID)
		}
	}

	busStopQueue := []int{source}
	isBusTaken := make([]bool, len(routes))
	for len(busStopQueue) > 0 {
		minBusNum++

		busStopNum := len(busStopQueue)
		for i := 0; i < busStopNum; i++ {
			busStop := busStopQueue[0]      // front
			busStopQueue = busStopQueue[1:] // pop
			for _, busID := range busIDs[busStop] {
				if isBusTaken[busID] {
					continue
				}
				isBusTaken[busID] = true

				busStops := routes[busID]
				for _, busStop := range busStops {
					if busStop == target {
						return
					}

					busStopQueue = append(busStopQueue, busStop)
				}
			}
		}

	}

	return -1
}
