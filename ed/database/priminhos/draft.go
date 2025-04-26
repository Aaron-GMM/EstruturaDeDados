package main
import "fmt"
func eh_primo(x int, div int) bool {
	if x ==0|| x == 1 || x <0 {
		return false
	}
	if div*div > x {
		return true
	}
	if x%div == 0 {
		return false
	}
	return eh_primo(x, div+1)
}
func printArr(arr []int) {
    fmt.Print("[")
    for i, v := range arr {
        fmt.Print(v)
        if i != len(arr)-1 {
            fmt.Print(", ")
        }
    }
    fmt.Println("]")
}
func primeirosPrimos(x int,contador int, numeroAtual int, primosArr []int) []int {
    if x == contador || x == 0 {
        return primosArr
    }
    if eh_primo(numeroAtual, 2) {
        primosArr = append(primosArr, numeroAtual)
        return primeirosPrimos(x, contador+1, numeroAtual+1, primosArr)
    }
    return primeirosPrimos(x, contador, numeroAtual+1, primosArr)
}
func main() {
	var x int
	fmt.Scan(&x)
    primosArr := primeirosPrimos(x, 0, 2, []int{})
    printArr(primosArr)
}