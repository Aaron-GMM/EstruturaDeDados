package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	par := make(map[int]int)
	for _, x := range vet {
		lvl := x
		if lvl < 0 {
			lvl = -lvl
		}
		par[lvl]++
	}
	keys := make([]int, 0, len(vet))
	for k := range par {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	res := make([]Pair, len(keys))
	for i, k := range keys {
		res[i] = Pair{One: k, Two: par[k]}
	}
	return res
}

func teams(vet []int) []Pair {
	res := []Pair{}
	tam := len(vet)
	for i := 0; i < tam; {
		j := i + 1
		for j < tam && vet[j] == vet[i] {
			j++
		}
		res = append(res, Pair{One: vet[i], Two: j - i})
		i = j
	}
	return res
}

func mnext(vet []int) []int {
	res := make([]int, len(vet))
	for i, x := range vet {
		if x > 0 {
			if (i > 0 && vet[i-1] < 0) || (i+1 < len(vet) && vet[i+1] < 0) {
				res[i] = 1
			}
		}
	}
	return res
}

func alone(vet []int) []int {
	res := make([]int, len(vet))
	for i, x := range vet {
		if x > 0 {
			if (i > 0 && vet[i-1] < 0) || (i+1 < len(vet) && vet[i+1] < 0) {
				continue
			} else {
				res[i] = 1
			}
		}
	}
	return res
}

func couple(vet []int) int {
	count := make(map[int]int)
	for _, x := range vet {
		count[x]++
	}
	couples := 0
	for k, v := range count {
		if k > 0 && count[-k] > 0 {
			if v < count[-k] {
				couples += v
			} else {
				couples += count[-k]
			}
		}
	}
	return couples
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	for j := 0; j < len(seq); j++ {
		if vet[pos+j] != seq[j] {
			return false
		}
	}
	return true
}

func subseq(vet []int, seq []int) int {
	if len(seq) == 0 || len(vet) < len(seq) {
		return -1
	}

	for i:= 0; i<= len(vet)-len(seq); i++{
		if  hasSubseq(vet, seq, i){ 
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	rem := make(map[int]bool)
	for _, pos := range posList {
		rem[pos] = true
	}
	newfila:=[]int{}
	for i, v := range vet {
		if !rem[i] {
			newfila = append(newfila, v)
		}
	}
	return newfila
}

func clear(vet []int, value int) []int {
	newvet := []int{}
	for _, v := range vet{
		if v != value {
			newvet = append(newvet, v)
		}
	}
	return newvet
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
