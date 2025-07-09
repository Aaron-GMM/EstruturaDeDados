package main
import (
	"fmt"
)

func solve(board [][]byte)  {
    rows := len(board)
	cols := len(board[0])
	if rows == 0 || cols == 0 {
		return
	}
	for c:= 0; c < cols; c++ {
		if board[0][c] == 'O' {
			dfs(board, 0, c, rows, cols) 
		}
		if board[rows-1][c] == 'O' {
			dfs(board, rows-1, c, rows, cols) 
		}
	}
	for r:= 0; r < rows; r++ {
		if board[r][0] == 'O' {
			dfs(board, r, 0, rows, cols) 
		}
		if board[r][cols-1] == 'O' {
			dfs(board, r, cols-1, rows, cols) 
		}
	}

	for r:=0; r< rows; r++ {
		for c:=0; c< cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}else if board[r][c] == '#' {
				board[r][c] = 'O'
			}
		}
	}				 
}

func dfs(board [][]byte, r, c, rows, cols int){
	if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != 'O' {
		return
	}
	board[r][c] = '#' 
	dfs(board, r-1, c, rows, cols) // Up
	dfs(board, r+1, c, rows, cols) // Down
	dfs(board, r, c-1, rows, cols) // Left
	dfs(board, r, c+1, rows, cols) // Right
}
func main() {
	fmt.Print("Hello, World!\n")
}