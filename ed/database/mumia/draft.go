package main
import "fmt"
func main() {
    var nome string = " "
    var idade int32 = 0
    var faixaEtaria string 
    fmt.Scanln(&nome)
    fmt.Scanln(&idade)
    switch {
	case idade < 12:
		faixaEtaria = "crianca"
	case idade < 18:
		faixaEtaria = "jovem"
	case idade < 65:
		faixaEtaria = "adulto"
    case idade<1000:
		faixaEtaria = "idoso"
    default:
         faixaEtaria = "mumia"
	}
    fmt.Println(nome+ " eh "+ faixaEtaria)
}
