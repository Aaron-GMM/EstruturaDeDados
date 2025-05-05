package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Vector struct {
	data     []int
	size     int
	capacity int
}
func (v *Vector) Reserve(capacity int) {
	if capacity < v.size{
		return
	}
	novo := make([]int, capacity)
	for i := range v.size{
		novo[i] = v.data[i]
	}
	v.capacity = capacity
	v.data = novo

}
func (v *Vector) PushBack(value int) {
	
	if v.size == v.capacity {
		v.Reserve(max(1, v.capacity*2))
	}
	v.data[v.size] = value
	v.size++
}
func (v *Vector) String() string {
	saida := "["
	for i := range v.size {
		if i != 0 {
			saida += ", "
		}
		saida += fmt.Sprint(v.data[i])
	}
	return saida + "]"
}
func (v *Vector) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
}
func NewVector(capacity int) *Vector {
	if capacity < 0 {
		capacity = 0
	}
	return &Vector{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}
func (v *Vector) GetSize() int {
	return v.size
}

func (v *Vector) PopBack()  error {
	if v.size == 0 {
		return  fmt.Errorf("vector is empty")
	}
	if len(v.data) > 0 {
		v.data = v.data[:v.size-1]
	}
	v.size--
	return  nil
}

func (v *Vector) Insert(index, value int) error {
	if index < 0 || index > v.size {
		return fmt.Errorf("index out of range")
	}
	if v.size == v.capacity {
		v.Reserve(max(1, v.capacity*2))
	}
	copy(v.data[index+1:], v.data[index:v.size])
	v.data[index] = value
	v.size++
	return nil
}

func (v *Vector) Erase(index int) error{
	if index<0 || index > v.size{
		return fmt.Errorf("index out of range")
	}
	copy(v.data[index:], v.data[index+1:v.size])
	v.size--
	return nil
}

func (v *Vector) Set(index, value int) error {
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}
	v.data[index] = value
	return nil
}
func (v *Vector) At(index int) (int, error) {
	if index < 0 || index >= v.size {
		return 0, fmt.Errorf("index out of range")
	}
	return v.data[index], nil
}

func (v *Vector) Clear() {
	v.data = make([]int, v.capacity)
	v.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

		v := NewVector(0)
		
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			 value, _ := strconv.Atoi(parts[1])
			 v = NewVector(value)
		case "push":
			 for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
			 	v.PushBack(value)
			 }
		case "show":
			 fmt.Println(v)
		case "status":
			 fmt.Println(v.Status())
		case "pop":
			 err := v.PopBack()
			 if err != nil {
			 	fmt.Println(err)
		 }
		case "insert":
			 index, _ := strconv.Atoi(parts[1])
			 value, _ := strconv.Atoi(parts[2])
			 err := v.Insert(index, value)
			 if err != nil {
			 	fmt.Println(err)
			 }
		case "erase":
			 index, _ := strconv.Atoi(parts[1])
			 err := v.Erase(index)
			 if err != nil {
			 	fmt.Println(err)
			 }
		case "indexOf":
			// value, _ := strconv.Atoi(parts[1])
			// index := v.IndexOf(value)
			// fmt.Println(index)
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
			// if v.Contains(value) {
			// 	fmt.Println("true")
			// } else {
			// 	fmt.Println("false")
			// }
		case "clear":
			 v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			 index, _ := strconv.Atoi(parts[1])
			 value, _ := strconv.Atoi(parts[2])
			 err := v.Set(index, value)
			 if err != nil {
			 	fmt.Println(err)
			 }
			 
		case "reserve":
			 newCapacity, _ := strconv.Atoi(parts[1])
			 v.Reserve(newCapacity)
		case "sort":
			// v.Sort()
		case "sorted":
			// fmt.Println("[" + Join(v.Sorted(), ", ") + "]")
		case "reverse":
			// v.Reverse()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
