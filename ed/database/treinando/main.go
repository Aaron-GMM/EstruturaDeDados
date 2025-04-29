package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	if len(vet) == 1 {
		return "["+fmt.Sprint(vet[0])+"]"
	}
			str := "["
			for i, v := range vet {
				if i> 0 {
					str += ", "
				}
				str += fmt.Sprint(v)
			}
			str += "]"
	return str
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	if len(vet) == 1 {
		return "["+fmt.Sprint(vet[0])+"]"
	}
		
			str := "["
			for i, v := range reverse(vet) {
				if i> 0 {
					str += ", "
				}
				str += fmt.Sprint(v)
			}
			str += "]"
	return str
}

// reverse: inverte os elementos do slice
func reverse(vet []int)[]int {
	for i, j := 0, len(vet)-1; i < j; i, j = i+1, j-1 {
		vet[i], vet[j] = vet[j], vet[i]
	}
	return vet
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	var soma int = 0
	if len(vet) == 0 {
		return 0
	}
	for _, v := range vet {
		soma += v
		
	}
	return soma
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	
		var soma int =1
		if len(vet) == 0 {
			return 1
		}
		for _, v := range vet {
			soma *= v
			
		}
		return soma
	
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1 
	}
	var rec func(v []int) (int, int)
    rec = func(v []int) (int, int) {
        if len(v) == 1 {
            return v[0], 0
        }
        min, idx := rec(v[1:])
        if v[0] <= min {
            return v[0], 0
        }
        return min, idx + 1
    }

    _, idx := rec(vet)
    return idx
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}