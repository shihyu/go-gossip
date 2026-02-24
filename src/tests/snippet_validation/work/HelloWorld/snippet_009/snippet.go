package main

import "os"
import "fmt"

func main() {
    fmt.Printf("Command: %s\n", os.Args[0])
    fmt.Printf("Hello, %s\n", os.Args[1])
}
