package main
import "fmt"
func exibirEstado(fila []int, posicaoEspada int) {
	for i, valor := range fila {
		marcador := ""
		if i == posicaoEspada {
			marcador = ">"
		}
        if  i ==0  {
            fmt.Printf("[ ")
        }
          fmt.Printf("%d%s ", valor, marcador)
        if i == len(fila)-1{
          fmt.Printf("]")

        }
	}
	fmt.Println()
}
func main() {
    numeroPessoas:= 0 
    pessoaInicial := 0

    fmt.Scan(&numeroPessoas,&pessoaInicial)
    filaCircular := make([] int, numeroPessoas)

    pessoas:= make(map[int] int)
   
    for i := 0; i<numeroPessoas; i++ {
        filaCircular[i] = i+1
        pessoas[i] = filaCircular[i]
    }
	posicaoEspada := pessoaInicial - 1

	exibirEstado(filaCircular, posicaoEspada)
  
    
    
    for len(filaCircular) > 1 {
		indiceAlvo := (posicaoEspada + 1) % len(filaCircular)
		
		filaCircular = append(filaCircular[:indiceAlvo], filaCircular[indiceAlvo+1:]...)
		
		posicaoEspada = indiceAlvo % len(filaCircular)
		
		exibirEstado(filaCircular, posicaoEspada)
	}


    // for  _, valor := range filaCircular{
    //     indiceAlvo  = pessoaInicial
        
      
    //     if valor > pessoas[indiceAlvo] {
    //         fmt.Println("morreu oh:",pessoas[indiceAlvo])
    //         delete(pessoas,indiceAlvo)
    //         pessoaInicial = indiceAlvo+1
    //     }
        
    // } 
}
  
    
