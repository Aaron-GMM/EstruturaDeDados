package main
import "fmt"

func __div(valor int){
    if(valor == 0){
        return
    }
    div:= valor / 2
    resto := valor %2
    __div(div)
    fmt.Printf("%v %v\n",div,resto)
}
func main() {
    valor:= 0

    fmt.Scan(&valor)
    __div(valor)
}
