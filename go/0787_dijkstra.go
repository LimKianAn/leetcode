// 2022.01.15

package main

import "container/heap"

func findCheapestPrice_di(n int, flights [][]int, src int, dst int, k int) int {
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

	q := &costMinPQ{{src, 0, 0}}
	heap.Init(q)
	for q.Len() > 0 {
		front := heap.Pop(q).([3]int)
		city, stop, cost := front[0], front[1], front[2]
		if city == dst {
			return cost
		}
		if stop > k {
			continue
		}
		for to, price := range priceMap[city] {
			heap.Push(q, [3]int{
				to,
				stop + 1,
				cost + price,
			})
		}
	}
	return -1
}

type costMinPQ [][3]int // city, stop, cost

func (q costMinPQ) Len() int {
	return len(q)
}

func (q costMinPQ) Less(i, j int) bool {
	return q[i][2] < q[j][2]
}

func (q costMinPQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *costMinPQ) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func (q *costMinPQ) Push(x interface{}) {
	*q = append(*q, x.([3]int))
}
