package main

import (
    "fmt"
    "sort"
)

func main() {
    family := []struct {
        Name string
        Age  int
    } {{"Irene", 12}, {"Justin", 45}, {"Monica", 42}}

    // 依年齡由小而大排序
    sort.SliceStable(family, func(i, j int) bool {
        return family[i].Age < family[j].Age
    })

    fmt.Println(family) // [{Irene 12} {Monica 42} {Justin 45}]

    idx := sort.Search(len(family), func (i int) bool {
        return family[i].Age == 45
    })
    fmt.Println(idx)
}
