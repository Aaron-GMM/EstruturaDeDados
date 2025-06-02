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
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root 
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
}

func  addsorted(lla *LList, value int) {
	if lla.root.next == lla.root { 
		lla.PushBack(value)
		return
	}

	node := lla.root.next
	for node != lla.root && node.Value < value {
		node = node.next
	}

	if node == lla.root { 
		lla.PushBack(value)
	} else {
		lla.insertBefore(node, value)
	}

}

func Equals(listA *LList, listB *LList) bool {

	if listA.root.Value != listB.root.Value {
		return true
	}
	nodeA := listA.root.next
	nodeB := listB.root.next

	for nodeA != listA.root && nodeB != listB.root {
		if nodeA.Value != nodeB.Value {
			return false
		}
		nodeA = nodeA.next
		nodeB = nodeB.next
	}

	return nodeA == listA.root && nodeB == listB.root
}
func reverse(l *LList) {
     node := l.root
    for {
        node.next, node.prev = node.prev, node.next
        node = node.prev 
        if node == l.root {
            break
        }
    }
}
func merge(lla *LList, llb *LList) *LList {
	merged := NewLList()
	nodeA := lla.root.next
	nodeB := llb.root.next

	for nodeA != lla.root && nodeB != llb.root {
		if nodeA.Value < nodeB.Value {
			merged.PushBack(nodeA.Value)
			nodeA = nodeA.next
		} else {
			merged.PushBack(nodeB.Value)
			nodeB = nodeB.next
		}
	}

	for nodeA != lla.root {
		merged.PushBack(nodeA.Value)
		nodeA = nodeA.next
	}

	for nodeB != llb.root {
		merged.PushBack(nodeB.Value)
		nodeB = nodeB.next
	}

	return merged
}
func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}
func String(l *LList ) string {
	var sb strings.Builder
	sb.WriteString("[")
	node := l.root.next
	for node != l.root {
		sb.WriteString(strconv.Itoa(node.Value))
		node = node.next
		if node != l.root {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if Equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(String(lla))
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(String(lla))
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(String(merged))
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
