package main

import "fmt"

type Point struct {
    X, Y int
}

func default_point() *Point {
    point := new(Point)
    point.X = 10
    point.Y = 10
    return point
}

func main() {
    point := default_point()
    fmt.Println(point) // &{10 10}
}
