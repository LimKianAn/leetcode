package main

import "sort"

// func minMalwareSpread(graph [][]int, initial []int) int {
// 	sort.Slice(initial, func(i, j int) bool { // so that the smaller node index stays upon competition
// 		return initial[i] < initial[j]
// 	})

// 	minNum := 1<<63 - 1
// 	ans := -1
// 	for _, initNode := range initial {
// 		infected := map[int]bool{}
// 		q := []int{}
// 		for _, initNd := range initial {
// 			if initNd == initNode {
// 				continue
// 			}
// 			infected[initNd] = true
// 			q = append(q, initNd)
// 		}

// 		for len(q) > 0 {
// 			for qLen := len(q); qLen > 0; qLen-- {
// 				front := q[0]
// 				q = q[1:]

// 				for i := range graph[front] {
// 					if graph[front][i] == 0 || infected[i] {
// 						continue
// 					}
// 					infected[i] = true
// 					q = append(q, i)
// 				}
// 			}
// 		}

// 		if len(infected) < minNum {
// 			minNum = len(infected)
// 			ans = initNode
// 		}
// 	}
// 	return ans
// }

func minMalwareSpread(graph [][]int, initial []int) int {
	sort.Slice(initial, func(i, j int) bool { // so that the smaller node index stays upon competition
		return initial[i] < initial[j]
	})

	minNum := 1<<63 - 1
	ans := -1
	for _, node := range initial {
		infected := map[int]bool{}
		q := []int{}
		for _, nd := range initial {
			if nd == node {
				continue
			}
			infected[nd] = true
			q = append(q, nd)
		}

		for i := 0; i < len(q); i++ {
			for j := range graph[q[i]] {
				if graph[q[i]][j] == 0 || infected[j] {
					continue
				}
				infected[j] = true
				q = append(q, j)
			}
		}

		if len(infected) < minNum {
			minNum = len(infected)
			ans = node
		}
	}
	return ans
}
