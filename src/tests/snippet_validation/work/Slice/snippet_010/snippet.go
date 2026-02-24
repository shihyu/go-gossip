package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3, 4, 5}
    slice1 := arr[:2]
    fmt.Println(slice1)      // [1 2]
    fmt.Println(len(slice1)) // 2
    fmt.Println(cap(slice1)) // 5

    slice2 := append(slice1, 6)
    fmt.Println(slice2)      // [1 2 6]
    fmt.Println(len(slice2)) // 3
    fmt.Println(cap(slice2)) // 5

    slice2[0] = 10
    fmt.Println(slice1) // [10 2]
    fmt.Println(slice2) // [10 2 6]
    fmt.Println(arr)    // [10 2 6 4 5]
}
