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
    fmt.Println(IndexOf([]int{1, 2, 3}, 2))           // 1
    fmt.Println(IndexOf([]string{"Go", "C"}, "Rust")) // -1
}
