package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3}
    for i := 0; i < len(arr); i++ {
        fmt.Printf("%d\n", arr[i])
    }
}
