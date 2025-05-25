package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Vetor struct {
    data []int
    size int
    capacity int
}







func NewVetor(capacity int) *Vetor {
    if capacity <= 0 {
		capacity = 1
	}
    return &Vetor{
        data:     make([]int, capacity),
        size:     0,
        capacity: capacity,
    }
}
func (v *Vetor) reserve(capacity int) {
    if capacity < v.size {
        return
    }
    novo := make([]int, capacity)
    copy(novo, v.data)
    v.data = novo
    v.capacity = capacity
}

func (v *Vetor) Read(value int){
   // fmt.Println("lendo", value)
    if value < 0  {
      return
    }
    if v.size >= v.capacity {
        v.reserve(v.capacity * 2)
    }
   // fmt.Println("inserindo", value, "na posicao", v.size)

    v.data[v.size] = value
    v.size++

}
func (v *Vetor) Show() {
    if v.size == 0 {
        fmt.Println("[]")
        return
    }
    fmt.Print("[ ")
    for i := 0; i < v.size; i++ {
        fmt.Printf("%d ", v.data[i])
    }
    fmt.Print("]")
    fmt.Println()
}
func (v *Vetor) Reverse() {
    if  v.size == 0 {
        fmt.Println("[]")
        return
    }
    
    for i, j := 0, v.size-1; i < j; i, j = i+1, j-1 {
        v.data[i], v.data[j] = v.data[j], v.data[i]
    }
    fmt.Print("[ ")
    for i := 0; i < v.size; i++ {
        fmt.Printf("%d ", v.data[i])
    }
    fmt.Print("]")
    fmt.Println()
}

func (v *Vetor) Cyclic(limite int) {
    if v.size == 0 {
        fmt.Println("[]")
        return
    }
    fmt.Print("[ ")
    for i := 0; i < limite; i++ {
        ind :=  i % v.size
        fmt.Printf("%d ", v.data[ind])
    }
    fmt.Print("]")
    fmt.Println()
}

func main() {
    var line, cmd string
    scanner := bufio.NewScanner(os.Stdin)
    vetor := NewVetor(0)
    for scanner.Scan() {
        fmt.Print("$")
        line = scanner.Text()
        args := strings.Fields(line)
        fmt.Println(line)
        if len(args) == 0 {
            continue
        }
        cmd = args[0]
        switch cmd {
        case "end":
            return
        case "reverse":
          vetor.Reverse()
        case "read":
            for _, tok := range args[1:] {
                value, _ := strconv.Atoi(tok)
                //fmt.Println(value)
                vetor.Read(value)
            }
        case "show":
            vetor.Show()
        case "cyclic":
            limit, _ := strconv.Atoi(args[1])
           vetor.Cyclic(limit)
        default:
			fmt.Println("fail: comando invalido")
        }
    }
}
