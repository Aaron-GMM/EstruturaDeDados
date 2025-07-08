package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}
type Btree struct {
	Root *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}
func NewBtree() *Btree {
	return &Btree{
		Root: nil,
	}
}
func NewBtreeSerial(serial string) *Btree {
	bt := &Btree{}
	parts := strings.Fields(serial)
	//fmt.Printf("DEBUG: : %v\n", parts)
	index := 0
	var buildTree func() *Node
	buildTree = func() *Node {
		if index >= len(parts) {
			return nil
		}
		part := parts[index]
		index++
		if part == "#" {
			return nil
		}
		value, err := strconv.Atoi(part)

		if err != nil {
			fmt.Printf("Error converting %s to int: %v\n", part, err)
			return nil
		}
		node := NewNode(value)
		node.Left = buildTree()
		node.Right = buildTree()

		return node
	}
	bt.Root = buildTree()
	return bt
}
// func (bt *Btree) recursiveBShow(node *Node, heranca string) {
// 	if node != nil && (node.Left != nil || node.Right != nil) {
// 		bt.recursiveBShow(node.Left, heranca+"l")
// 	}

// 	for i := 0; i < len(heranca)-1; i++ {
// 		if heranca[i] != heranca[i+1] {
// 			fmt.Print("│   ")
// 		} else {
// 			fmt.Print("    ")
// 		}
// 	}

// 	if heranca != "" {
// 		if heranca[len(heranca)-1] == 'l' {
// 			fmt.Print("┌───")
// 		} else {
// 			fmt.Print("└───")
// 		}
// 	}

// 	if node == nil {
// 		fmt.Println("#")
// 		return
// 	}

// 	fmt.Println(node.Value)

// 	if node.Left != nil || node.Right != nil {
// 		bt.recursiveBShow(node.Right, heranca+"r")
// 	}
// }
//func (bt *Btree) bShow() {bt.recursiveBShow(bt.Root, "")}
func (bt *Btree) Destroy() { bt.Root = nil }

func (bt *Btree) ShowInOrder() {
	fmt.Print("[ ")
	bt.recursiveShowInOrder(bt.Root)
	fmt.Println("]")
}

func (bt *Btree) recursiveShowInOrder(node *Node) {
	if node == nil {
		return
	}
	bt.recursiveShowInOrder(node.Left)
	fmt.Printf("%d ", node.Value)
	bt.recursiveShowInOrder(node.Right)
}

func main() {
	var line string
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line = scanner.Text()
//	fmt.Printf("DEBUG:  '%s'\n", line)
	bt := NewBtreeSerial(line)
	bt.ShowInOrder()

	//bt.bShow()
}
