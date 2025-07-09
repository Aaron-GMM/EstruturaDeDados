package main
import (
	"fmt"
)
func countBattleships(board [][]byte) int {
	nl := len(board)
	nc := len(board[0])
	if nl == 0 || nc == 0 {
		return 0
	}

	count := 0 

	for i:=0; i < nl; i++ {
		for j:=0; j < nc; j++ {
			if board[i][j] == 'X' {
				if i==0 || board[i-1][j] == '.' {
					if j==0 || board[i][j-1] == '.' {
						count++
					}
				}
			}
		}
	}

	return count
}
func main() {
	fmt.Print("Hello, World!\n")
}