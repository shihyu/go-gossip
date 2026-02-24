package main

import "fmt"

func main() {
    defer fmt.Println("deffered 1")    
    defer fmt.Println("deffered 2")
    fmt.Println("Hello, 世界")    
}
