package main

import "fmt"

func main() {
    input := 10
    remain := input % 2
    if remain == 1 {
        fmt.Printf("%d 為奇數\n", input)
    } else {
        fmt.Printf("%d 為偶數\n", input)
    }
}
