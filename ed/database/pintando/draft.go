package main
import (
    "fmt"
    "math"
)
func main() {
    var ladoA float32  = 0.00
    var ladoB float32  = 0.00
    var ladoC float32  = 0.00

    var semiPerimetro float32 = 0.0
    var raiz float32 = 0.0

    fmt.Scanln(&ladoA)
    fmt.Scanln(&ladoB)
    fmt.Scanln(&ladoC)



    semiPerimetro = (ladoA+ladoB+ladoC)/ 2
    raiz = semiPerimetro * ((semiPerimetro-ladoA) * (semiPerimetro-ladoB) * (semiPerimetro-ladoC))
    var area float32 = fraiz(raiz) 
    
    fmt.Printf("%.2f\n", area)
}

func fraiz(raiz float32) float32{
    raiz = float32(math.Sqrt(float64(raiz)))
    return raiz

}
