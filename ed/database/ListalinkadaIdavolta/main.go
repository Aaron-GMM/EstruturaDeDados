package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (node *Node) InserirDepois(value int) {
	i :=NewNode(value)
	i.prev = node
	i.next = node.next
	node.next.prev = i
	node.next = i
}

func (node *Node) Sair() {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		next:  nil,
		prev:  nil, // previous
	}
}


type List struct {
	head *Node
	tail *Node
}

func NewList() *List {
	 	head := NewNode(0)
		tail := NewNode(0)
		head.next = tail
		tail.prev = head
	return &List{
		head: head,
		tail: tail,
	}
}
func main() {
	a := NewNode(4)
	b := NewNode(7)
	c := NewNode(3)

	a.next = b
	b.prev = a
	b.next = c
	c.prev = b

	// for node := a; node != nil; node = node.next {
	// 	fmt.Println(node.value)
	// }
	a.InserirDepois(1)

	b.InserirDepois(5)
	for node := a; node != nil; node = node.next {
		fmt.Println(node.value)
             	}
}
