package main

import (
    "fmt"
    "unsafe"
)

func main() {
    fmt.Println(unsafe.Sizeof(true))
    fmt.Println(unsafe.Sizeof(false))
}
