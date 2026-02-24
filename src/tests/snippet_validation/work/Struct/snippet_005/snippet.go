package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {  
    point1 := Point{X: 10, Y: 20}
    point2 := point1

    point1.X = 20

    fmt.Println(point1)  // {20, 20}
    fmt.Println(point2)  // {10 20}
}
