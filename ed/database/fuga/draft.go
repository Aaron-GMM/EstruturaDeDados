package main
import "fmt"
func main() {
        var positionHelic int8
        var positionPolic int8
        var positionFugit int8
        var directionFugit int8
        fmt.Scan(&positionHelic)
        fmt.Scan(&positionPolic)
        fmt.Scan(&positionFugit)
        fmt.Scan(&directionFugit)
    
       
        if directionFugit == -1{
                
        var p int8  = positionFugit       
        for i := 0; i<16; i++  {
                   if p == positionHelic  {
                        fmt.Printf("S\n")
                        return
                   } 
                   if  p == positionPolic {
                        fmt.Printf("N\n")
                        return
                          
                   }  
                p--
                if p < 0{
                        p = 15
                          }
                     
        }
                        }


        if directionFugit == 1{
              var p int8 = positionFugit 
                for i := 0; i<16; i++  {
                   if p == positionPolic  {
                        fmt.Printf("N\n") 
                        return
                      
                   } 
                   if  p == positionHelic {
                        fmt.Printf("S\n") 
                        return
                          
                   }  
                p++
                if p < 0{
                        p = 15
                }
                    
                }
        }
        
}
