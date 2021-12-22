package main

// time: O(m * log n)
func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for l, h := 0, len(matrix[0])-1; l <= h; { // l := low, h := high
			m := l + (h-l)>>1 // m := middle
			if target < row[m] {
				h = m - 1
			} else if row[m] < target {
				l = m + 1
			} else {
				return true
			}
		}
	}
	return false
}
