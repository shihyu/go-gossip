package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {  
    var point Point
    fmt.Println(point)      // {0 0}           
}
