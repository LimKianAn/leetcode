package main

import "container/heap"

func reachableNodes_d(edges [][]int, maxMoves int, n int) int {
	graph := map[int]map[int]int{}
	for _, e := range edges {
		u, v, nodeNum := e[0], e[1], e[2]
		if graph[u] == nil {
			graph[u] = map[int]int{}
		}
		if graph[v] == nil {
			graph[v] = map[int]int{}
		}
		graph[u][v] = nodeNum
		graph[v][u] = nodeNum
	}

	nodeToQuota := map[int]int{} // node ID to quota of moves

	q := &quotaPQ{}
	heap.Init(q)
	heap.Push(q, [2]int{0, maxMoves})
	for q.Len() > 0 {
		front := heap.Pop(q).([2]int)
		cur, quota := front[0], front[1]
		if _, ok := nodeToQuota[cur]; ok {
			continue
		}
		nodeToQuota[cur] = quota

		for next, nodeNum := range graph[cur] {
			nextQuota := quota - nodeNum - 1
			if nextQuota < 0 {
				continue
			}
			if _, ok := nodeToQuota[next]; ok {
				continue
			}
			heap.Push(q, [2]int{next, nextQuota})
		}
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	ans := len(nodeToQuota)
	for _, e := range edges {
		u, v, nodeNum := e[0], e[1], e[2]
		uv := nodeToQuota[u]
		vu := nodeToQuota[v]
		ans += min(uv+vu, nodeNum)
	}
	return ans
}

type quotaPQ [][2]int // PQ := priority queue, [2]int{nodeID, moveQuota}

func (q quotaPQ) Len() int {
	return len(q)
}

func (q quotaPQ) Less(i, j int) bool {
	return q[i][1] > q[j][1] // max heap!
}

func (q quotaPQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *quotaPQ) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func (q *quotaPQ) Push(x interface{}) {
	*q = append(*q, x.([2]int))
}
