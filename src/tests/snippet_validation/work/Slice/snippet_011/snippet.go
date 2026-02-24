package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3, 4, 5}
    slice1 := arr[:]
    fmt.Println(slice1)      // [1 2 3 4 5]
    fmt.Println(len(slice1)) // 5
    fmt.Println(cap(slice1)) // 5

    slice2 := append(slice1, 6)
    fmt.Println(slice2)      // [1 2 3 4 5 6]
    fmt.Println(len(slice2)) // 6
    fmt.Println(cap(slice2)) // 12

    slice2[0] = 10
    fmt.Println(slice1) // [1 2 3 4 5]
    fmt.Println(slice2) // [10 2 3 4 5 6]
    fmt.Println(arr)    // [1 2 3 4 5]
}
