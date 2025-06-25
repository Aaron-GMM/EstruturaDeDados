package main
import "fmt"
type Node struct {
    Value int 
    Left *Node
    Right *Node
}
type Btree struct {
    Root *Node
}
func NewNode(value int) *Node {
    return &Node{
        Value: value,
        Left:  nil,
        Right: nil,
    }
}
func NewBtree() *Btree {
    return &Btree{
        Root: nil,
    }
}
func main() {
    fmt.Println("qxcode")
}
