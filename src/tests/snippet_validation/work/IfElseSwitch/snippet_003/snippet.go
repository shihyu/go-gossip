package main

import "fmt"

func main() {
    input := 10
    if remain := input % 2; remain == 1 {
        fmt.Printf("%d 為奇數\n", input)
    } else {
        fmt.Printf("%d 為偶數\n", input)
    }
}
