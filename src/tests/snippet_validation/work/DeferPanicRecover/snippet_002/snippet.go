package main

import "fmt"

func deferredFunc1() {
    fmt.Println("deferredFunc1")
}

func deferredFunc2() {
    fmt.Println("deferredFunc2")
}

func main() {
    defer deferredFunc1()
    defer deferredFunc2()
    fmt.Println("Hello, 世界")
}
