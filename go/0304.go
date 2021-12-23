package main

type NumMatrix struct {
	sum [][]int
}

// O(m*n)
func Constructor0304(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])

	// sum is (m+1)*(n+1) so that the 0th row and column of matrix
	// can be treated uniformly
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i := range matrix {
		sum[i+1] = make([]int, n+1)
		for j := range matrix[i] {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] - this.sum[row2+1][col1] - this.sum[row1][col2+1] + this.sum[row1][col1]
}
