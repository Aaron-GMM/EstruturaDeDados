package main

import (
	"fmt"	
)
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    nl := len(image)
    nc := len(image[0])
    if  nl == 0 || nc == 0{
        return image
    }
    
    if image[sr][sc] == color {
        return image
    } 
    marcar(image, sr, sc, image[sr][sc], color, nl, nc)
    return image
}

func marcar(image[][]int, l, c int, originalColor, newColor, nl, nc int){
    if(l < 0 || l>= nl || c<0 || c>=nc || image[l][c]!= originalColor){
        return
    }
    image[l][c] = newColor
    marcar(image, l-1, c , originalColor, newColor, nl, nc)
    marcar(image, l+1, c , originalColor, newColor, nl, nc)
    marcar(image, l, c-1 , originalColor, newColor, nl, nc)
    marcar(image, l, c+1 , originalColor, newColor, nl, nc)
}

func main() {
	fmt.Print("Hello, World!\n")
}