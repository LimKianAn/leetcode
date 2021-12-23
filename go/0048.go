package main

func rotate0048(matrix [][]int) {
	n := len(matrix)

	// 1, 2
	// 3, 4

	// Flip along the diagonal
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 1, 3
	// 2, 4

	// Flip along the cenerline
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}

	// 3, 1
	// 4, 2
}
