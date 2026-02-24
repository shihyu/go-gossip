package main

import "fmt"

func main() {
    slice := make([]int, 5, 10)
    fmt.Println(slice)       // [0 0 0 0 0]
    fmt.Println(len(slice))  // 5
    fmt.Println(cap(slice))  // 10
}
