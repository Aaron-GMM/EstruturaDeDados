package main
import "fmt"
func main() {
    var album int 
    var qtdFigurinhaBaruel int
    fmt.Scan(&album,&qtdFigurinhaBaruel)

    figurinhasBaruel := make([] int,qtdFigurinhaBaruel)
    figurinhasRepetidas := [] int {}
    for i := 0; i < qtdFigurinhaBaruel; i++ {
        fmt.Scan(&figurinhasBaruel[i])
    }
    sequencia := [] int {}
    qtd:= 0;
    for indice:=0; indice<album; indice++{
      qtd +=1
      sequencia = append(sequencia, qtd)
    }



    
    for i := 0; i < qtdFigurinhaBaruel-1; i++ {
      if  figurinhasBaruel[i] == figurinhasBaruel[i+1] {
          figurinhasRepetidas = append(figurinhasRepetidas, figurinhasBaruel[i])     
      }
      }
      
      faltantes := [] int {}
      for i := 0; i < len(sequencia); i++ {
        num := sequencia[i]
        encontrado := false
        for j := 0; j < qtdFigurinhaBaruel; j++ {
          if figurinhasBaruel[j] == num {
            encontrado = true
            break
          }
        }
        if !encontrado {
          faltantes = append(faltantes, num)
        }
      }

    


   if len(figurinhasRepetidas) == 0 {
    fmt.Println("N")
    } else {
     for i := 0; i < len(figurinhasRepetidas); i++ {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(figurinhasRepetidas[i])
      }
    fmt.Println()
}

  if len(faltantes) == 0 {
    fmt.Println("N")
  } else {
    for i := 0; i < len(faltantes); i++ {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(faltantes[i])
    }
    fmt.Println()
}
}
