package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Node struct {
	value int 
	next *Node
	prev *Node
}
type List struct{
	root *Node
}
func NewNode(value int) *Node{
	return &Node{
		value: value,
		next:  nil,
		prev:  nil, 
	}
}

func (node *Node) exit(){
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}
func (node *Node) inserirDepois(value int ){
	if node == nil {
		return
	}
	

	newNode := NewNode(value)
	newNode.prev = node 
	newNode.next = node.next
	node.next.prev = newNode
	node.next = newNode


}
func (node *Node) inserirAntes(value int){
	if node == nil{
		return
	}
	newNode := NewNode(value)
	newNode.next = node
	newNode.prev = node.prev
	node.prev.next = newNode
	node.prev = newNode
	
}

func NewList() *List {
	root := NewNode(0)
	root.next = root
	root.prev = root
	return &List{
		root: root,
	}
}

func (l *List) PushBack(value int) {
		 l.root.inserirAntes(value)
}

func (l *List) PushFront(value int) {
		l.root.inserirDepois(value)
}
func (l *List) String() string {
	var sep string = ", "
	var result string
	if l.root.next == l.root {	
		return "[]"
	}
	result += "["
	result += fmt.Sprintf("%d", l.root.next.value)
	for node := l.root.next.next; node != l.root; node = node.next {
		if node.next == nil{
			 result += fmt.Sprintf("%d", node.value)
		}else{
			result +=  sep +  fmt.Sprintf("%d", node.value) 
		}

	}
	result += "]"
	return result
}
func (l *List) PopBack() {
	if l.root.prev == l.root {
		return 
	}
	l.root.prev.exit()
}
func (l *List) PopFront() {
	if l.root.next == l.root{
		return
	}
	l.root.next.exit()
}

func (l *List) Size() int {
	count := 0
	for node := l.root.next; node != l.root; node = node.next {
		count++
	}
	return count
}
func (l *List) Clear(){
	l.root.next = l.root
	l.root.prev = l.root
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			 fmt.Println(ll.String())
		case "size":
			 fmt.Println(ll.Size())
		case "push_back":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushBack(num)
			 }
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			 ll.PopBack()
		case "pop_front":
			 ll.PopFront()
		case "clear":
			 ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
