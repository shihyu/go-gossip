package main

import (
    "fmt"
    "reflect"
)

func main() {
    s1 := []int{1, 2, 3, 4, 5}
    a1 := [...]int{1, 2, 3, 4, 5}
    fmt.Println(reflect.TypeOf(s1)) // []int
    fmt.Println(reflect.TypeOf(a1)) // [5]int
}
