package main

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	const red, blue = 0, 1

	edges := make([][]map[int]struct{}, 2) // [color][from][to]
	for i := range edges {
		edges[i] = make([]map[int]struct{}, n)
		for j := range edges[i] {
			edges[i][j] = map[int]struct{}{}
		}
	}

	for _, redEdge := range redEdges {
		from, to := redEdge[0], redEdge[1]
		edges[red][from][to] = struct{}{}
	}
	for _, blueEdge := range blueEdges {
		from, to := blueEdge[0], blueEdge[1]
		edges[blue][from][to] = struct{}{}
	}

	seen := make([]map[int]bool, 2) // [color][to]
	for i := range seen {
		seen[i] = make(map[int]bool, n)
	}
	seen[red][0] = true
	seen[blue][0] = true

	type state struct {
		lastEdgeColor, node int
	}
	q := []state{
		{red, 0},
		{blue, 0},
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1 // distinguishes unreachable from zero
	}
	step := 0
	for len(q) > 0 {
		for qLen := len(q); qLen > 0; qLen-- {
			stt := q[0]
			q = q[1:]

			color := stt.lastEdgeColor
			node := stt.node
			if ans[node] == -1 {
				ans[node] = step
			} else {
				ans[node] = min(ans[node], step)
			}

			for to := range edges[color][node] {
				nextColor := 1 - color
				if seen[nextColor][to] {
					continue
				}
				seen[nextColor][to] = true
				q = append(q, state{nextColor, to})
			}
		}
		step++
	}
	return ans
}
