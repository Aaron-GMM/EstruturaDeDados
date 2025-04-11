package main
import "fmt"

type gomosPosicao struct {
    x int
    y int
}

func atualizar_cobra(cobra []gomosPosicao, direcao string){
    posiAntiga := cobra[0]
    switch direcao {
    case "L":
        cobra[0].x--
    case "R":
        cobra[0].x++
    case "U":
        cobra[0].y--
    case "D":
        cobra[0].y++
    }
    for i := 1; i < len(cobra); i++ {
        temp := cobra[i]
        cobra[i] = posiAntiga
        posiAntiga = temp
    }
}
func main() {
  
    var tamanhoCobra int
    var direcao string = ""
    fmt.Scan(&tamanhoCobra,&direcao)
    cobra := make([]gomosPosicao, tamanhoCobra)
    
    for i := 0; i<tamanhoCobra; i++{
        fmt.Scan(&cobra[i].x, &cobra[i].y)
    }
    atualizar_cobra(cobra, direcao)

    for _, pos := range cobra {
        fmt.Println(pos.x, pos.y)
    }
}
