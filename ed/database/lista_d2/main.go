package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct{
	value int
	next *Node
	prev *Node
	root *Node
}

type LList struct {
	root *Node
	size int
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		next:  nil,
		prev:  nil,
		root: nil,
	}
}

func NewLList() *LList {
	root := NewNode(0)
	root.next = root
	root.prev = root
	root.root = root

	return &LList{
		root: root,
		size: 0,
	}
}
func (node *Node) Next() *Node {
	if node == node.root{
		return nil
	}
	return node.next
}
func (node *Node) Prev() *Node {
	if node == node.root{
		return nil
	}
	return node.prev
}

func (ll *LList) Size() int {
	return ll.size
}
func (ll *LList) Clear() {
	ll.root = nil
	ll.size = 0
}
func (ll *LList) Front() *Node {
	if ll.root == nil {
		return nil
	}
	return ll.root.next
}
func (ll *LList) Back() *Node {
	if ll.root == nil {
		return nil
	}
	return ll.root.prev
}
func (node *Node) inserirAntes(value int) {
	if node == nil {
		return
	}
	newNode := NewNode(value)
	newNode.next = node
	newNode.prev = node.prev
	node.prev.next = newNode
	node.prev = newNode
}
func (node *Node) inserirDepois(value int) {
	if node == nil {
		return
	}
	newNode := NewNode(value)
	newNode.prev = node
	newNode.next = node.next
	node.next.prev = newNode
	node.next = newNode
}

func (ll *LList) PushBack(value int){
	ll.root.inserirAntes(value)
	ll.size++
}
func (ll *LList) PushFront(value int) {
	ll.root.inserirDepois(value)
	ll.size++
}

func (ll *LList) String() string{
	var sep string = ", "
	var result string
	if ll.root.next == ll.root{
		return "[]"
	}
	result+= "["
	result += fmt.Sprintf("%d", ll.root.next.value)
	for node := ll.root.next.next; node != ll.root; node = node.next{
		if node.next == nil {
			result += fmt.Sprintf("%d", node.value)
		} else {
			result += sep + fmt.Sprintf("%d", node.value)
		}
	}
	result += "]"
	return result
}

func (ll *LList) PopBack() {
	if ll.root.prev == ll.root {
		return
	}
	node := ll.root.prev
	node.prev.next = ll.root
	ll.root.prev = node.prev
	node.prev = nil
	node.next = nil
	ll.size--
}
func (ll *LList) PopFront() {
	if ll.root.next == ll.root {
		return
	}
	node := ll.root.next
	node.next.prev = ll.root
	ll.root.next = node.next
	node.prev = nil
	node.next = nil
	ll.size--
}

func (ll *LList) Search(value int) *Node {
	for node := ll.root.next; node != ll.root; node = node.next {
		if node.value == value {
			return node
		}
	}
	return nil
}
func (ll *LList) Insert(node *Node, value int) error{
	if node == nil || node == ll.root {
		return fmt.Errorf("fail: not found")
	}
	newNode := NewNode(value)
	newNode.prev = node.prev
	newNode.next = node
	node.prev.next = newNode
	node.prev = newNode
	ll.size++
	return nil
}
func (ll *LList) Remove(node *Node) {
	if node == nil || node == ll.root {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
	ll.size--
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				if node == ll.root {
					continue
					} else {
				fmt.Printf("%v ", node.value)
					
				}
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				if node == ll.root {
					continue
				} else {
				fmt.Printf("%v ", node.value)
				}
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
