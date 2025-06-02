package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}
type List struct {
	Head *Node
	Tail *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}
func NewList() *List {
	head := NewNode(0)
	tail := NewNode(0)
	head.Next = tail
	tail.Prev = head
	return &List{
		Head: head,
		Tail: tail,
	}
}

func ToStr(l *list.List, sword *list.Element) string {
	result := ""
	result += "[ "
	for e := l.Front(); e != nil; e = e.Next() {
		if e == sword {
			result += fmt.Sprintf("%d> ", e.Value)
		} else {
			result += fmt.Sprintf("%d ", e.Value)
		}
	}
	result += "]\n"
	return result[:len(result)-1] 

}


func Next(l *list.List, it *list.Element) *list.Element {
	if it == nil {
		return nil
	}
	if it.Next() == nil {
		return l.Front()
	}
	return it.Next()

}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := list.New()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Remove(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
