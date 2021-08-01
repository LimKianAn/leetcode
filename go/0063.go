package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	a := new2DArray(m, n) // a := array

	// Init the first column
	for i := 1; i < m; i++ {
		if a[i-1][0] == 1 && obstacleGrid[i][0] == 0 {
			a[i][0] = 1
		}
	}

	// Init the first row
	for j := 1; j < n; j++ {
		if a[0][j-1] == 1 && obstacleGrid[0][j] == 0 {
			a[0][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				a[i][j] = a[i][j-1] + a[i-1][j]
			}
		}
	}

	return a[m-1][n-1]
}

func new2DArray(m, n int) [][]int {
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
	}
	a[0][0] = 1
	return a
}
