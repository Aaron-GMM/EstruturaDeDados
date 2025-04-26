package main
import "fmt"
func printArr(arr []int) {
    
    fmt.Print("[ ")
    for i, v := range arr {
        fmt.Print(v)
        if i != len(arr)-1 {
            fmt.Print(" ")
        }
    }
    fmt.Print(" ]\n")
}
func main() {
    var tamnho int  = 0
    var rotacao int  = 0

    fmt.Scan(&tamnho, &rotacao)
     
    var arr  = make([]int, tamnho)

    for i := range tamnho {
        fmt.Scan(&arr[i])
    }

    rotacao = rotacao % tamnho
  
    newArr := make([]int, tamnho)
    for i := 0; i < tamnho; i++ {
        np := (i+rotacao) % tamnho
        newArr[np] = arr[i]
    }
    
    printArr(newArr)
}
