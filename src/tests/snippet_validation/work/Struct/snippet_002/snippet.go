package main

import "fmt"

func main() {
    point := struct{ x, y int }{10, 20}
    fmt.Printf("{%d %d}\n", point.x, point.y) // {10 20}

    point.x = 20
    point.y = 30

    fmt.Printf("{%d %d}\n", point.x, point.y) // {20 30}
}
