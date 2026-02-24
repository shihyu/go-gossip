package main

import "fmt"

func main() {
    x := 10
    y := 20
    fmt.Printf("{%d %d}\n", x, y) // {10 20}

    x = 20
    y = 30
    fmt.Printf("{%d %d}\n", x, y) // {20 30}
}
