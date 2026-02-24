package main

import "fmt"

func main() {
    i := 10
    defer fmt.Println(i)    
    i++
    fmt.Println(i) 
}
