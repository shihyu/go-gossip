package main

import "fmt"

func main() {
    arr1 := [...]int{1, 2, 3}
    arr2 := arr1
    fmt.Println(arr1) // [1 2 3]
    fmt.Println(arr2) // [1 2 3]
    arr1[0] = 10
    fmt.Println(arr1) // [10 2 3]
    fmt.Println(arr2) // [1 2 3]
}
