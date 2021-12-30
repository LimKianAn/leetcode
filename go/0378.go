package main

func kthSmallestBinarySearch(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]+1 // [l,r)
	for l < r {
		m := l + (r-l)>>1
		count := 0 // how many <= m
		j := n - 1
		for i := range matrix {
			for ; j >= 0 && matrix[i][j] > m; j-- {
			}
			count += j + 1 // [i][0] to [i][j] <= m

			if j < n-1 {
				j++
			}
		}

		if k <= count {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
