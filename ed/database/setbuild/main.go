package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int 
	size     int  
	capacity int   
}


func NewSet(capacity int) *Set {
	if capacity <= 0 {
		capacity = 1
	}
	return &Set{
		data:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}


func (s *Set) reserve(newCap int) {
	if newCap <= s.capacity {
		return
	}
	novoSlice := make([]int, s.size, newCap)
	copy(novoSlice, s.data)
	s.data = novoSlice
	s.capacity = newCap
}


func (s *Set) binarySearch(value int) (found bool, idx int) {
	low, high := 0, s.size-1
	for low <= high {
		mid := (low + high) / 2
		if s.data[mid] == value {
			return true, mid
		}
		if value < s.data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false, low
}


func (s *Set) Insert(value int) {
	found, idx := s.binarySearch(value)
	if found {
		return
	}

	if s.size == s.capacity {
	
		newCap := s.capacity * 2
		if newCap < 1 {
			newCap = 1
		}
		s.reserve(newCap)
	}

	s.data = append(s.data, 0)           
	copy(s.data[idx+1:], s.data[idx:s.size]) 
	s.data[idx] = value
	s.size++
}

func (s *Set) Contains(value int) bool {
	found, _ := s.binarySearch(value)
	return found
}


func (s *Set) Erase(value int) bool {
	found, idx := s.binarySearch(value)
	if !found {
		fmt.Println("value not found")
		return false
	}
	copy(s.data[idx:], s.data[idx+1:s.size])
	s.size--
	s.data = s.data[:s.size]
	return true
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
	scanner := bufio.NewScanner(os.Stdin)
	var s *Set

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fmt.Print("$")
		fmt.Println(line)

		parts := strings.Fields(line)
		cmd := parts[0]

		switch cmd {
		case "end":
			return

		case "init":
			if len(parts) != 2 {
				fmt.Println("fail: comando invalido")
				continue
			}
			capVal, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("fail: valor invalido")
				continue
			}
			s = NewSet(capVal)

		case "show":
			printVet(s.data)

		case "insert":
			if len(parts) < 2 {
				fmt.Println("fail: comando invalido")
				continue
			}
			for _, tok := range parts[1:] {
				v, err := strconv.Atoi(tok)
				if err != nil {
					fmt.Println("fail: valor invalido")
					continue
				}
				s.Insert(v)
			}

		case "contains":
			if len(parts) != 2 {
				fmt.Println("fail: comando invalido")
				continue
			}
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("fail: valor invalido")
				continue
			}
			if s.Contains(v) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}

		case "erase":
			if len(parts) != 2 {
				fmt.Println("fail: comando invalido")
				continue
			}
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("fail: valor invalido")
				continue
			}
			s.Erase(v)

		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
