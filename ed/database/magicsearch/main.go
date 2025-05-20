package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func mangicSearch(value int, slice []int, infer int, super int)int{
	if infer > super {
		return infer
	}
	meio := (infer + super) / 2
	if slice[meio] == value {
		return meio
	}
	if value < slice[meio] {
		return mangicSearch(value, slice, infer, meio-1)
	}
	return mangicSearch(value, slice, meio+1, super)
}

func MagicSearch(slice []int, value int)int{
	index :=  mangicSearch(value, slice, 0, len(slice)-1)
	if index < len(slice) && slice[index] == value {
		for index+1 < len(slice) && slice[index+1] == value {
			index++
		}
	}
	return index
}



func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
