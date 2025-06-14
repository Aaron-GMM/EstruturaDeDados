package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func timesToList() *Queue[string] {
	times := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"}
	q := NewQueue[string]()
	for _, time := range times {
		q.Enqueue(time)
	}
	return q
}

func processarChaves(times *Queue[string], resultados [][2]int) string {
	for _, resultado := range resultados {
		time1 := times.Dequeue()
		time2 := times.Dequeue()

		if resultado[0] > resultado[1] {
			times.Enqueue(time1)
		} else {
			times.Enqueue(time2)
		}

	}
	return string(times.Dequeue())

}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	campeao := ""
	resultado := make([][2]int, 0, 15)

	for scanner.Scan() {
		partes := strings.Fields(scanner.Text())
		gol1, _ := strconv.Atoi(partes[0])
		gol2, _ := strconv.Atoi(partes[1])
		resultado = append(resultado, [2]int{gol1, gol2})

	}
	campeao += processarChaves(timesToList(), resultado)

	if campeao != "" {
		fmt.Println(campeao)
	} else {
		println("Não há campeão")
	}
}
