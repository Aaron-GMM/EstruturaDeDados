package main

import (
	"fmt"
)
func main() {
    var n int 
    var qcasais int = 0
    fmt.Scanln(&n)

	casais := make([] int, n) 
   
    for i:= 0; i < n; i++{
        fmt.Scan(&casais[i])
    
    }

    for i := 0; i < n-1; i++ {
        if casais[i] == 9999 {
            continue
        }
        for j:=i+1; j<n; j++{
            if  casais[j] == 9999 {
                continue
            }
            if  casais[i]+casais[j] == 0 {
                qcasais++
                casais[i] = 9999
                casais[j] = 9999
                break
            }
        }
     
    }
    fmt.Printf("%d\n", qcasais)
}
