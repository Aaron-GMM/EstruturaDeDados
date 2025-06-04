package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Value rune
	Next  *Node
	Prev  *Node
}
type Editor struct {
	Head   *Node
	Tail   *Node
	Cursor *Node
}

func NewNode(value rune) *Node {
	return &Node{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}
func NewEditor() *Editor {
	head := NewNode(0)
	tail := NewNode(0)
	head.Next = tail
	tail.Prev = head
	return &Editor{
		Head:   head,
		Tail:   tail,
		Cursor: tail,
	}
}

func (e *Editor) insert(value rune) {
	newNode := NewNode(value)
	prevNode := e.Cursor.Prev

	prevNode.Next = newNode
	newNode.Prev = prevNode
	newNode.Next = e.Cursor
	e.Cursor.Prev = newNode
}

func (e *Editor) moveLeft() {
	if e.Cursor.Prev != e.Head {
		e.Cursor = e.Cursor.Prev
	}
}
func (e *Editor) moveRight() {
	if e.Cursor != e.Tail {
		e.Cursor = e.Cursor.Next
	}
}
func (e *Editor) backspace() {
	toDelete := e.Cursor.Prev
	if toDelete != e.Head {
		left := toDelete.Prev
		right := e.Cursor
		left.Next = right
		right.Prev = left
	}
}

func (e *Editor) deleteAtCursor() {
	if e.Cursor != e.Tail {
		toDelete := e.Cursor
		left := toDelete.Prev
		right := toDelete.Next
		left.Next = right
		right.Prev = left
		e.Cursor = right
	}
}
func (e *Editor) showDisplay() string {
	var result []rune 
	for node := e.Head.Next; node != e.Tail; node = node.Next {
		if node == e.Cursor {
			result = append(result, '|')
		}
		if node.Value == '\n' {
			result = append(result, '\n')
		} else {
			result = append(result, node.Value)
		}
	}
	if e.Cursor == e.Tail {
		result = append(result, '|')
	}
   
	return string(result) 
}
func processInput(input string) string {
	editor := NewEditor()
	for _, char := range input {
		switch char {
		case 'R':
			editor.insert('\n')

		case 'B':
			editor.backspace()

		case 'D':
			editor.deleteAtCursor()

		case '<':
			editor.moveLeft()

		case '>':
			editor.moveRight()  

		default:
			if (char >= 'a' && char <= 'z') || char == '-' {
				editor.insert(char)
			}
		}
	}

	return editor.showDisplay()

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		input := scanner.Text()

		if input == "" || input == "exit" {
			break
		}

		resul := processInput(input)
		fmt.Println(resul)
   

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

}
