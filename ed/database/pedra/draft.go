package main
import "fmt"
func main() {
    var numeroCompetidores int8
	fmt.Scanln(&numeroCompetidores)

	distancias := make([][2] int, numeroCompetidores) 

	for i:= 0; i< int(numeroCompetidores); i++{
		fmt.Scan(&distancias[i][0],&distancias[i][1])
	}

	vencedor := -1
	melhorPont := 91

	for i := 0; i < int(numeroCompetidores); i++ {
		pedraA:= distancias[i][0]
		pedraB:= distancias[i][1]
		if (pedraA <10 || pedraB <10){
			continue
		}
		var diferenca int
		if pedraA > pedraB{
		 diferenca = pedraA-pedraB
		}else{
			diferenca = pedraB-pedraA
		}
		if diferenca < melhorPont {
			melhorPont = diferenca
			vencedor = i
		}
	}
	if vencedor != -1 {
		fmt.Println(vencedor)
	} else {
		fmt.Println("sem ganhador")
	}
}
