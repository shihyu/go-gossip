package main

import (
    "fmt"
    "sort"
)

func main() {
    s := sort.IntSlice([]int{8, 2, 6, 3, 1, 4})
    sort.Sort(sort.Reverse(s))
    fmt.Println(s)                     // [8 6 4 3 2 1]
}
