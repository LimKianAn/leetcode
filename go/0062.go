package main

func uniquePaths(m int, n int) int {
	a := newArray(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				// only one unique path along the first row and column
				a[i][j] = 1
				continue
			}
			a[i][j] = a[i][j-1] + a[i-1][j]
		}
	}
	return a[m-1][n-1]
}

func newArray(m, n int) [][]int {
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
	}
	return a
}
