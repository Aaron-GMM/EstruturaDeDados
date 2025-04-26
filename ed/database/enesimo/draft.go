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
func encontrarPrimo(x int, contador int, numeroAtual int) int {
    if contador == x {
        return numeroAtual - 1
    }
    if eh_primo(numeroAtual, 2) {
        return encontrarPrimo(x, contador+1, numeroAtual+1)
    }
    return encontrarPrimo(x, contador, numeroAtual+1)
}
func main() {
	var x int
	fmt.Scan(&x)
    enesimo :=  encontrarPrimo(x, 0, 1)
    fmt.Println(enesimo)
}
