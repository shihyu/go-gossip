package main

import (
    "fmt"
    "sort"
)

func main() {
    s := []int{8, 2, 6, 3, 1, 4} 
    sort.Slice(s, func(i, j int) bool {
        return s[i] > s[j]
    })
    fmt.Println(s)  // [8 6 4 3 2 1]
}
