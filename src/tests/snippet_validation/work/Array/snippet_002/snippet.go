package main

import "fmt"

func main() {
    arr1 := [3]int{1, 2, 3}
    arr2 := [5]int{1, 2, 3}
    arr3 := [...]int{1, 2, 3, 4, 5}
    fmt.Println(arr1) // [1 2 3]
    fmt.Println(arr2) // [1 2 3 0 0]
    fmt.Println(arr3) // [1 2 3 4 5]
}
