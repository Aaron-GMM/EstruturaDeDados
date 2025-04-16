package main

import (
	"container/list"
	"fmt"
)
type Produto struct{
	Nome string
	preco float64
	Estoque int
}

func criarProduto(Nome string, preco float64, Estoque int)Produto{
	produto := make([] Produto, 0)
	produto = append(produto, Produto{Nome: Nome, preco: preco, Estoque: Estoque})
	return produto[0]
}
func venderProduto(prod *Produto, quantidade int) {
    if quantidade > prod.Estoque {
        fmt.Println("Estoque insuficiente para a venda!")
    } else {
        prod.Estoque -= quantidade
        fmt.Println("Venda realizada com sucesso!")
    }
}
func mostrarProduto(prod Produto) {
	fmt.Printf("Produto:%s\n preço:%.2f\n Estoque:%v",prod.Nome, prod.preco, prod.Estoque)
}
func main(){

	produto1 := criarProduto("Notebook",1499.99,3)

	mostrarProduto(produto1)

	venderProduto(&produto1, 4)
}