package main

func shortestPathLength(graph [][]int) int {
	n := len(graph)

	type state struct {
		curNode, visitedNodes int
	}
	happened := map[state]bool{}

	q := make([]state, 0, n)
	for i := range graph { // different starting points
		st := state{
			curNode:      i,
			visitedNodes: 1 << i, // uses int as a bit array to represent if each node is visited
		}
		q = append(q, st)
		happened[st] = true
	}

	length := 0
	for len(q) > 0 {
		for qLen := len(q); qLen > 0; qLen-- {
			st := q[0]
			q = q[1:]

			cur := st.curNode
			visited := st.visitedNodes
			if visited == 1<<n-1 { // all nodes are visited
				return length
			}

			for _, next := range graph[cur] {
				st := state{
					curNode:      next,
					visitedNodes: visited | 1<<next,
				}
				if happened[st] { // this path has been taken
					continue
				}
				q = append(q, st)
				happened[st] = true
			}
		}
		length++
	}
	return -1
}
