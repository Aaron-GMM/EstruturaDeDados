package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func (q *Queue[T]) Enqueue(value T)
// func (q *Queue[T]) Dequeue() (T, bool)
// func (q *Queue[T]) Peek() (T, bool)
// func (q *Queue[T]) Size() int
// func (q *Queue[T]) IsEmpty() bool
// func (q *Queue[T]) Clear()


type Node[T any] struct {
	Value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}
func NewNode[T any](value T) *Node[T]{
	return &Node[T]{
		Value: value,
		 next: nil}
}

func (q *Queue[T]) String() string {
	result := "["
	for n := q.head; n != nil; n = n.next {
		if n != q.head {
			result += ", "
		}
		result += fmt.Sprintf("%v", n.Value)
	}
	return result + "]"
}
func (q *Queue[T]) Enqueue(value T){
	node := NewNode(value)
	if q.tail == nil{
		q.head = node
		q.tail = node
	
	}else{
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue[T]) Dequeue() (T, bool){
		if q.head == nil{
			var zero T
			return zero, false
		}
	value := q.head.Value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return value, true
}

func (q *Queue[T]) Peek()(T,bool){

	if q.head == nil {
		var zero T 
		return zero, false
	}
	return q.head.Value, true
}
func (q *Queue[T]) Size() int{return q.size}
func (q *Queue[T]) IsEmpty() bool{ return q.Size() == 0 }
func (q *Queue[T]) Clear() {
	q.head = nil
	q.tail = nil
	q.size = 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := NewQueue[int]()

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println("$" + line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "end":
			break
		case "show":
			fmt.Println(queue)
		case "push":
			for _, arg := range args[1:] {
				value, _ := strconv.Atoi(arg)
				queue.Enqueue(value)
			}
		case "pop":
			if _, ok := queue.Dequeue(); !ok {
				fmt.Println("falha: fila vazia")
			}
		case "peek":
			if value, ok := queue.Peek(); ok {
				fmt.Println(value)
			} else {
				fmt.Println("falha: fila vazia")
			}
		default:
			fmt.Println("Unknown command:", args[0])
		}
	}
}
