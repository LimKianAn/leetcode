package main

func updateMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	isVisited := make([][]bool, m)
	for i := range isVisited {
		isVisited[i] = make([]bool, n)
	}

	queue := [][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				isVisited[i][j] = true
				queue = append(queue, [2]int{i, j}) // we start from the zeros
			}
		}
	}

	steps := 0
	for len(queue) > 0 {
		qLen := len(queue)
		for qLen > 0 {
			node := queue[0]
			queue = queue[1:]
			qLen--

			i, j := node[0], node[1]
			ans[i][j] = steps
			for _, dir := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				x := i + dir[0]
				y := j + dir[1]
				if x < 0 || x >= m || y < 0 || y >= n || isVisited[x][y] { // outside or is visited
					continue
				}
				isVisited[x][y] = true
				queue = append(queue, [2]int{x, y})
			}

		}
		steps++ // next level
	}

	return ans
}
