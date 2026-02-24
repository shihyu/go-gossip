package main

import (
    "fmt"
    "sort"
)

func main() {
    s := []int{8, 2, 6, 3, 1, 4} 
    sort.Ints(s)
    fmt.Println(sort.IntsAreSorted(s)) // true
    fmt.Println(s)                     // [1 2 3 4 6 8]
    fmt.Println(sort.SearchInts(s, 7)) // 5
}
