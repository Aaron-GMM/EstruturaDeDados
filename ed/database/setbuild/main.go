package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Set struct {
	data [] int
	size int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}
func (s *Set) Reserve(capacity int) {
	if capacity < s.size {
		return
	}
	novo := make([]int, capacity)
	for i := range s.size {
		novo[i] = s.data[i]
	}
	s.capacity = capacity
	s.data = novo
}
func binarySearch(value int, slice []int, inf int, sup int) int{
	if sup < inf{
		return -1
	}
	meio := (inf+sup)/2
	if slice[meio] == value{
		return meio
	}

	if value < slice[meio]{
		return binarySearch(value, slice, inf, meio-1)
	} 
	return binarySearch(value, slice, meio+1, sup)
}

func BinarySearch(slice []int, value int) int{
	return binarySearch(value, slice, 0, len(slice)-1)
}

func (s *Set) insert(value int, index int) error{
	if index < 0 || index >= s.size {
		return fmt.Errorf("index out of range")
	}
	if s.size == s.capacity {
		s.Reserve(max(1, s.capacity*2))
	}
	s.data = append(s.data, 0)
	copy(s.data[index+1:], s.data[index:])
	s.data[index] = value	
	s.size++
	return nil 
}

func (s *Set) Insert(value int)  {
	index:= BinarySearch(s.data, value)
	if index != -1 {
		return
	}
	if index < s.size && s.data[index] == value {
		return
	}
	i := s.insert(value, index)
	if i == nil {
		return
	}

	// if s.size == s.capacity {
	// 	s.Reserve(max(1, s.capacity*2))
	// }
	// s.data = append(s.data, 0)
	// copy(s.data[index+1:], s.data[index:])
	// s.data[index] = value
	// s.size++

	

}




func printVet(v []int) {
	fmt.Print("[")
	for i := 0; i < len(v); i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v[i])
	}
	fmt.Println("]")
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
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
			 v = NewSet(value)
		case "insert":
			 for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			 printVet(v.data)
		case "erase":
			//  value, _ := strconv.Atoi(parts[1])
			//  v.Erase(value)
		case "contains":
			//  value, _ := strconv.Atoi(parts[1])
			//  v.Contains(value)
		case "clear":
			//  v.Clear()

		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
