package main

func generate(numRows int) [][]int {
	a := make([][]int, numRows)
	for i := range a {
		a[i] = make([]int, i+1)
		a[i][0], a[i][i] = 1, 1
		for j := 1; j < i; j++ {
			a[i][j] = a[i-1][j-1] + a[i-1][j]
		}
	}
	return a
}
