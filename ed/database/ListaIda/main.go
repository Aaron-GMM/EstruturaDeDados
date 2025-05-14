package main

import "fmt"

type Node struct {
	value int
	next  *Node
	
}

type List struct {
	head *Node
}


func main() {
	fmt.Println("Hello, World!")
}
