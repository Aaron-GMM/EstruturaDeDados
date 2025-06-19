package main
import "fmt"

func subindo(numeroDeDegraus int)  int {

    if numeroDeDegraus == 0 {
        return 0
    }
    if numeroDeDegraus == 1 || numeroDeDegraus == 2{
        return 1
    }
    if numeroDeDegraus == 3 {
        return 2
    }
       
    return subindo(numeroDeDegraus-1) + subindo(numeroDeDegraus-3)
    
}

func main() {
    
  var numeroDeDegraus int 
  fmt.Scan(&numeroDeDegraus)

  fmt.Println(subindo(numeroDeDegraus))
    

}
