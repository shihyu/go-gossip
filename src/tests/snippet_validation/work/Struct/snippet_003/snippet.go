package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {
    point1 := Point{10, 20}
    fmt.Println(point1) // {10 20}

    point2 := Point{Y: 20, X: 30}
    fmt.Println(point2) // {30 20}
} 
