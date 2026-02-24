package main

import "fmt"

func First[T any](elems []T) (T, bool) {
    if len(elems) == 0 {
        var zero T
        return zero, false
    }
    return elems[0], true
}

func main() {
    if v, ok := First([]int{10, 20, 30}); ok {
        fmt.Println(v) // 10
    }
    if v, ok := First([]string{"Go", "Rust"}); ok {
        fmt.Println(v) // Go
    }
}
