package main

import "fmt"
func longestIncreasingPath(matrix [][]int) int {
        nl := len(matrix)
        nc := len(matrix[0])
        memo := make([][]int, nl)
        if nl == 0 || nc ==0{
            return 0
        }
        for i := range memo{
            memo[i] = make([]int, nc)
        }
        maiorCaminhoGeral := 0
        for r:=0; r < nl; r++{
            for c:=0; c< nc; c++{
                			maiorCaminhoGeral = max(maiorCaminhoGeral, dfs(matrix, r, c, nl, nc, memo))
            }
        }
	return maiorCaminhoGeral
}
func dfs(matrix [][]int, r,c, rows, cols int, memo [][]int) int{
    if memo[r][r] !=0{
        return memo[r][c]
    }
    currentPathLength := 1
    if r-1>=0 && matrix[r-1][c]>matrix[r][c]{
        		currentPathLength = max(currentPathLength, 1+dfs(matrix, r-1, c, rows, cols, memo))
    }
    if r+1>=0 && matrix[r+1][c]>matrix[r][c]{
        		currentPathLength = max(currentPathLength, 1+dfs(matrix, r-1, c, rows, cols, memo))
    }
	if c-1 >= 0 && matrix[r][c-1] > matrix[r][c] {
        		currentPathLength = max(currentPathLength, 1+dfs(matrix, r, c-1, rows, cols, memo))
    }
    if c+1 < cols && matrix[r][c+1] > matrix[r][c] {
        		currentPathLength = max(currentPathLength, 1+dfs(matrix, r, c+1, rows, cols, memo))
    }
    memo[r][c] = currentPathLength
	return currentPathLength
}
func max(a, b int) int{
    if a>b{
        return a
    }
    return b
}
func main(){
 fmt.Print()
}