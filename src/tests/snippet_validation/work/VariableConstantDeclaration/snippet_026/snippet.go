package main

import "fmt"

func main() {
    p := new(42)
    q := new(int64(300))
    fmt.Println(*p, *q) // 42 300
}
