package main

import "fmt"

func deferredFunc() {
    fmt.Println("deferredFunc")    
}

func main() {
    defer deferredFunc()
    fmt.Println("Hello, 世界")    
}
