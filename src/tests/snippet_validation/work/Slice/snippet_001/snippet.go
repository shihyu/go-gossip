package main

import "fmt"

func main() {
    s1 := make([]int, 5)
    s2 := s1
    fmt.Println(s1) // [0 0 0 0 0]
    fmt.Println(s2) // [0 0 0 0 0]
    s1[0] = 1
    fmt.Println(s1) // [1 0 0 0 0]
    fmt.Println(s2) // [1 0 0 0 0]
    s2[1] = 2
    fmt.Println(s1) // [1 2 0 0 0]
    fmt.Println(s2) // [1 2 0 0 0]
}
