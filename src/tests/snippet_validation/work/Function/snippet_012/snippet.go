package main

import "fmt"

func add1To(n *int) {
    *n = *n + 1
}

func main() {
    number := 1
    add1To(&number)
    fmt.Println(number) // 2
}
