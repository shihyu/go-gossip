package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3, 4, 5}
    slice1 := arr[0:2:4]
    fmt.Println(slice1)      // [1 2]
    fmt.Println(len(slice1)) // 2
    fmt.Println(cap(slice1)) // 4
}
