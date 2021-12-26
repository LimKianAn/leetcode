package main

func shortestPathLength(graph [][]int) int {
	type state struct {
		currNode, visitedNodes int
	}

	q := []state{}
	for i := range graph {
		q = append(q, state{
			currNode:     i,
			visitedNodes: 1 << i, // uses int as a bit array to represent if each node is visited
		})
	}

	steps := 0
	hasHappened := map[state]bool{}
	for len(q) > 0 {
		qLen := len(q)
		for qLen > 0 {
			st := q[0]
			q = q[1:]
			qLen--

			curr := st.currNode
			visited := st.visitedNodes
			if visited == 1<<len(graph)-1 { // all nodes are visited (-1!)
				return steps
			}

			if hasHappened[st] { // this route has been taken
				continue
			}
			hasHappened[st] = true

			for _, next := range graph[curr] {
				q = append(q, state{
					currNode:     next,
					visitedNodes: visited | 1<<next, // sets the bit
				})
			}
		}
		steps++
	}
	return -1
}
