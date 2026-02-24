package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    slice1 := arr[:5]
    slice2 := slice1[:3]

    fmt.Println(slice1) // [1 2 3 4 5]
    fmt.Println(slice2) // [1 2 3]

    slice2[0] = 10
    fmt.Println(slice1) // [10 2 3 4 5]
    fmt.Println(slice2) // [10 2 3]
    fmt.Println(arr)    // [10 2 3 4 5 6 7 8 9 10]
}
