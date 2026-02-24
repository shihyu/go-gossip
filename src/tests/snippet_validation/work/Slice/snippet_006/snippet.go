package main

import (
    "fmt"
    "reflect"
)

func main() {
    arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    slice := arr[1:4]
    fmt.Println(reflect.TypeOf(arr))   // [10]int
    fmt.Println(reflect.TypeOf(slice)) // []int
    fmt.Println(len(slice))            // 3
    fmt.Println(cap(slice))            // 9

    fmt.Println(slice)   // [2 3 4]
    fmt.Println(arr)     // [1 2 3 4 5 6 7 8 9 10]

    slice[0] = 20
    fmt.Println(slice)   // [20 3 4]
    fmt.Println(arr)     // [1 20 3 4 5 6 7 8 9 10]
}
