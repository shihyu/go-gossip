package main

import (
    "fmt"
    "reflect"
)

func main() {
    s1 := make([]int, 5)
    s2 := s1
    fmt.Println(reflect.ValueOf(s1).Pointer() == reflect.ValueOf(s2).Pointer())
}
