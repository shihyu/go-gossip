package main

import "fmt"

type Point struct {
    X, Y int
}

type Node struct {
    point *Point
    next  *Node
}

func main() {
    node := new(Node)

    node.point = &Point{10, 20}
    node.next = new(Node)

    node.next.point = &Point{10, 30}

    fmt.Println(node.point)      // &{10 20}
    fmt.Println(node.next.point) // &{10 30}
}
