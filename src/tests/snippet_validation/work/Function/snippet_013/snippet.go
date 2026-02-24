package main

import "fmt"

func IndexOf[T comparable](elems []T, target T) int {
    for i, elem := range elems {
        if elem == target {
            return i
        }
    }
    return -1
}

func main() {
    fmt.Println(IndexOf([]int{10, 20, 30}, 20))          // 1
    fmt.Println(IndexOf([]string{"Go", "Rust"}, "Rust")) // 1
}
