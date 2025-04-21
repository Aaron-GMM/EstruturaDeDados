package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	newVet := make([]int, 0, len(vet))
	for _, v := range vet{
		if v >= 0 {
			newVet =  append(newVet, v)		
		}
	}
	return newVet
}

func getCalmWomen(vet []int) []int {
	calmWomen := make([]int,0, len(vet))
	
	for _, v := range vet{			
		if v<0 && -v <10{
				calmWomen = append(calmWomen, v)
		}
	}
return calmWomen
}

func sortVet(vet []int) []int {
	sort.Ints(vet)
	return vet
}

func sortStress(vet []int) []int {
	sortvet := make([]int, len(vet))
	copy(sortvet, vet)
	sort.Slice(sortvet, func (i,j int) bool  {
		ai, aj := sortvet[i], sortvet[j]
		if ai<0{
			ai = -ai
		}
		if aj < 0 {
			aj = -aj
		}
		return ai < aj
	})
	return sortvet
}

func reverse(vet []int) []int {
	rVetor := make([]int, 0, len(vet))

	for i:= len(vet); i>0; i--{
		rVetor = append(rVetor, vet[i-1])
		
	}
	
	return rVetor
}

func reverseInplace(vet []int) {
	_ = vet
}
func unique(vet []int) []int {
	 uniqueVet := make([]int, 0, len(vet))
	 controle := make(map[int]bool,len(vet))
	 for _,v := range vet{
		if !controle[v] {
			controle[v] = true
			uniqueVet = append(uniqueVet, v)
		}
	 }
	return uniqueVet
}

func repeated(vet []int) []int {
	repeatedVetor := make([]int, 0, len(vet))
	controle := make(map[int]bool,len(vet))
	for _,v := range vet{
	   if controle[v] {
		  repeatedVetor = append(repeatedVetor, v)
	   }else{
		controle[v] = true

	   }

	}
   return repeatedVetor
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			printVec(reverse(str2vet(args[1])))
		case "reverse_inplace":
			vet := str2vet(args[1])
			reverseInplace(vet)
			printVec(vet)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

