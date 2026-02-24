package main

import "fmt"

func main() {
    number := 1
    fmt.Printf("2 的 0 次方: %d\n", number)        // 1
    fmt.Printf("2 的 1 次方: %d\n", number << 1)   // 2
    fmt.Printf("2 的 2 次方: %d\n", number << 2)   // 4
    fmt.Printf("2 的 3 次方: %d\n", number << 3)   // 8
}
