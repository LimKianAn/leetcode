package main

func minimumTotal(triangle [][]int) int {
	for r := len(triangle) - 2; r >= 0; r-- { // r := row
		for c := 0; c < len(triangle[r]); c++ { // c := column
			triangle[r][c] += min(triangle[r+1][c], triangle[r+1][c+1])
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
