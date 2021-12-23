package main

func exist(board [][]byte, word string) bool {
	visited := makeBoolArray(len(board), len(board[0]))
	c := ctx{board: board, word: word}
	for i, v := range board {
		for j := range v {
			if c.existing(visited, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func makeBoolArray(m, n int) [][]bool {
	a := make([][]bool, m)
	for i := range a {
		a[i] = make([]bool, n)
	}
	return a
}

type ctx struct {
	board [][]byte
	word  string
}

func (c *ctx) existing(visited [][]bool, i, j int, index int) bool {
	if index == len(c.word)-1 {
		return c.board[i][j] == c.word[index]
	}

	if c.board[i][j] == c.word[index] {
		visited[i][j] = true
		for _, v := range directions {
			newI := i + v[0]
			newJ := j + v[1]
			if c.isIn(newI, newJ) && !visited[newI][newJ] && c.existing(visited, newI, newJ, index+1) {
				return true
			}
		}
		visited[i][j] = false
	}
	return false
}

func (c *ctx) isIn(i, j int) bool {
	return 0 <= i && i < len(c.board) && 0 <= j && j < len(c.board[0])
}
