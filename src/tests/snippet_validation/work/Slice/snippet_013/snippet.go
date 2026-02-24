package main

import "fmt"

func main() {
    src := []int{1, 2, 3, 4, 5}
    dest := make([]int, len(src), (cap(src)+1)*2)
    fmt.Println(copy(dest, src)) // 5
    fmt.Println(src)             // [1 2 3 4 5]
    fmt.Println(dest)            // [1 2 3 4 5]

    src[0] = 10
    fmt.Println(src)  // [10 2 3 4 5]
    fmt.Println(dest) // [1 2 3 4 5]
}
