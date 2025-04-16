package main

import (
	"fmt"
	"math"
)
type Player struct{
	p1 float64
	p2 float64
}

func (p Player) calcPont() int{
	return int( math.Abs(p.p1-p.p2))
}
func (p Player) valido() bool{
	return p.p1>=10 && p.p2>=10
}
func main() {
    var quantidade int
	fmt.Scan(&quantidade)

	players := make([]Player, quantidade)

	for i := range quantidade{
		player:= &players[i]
		fmt.Scan(&player.p1,&player.p2)
	}
	//fmt.Println(players)

	melhorPlayer:= -1
	
	for idice, player := range players{
	//fmt.Println(idice,"\n",player)

		if player.valido() {
			if melhorPlayer == -1 || (player.calcPont() < players[melhorPlayer].calcPont()){
				melhorPlayer = idice 
			}
		}
	}
	if  melhorPlayer == -1  {
		fmt.Println("sem ganhador")
	}else{
		fmt.Println(melhorPlayer)
	}
}
