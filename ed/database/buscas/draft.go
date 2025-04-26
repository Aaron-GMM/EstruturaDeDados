package main
import "fmt"
func main() {
    var  quantidadeStrings int = 0
    fmt.Scan(&quantidadeStrings)

    var stringsarr = make([]string, quantidadeStrings)
    for i := range quantidadeStrings {
        fmt.Scan(&stringsarr[i])
    }

    fmt.Scan(&quantidadeStrings)
    var stringsSearch  = make([]string, quantidadeStrings)
    for i := range quantidadeStrings {
        fmt.Scan(&stringsSearch[i])
    }
    var arrCont = make(map[string]int, quantidadeStrings)

    for _, str := range stringsarr {
        arrCont[str] ++
    }

    for i, s := range stringsSearch {
        if i >0{
            fmt.Print(" ")
        }
        fmt.Print(arrCont[s])
    }
    fmt.Println()

    // for idx, str := range stringsSearch {
    //     for _, str2 := range stringsarr {
    //         if str == str2 {
    //             arrCont[str]++
    //         }
    //     }

    //     if( idx == len(stringsSearch)-1){
    //         fmt.Printf("%v\n",  arrCont[idx])
    //     }else{
    //     fmt.Printf("%v ",  arrCont[idx])
    //     }
    // }
    
}
