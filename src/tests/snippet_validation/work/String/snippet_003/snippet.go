package main

import "fmt"

func main() {
    text := `Go語言
                 Cool`
    fmt.Printf("%q\n", text) // "Go語言\n                  Cool"
}
