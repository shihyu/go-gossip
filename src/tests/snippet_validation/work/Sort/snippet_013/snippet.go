package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

type Family []Person

func (f Family) Len() int {
    return len(f)
}

func (f Family) Less(i, j int) bool {
    return f[i].Age < f[j].Age
}

func (f Family) Swap(i, j int) {
    f[i], f[j] = f[j], f[i]
}

func main() {
    family := Family{{"Irene", 12}, {"Justin", 45}, {"Monica", 42}}

    sort.Sort(family)
    fmt.Println(family)  // [{Irene 12} {Monica 42} {Justin 45}]

    sort.Sort(sort.Reverse(family))
    fmt.Println(family)  // [{Justin 45} {Monica 42} {Irene 12}]
}
