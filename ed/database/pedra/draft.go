package main
import "fmt"
func main() {
    var numCompeditdor int8
    //inserir vetores das pedras 
    // vetores de competidores

    for i:= 0; i<int(numCompeditdor); i++  {
        fmt.Scan(&disPedraA)
        fmt.Scan(&disPedraB)
        if ( i==0) {
            difCompetidoUm = disPedraA-disPedraB
        }
        if(i ==1){
            difCompetidoDois = disPedraA - disPedraB
        }
    }

    if difCompetidoUm<0 {
         difCompetidoUm  *=(-1)
    }
    if difCompetidoDois<0 {
        difCompetidoDois *=(-1)
    }


}
