package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3}
    for index, element := range arr {
        fmt.Printf("%d: %d\n", index, element)
    }
}
