package main

func getRow(rowIndex int) []int {
	a := make([]int, rowIndex+1)
	a[0] = 1
	for i := 1; i <= rowIndex; i++ {
		a[i] = a[i-1] * (rowIndex - (i - 1)) / i
	}
	return a
}
