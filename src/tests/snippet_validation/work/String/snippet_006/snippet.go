package main

import "fmt"

func main() {
    text := "Go語言"
    for i := 0; i < len(text); i++ {
        fmt.Printf("%x ", text[i])
    }
}
