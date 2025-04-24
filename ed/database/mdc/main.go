package main

import (
	"fmt"
)

func mdc(a, b int) int {
	
	if a == 0{
		return b
	}
	if  b == 0 {
		return a
	}
		r:= a % b 
	return mdc(b, r)
	}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
