package main

var dir = [][]int{
	{ 1,  0},
	{-1,  0},
	{ 0,  1},
	{ 0, -1},
}

func solve(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if isOnBorder(board, i, j) && board[i][j] == 'O' {
				replaceO(board, i, j) // replacing the whole area on the border
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '0' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func isOnBorder(board [][]byte, i, j int) bool {
	if i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1 {
		return true
	}
	return false
}

func isOnBoard(board [][]byte, i, j int) bool {
	if 0 <= i && i < len(board) && 0 <= j && j < len(board[0]) {
		return true
	}
	return false
}

func replaceO(board [][]byte, i, j int) {
	if !isOnBoard(board, i, j) {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = '0' // replaced with some other symbol
		for k := range dir {
			replaceO(board, i+dir[k][0], j+dir[k][1])
		}
	}
}
