package main

import "fmt"

func main() {
    for i, j := 0, 0; i < 10; i, j = i+1, j+1 {
        fmt.Printf("%d * %d = %2d\n", i, j, i*j)
    }
}
