package main

import "fmt"

func main() {
    scores := []int{10, 20, 30}
    fmt.Println(min(10, 3, 22)) // 3
    fmt.Println(max(10, 3, 22)) // 22

    clear(scores)
    fmt.Println(scores) // [0 0 0]
}
