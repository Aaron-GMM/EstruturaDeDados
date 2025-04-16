package main

import "fmt"
type Pessoa struct{
    id int
}
func removerPessoas(filaPessoas []Pessoa, pessoasRemovidas []Pessoa ) []Pessoa{
        for _, rem := range pessoasRemovidas  {
                for i, p := range filaPessoas{
                     if p.id == rem.id{
                        filaPessoas = append(filaPessoas[:i], filaPessoas[i+1:]...)
                        break
                    }
                }
        }
        return filaPessoas
}

func main() {
    var numeroPessoas int8 = 0
    fmt.Scan(&numeroPessoas)
    filaPessoas:= make([]Pessoa, numeroPessoas)

    for i:= range numeroPessoas{
        pessoa:= &filaPessoas[i]
        fmt.Scan(&pessoa.id)  
    }
    
    var numeroPessoasFora int8 = 0
    fmt.Scan(&numeroPessoasFora)
    pessoasRemovidas:= make([]Pessoa, numeroPessoasFora)
    
    for i:= range numeroPessoasFora  {
        pessoasFora:= &pessoasRemovidas[i]
        fmt.Scan(&pessoasFora.id)
    }

    newFila:= removerPessoas(filaPessoas,pessoasRemovidas)

    for _, v := range newFila{
        fmt.Printf("%v ",v.id)
    }
    fmt.Println()

    




}
