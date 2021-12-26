package main

func numIslandsDFS(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	num := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num += int(grid[i][j] - '0') // increments upon 1
			dfs(i, j)                    // removes its 1 and surrounding 1s
		}
	}
	return num
}
