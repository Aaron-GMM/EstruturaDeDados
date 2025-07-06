 package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0 
	}
	var contIslan int
	m := len(grid)
	n := len(grid[0])
	for i := 0; i<m; i++ {
		for j := 0; j<n; j++{
			if grid[i][j] == '1' {
				contIslan++				
				marcarIlha(grid, i, j) 
			}
		}
	}

	
	return contIslan
}
func marcarIlha(grid [][]byte, i, j int){
	m := len(grid)
	n := len(grid[0])
	if (i < 0 || i>= m) || (j <0 || j>=n){
		return
	}
	if grid[i][j] != '1'{
		return
	}
	grid[i][j] = '0' 
	marcarIlha(grid, i-1, j) // up
	marcarIlha(grid, i+1, j) // down
	marcarIlha(grid, i, j-1) // left
	marcarIlha(grid, i, j+1) // reight
}
	
// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
