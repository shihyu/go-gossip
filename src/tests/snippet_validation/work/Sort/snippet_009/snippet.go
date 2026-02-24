package main

import (
    "fmt"
    "sort"
)

func main() {
    s := sort.IntSlice([]int{8, 2, 6, 3, 1, 4})
    sort.Sort(s)
    fmt.Println(s)                     // [1 2 3 4 6 8]
}
