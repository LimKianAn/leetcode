package main

func numIslands(grid [][]byte) int {
	return newGrid(grid).numOfIslands()
}

type grid struct {
	grid    [][]byte
	visited [][]bool
}

func newGrid(source [][]byte) *grid {
	return &grid{
		grid:    source,
		visited: boolSlice(source),
	}
}

func (g *grid) isInGrid(i, j int) bool {
	return 0 <= i && i < len(g.grid) && 0 <= j && j < len(g.grid[0])
}

func (g *grid) mark(i, j int) {
	g.visited[i][j] = true
	for _, v := range directions {
		newI := i + v[0]
		newJ := j + v[1]
		if g.isInGrid(newI, newJ) && !g.visited[newI][newJ] && g.grid[newI][newJ] == '1' {
			g.mark(newI, newJ)
		}
	}
}

func (g *grid) numOfIslands() (n int) {
	for i := range g.grid {
		for j := range g.grid[i] {
			if !g.visited[i][j] && g.grid[i][j] == '1' {
				n++
				g.mark(i, j)
			}
		}
	}
	return
}

func boolSlice(source [][]byte) [][]bool {
	a := make([][]bool, len(source))
	for i := range a {
		a[i] = make([]bool, len(source[0]))
	}
	return a
}
