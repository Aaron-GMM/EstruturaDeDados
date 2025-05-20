package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}
func (ms *MultiSet) expand(newCapacity int) {
	if newCapacity < ms.size {
		return
	}
	newData := make([]int, newCapacity)
	copy(newData, ms.data)
	ms.data = newData
	ms.capacity = newCapacity
}

func (ms MultiSet) magicSearch(value int) int {
	left, right := 0, len(ms.data)-1
	for left <= right {
		mid := left + (right-left)/2
		if ms.data[mid] == value {
			return mid
		} else if ms.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func (ms *MultiSet) findInsertPos(value int) int {
	left, right := 0, ms.size
	for left < right {
		mid := (left + right) / 2
		if ms.data[mid] < value {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
func (ms *MultiSet) Insert(value int) {
	if ms.size == ms.capacity {
		newCapacity := ms.capacity * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		ms.expand(newCapacity)
	}
	pos := ms.findInsertPos(value)
	for i := ms.size; i > pos; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[pos] = value
	ms.size++

}
func (ms *MultiSet) Erase(value int) error {

	pos := ms.magicSearch(value)
	if pos == -1 {
		return fmt.Errorf("value not found")
	}
	for i := pos; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
	return nil
}

func (ms *MultiSet) Contains(value int) bool {
	return ms.magicSearch(value) != -1
}
func (ms *MultiSet) Count(value int) int {
	count := 0
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			count++
		}
	}
	return count
}
func (ms *MultiSet) Unique() int {
	uniqueCount := 0
	for i := 0; i < ms.size; i++ {
		if i == 0 || ms.data[i] != ms.data[i-1] {
			uniqueCount++
		}
	}
	return uniqueCount
}
func (ms *MultiSet) Clear() {
	ms.data = make([]int, ms.capacity)
	ms.size = 0
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return "[]"
	}
	result := "["
	result += fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	result += "]"
	return result
}

func main() {

	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)
	_ = ms
	_ = ms.magicSearch(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(Join(ms.data[:ms.size], ", "))

		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			 value, _ := strconv.Atoi(args[1])
			 count := ms.Count(value)
			 fmt.Println(count)
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
