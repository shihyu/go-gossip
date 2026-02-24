package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3}
    for _, element := range arr {
        fmt.Printf("%d\n", element)
    }
}
