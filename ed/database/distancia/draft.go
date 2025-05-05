package main
import "fmt"

func solve(seq []rune, L , i int) bool{
    if i == len(seq) {
        return true
    }

    if seq[i] != '.' {
        return solve(seq, L, i+1)
    }

    for d:= 0; d<=L; d++ {
        r:= rune('0'+d)
        ok := true
        
    
    for j:= i-1; j >= 0 && j>= i-L; j-- {
        if seq[j] == r {
            ok = false
            break
        }
        }

    if !ok{ continue }

    for j:= i+1; j < len(seq) &&  j <= i+L; j++ {
        if seq[j] == r {
            ok = false
            break
        }
    }
    if !ok{ continue }

    seq[i] = r
    if solve(seq, L, i+1) {
        return true
    }
    seq[i] = '.'
}
    return false
}
func main() {
   var line string

   var L int 

   fmt.Scanln(&line)
   fmt.Scanln(&L)

   seq := []rune(line)
   solve(seq, L, 0)
   fmt.Println(string(seq))


}
