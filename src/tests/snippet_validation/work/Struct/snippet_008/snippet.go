package main

import "fmt"

type Point struct {
    X, Y int
}

func changeX(point *Point) {
    point.X = 20
    fmt.Printf("&{%d %d}\n", point.X, point.Y)
}

func main() {
    point := Point{X: 10, Y: 20}

    changeX(&point)    // &{20 20}
    fmt.Println(point) // {20 20}
}
