package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)
type Sudoku struct {
    tamanho int // tamanho do tabuleiro 4 ou 9
    dim int
    grid [][]rune
    holes [][2]int

}

func newSudoku(lines []string) *Sudoku {
    tamanho := len(lines)
    dim := int(math.Sqrt(float64(tamanho)))
    grid := make([][]rune, tamanho)
    var holes [][2]int

    for i := 0; i < tamanho; i++ {
        grid[i] = []rune(lines[i])
        for j, ch := range grid[i] {
            if ch == '.' {
                holes = append(holes, [2]int{i, j})
            }
        }
    }

    return &Sudoku{tamanho: tamanho, dim: dim, grid: grid, holes: holes}
}
func (s *Sudoku) canPlace(a,b int, val rune) bool {
    for i:=0; i < s.tamanho; i++ {
        if( s.grid[a][i] == val) || (s.grid[i][b] == val) {
            return false
        }
    } 

    ba := (a/ s.dim) * s.dim
    bb := (b/ s.dim) * s.dim

    for i := ba; i < s.dim; i++ {
        for j := bb; j < s.dim; j++ {
            if s.grid[i][j] == val {
                return false
            }
        }
    }

    return true
}
func (s *Sudoku)solve(idx int) bool {
    if idx == len(s.holes){
        return true
    }
    a, b := s.holes[idx][0], s.holes[idx][1]

    for ch := 1; ch<=s.tamanho  ; ch++ {
        val := rune(ch + '0')

        if s.canPlace(a, b, val) {
            s.grid[a][b] = val
            if s.solve(idx + 1) {
                return true
            }
            s.grid[a][b] = '.' // backtrack
        }
    }
    return false
}
func (s *Sudoku) String() string {
    var result string
    for i := 0; i < s.tamanho; i++ {
        for j := 0; j < s.tamanho; j++ {
            result += string(s.grid[i][j])
        }
       
        if  (i < s.tamanho-1) {
             result += "\n"
        }
    }
    return result
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
        
    scanner.Scan()
    var N int
    fmt.Sscan(scanner.Text(), &N)

    lines := make([]string, N)
    for i := 0; i < N; i++ {
        scanner.Scan()
        lines[i] = scanner.Text()
    }
    sudoku := newSudoku(lines)
    if sudoku.solve(0) {
        fmt.Println(sudoku)
        
    }
}
