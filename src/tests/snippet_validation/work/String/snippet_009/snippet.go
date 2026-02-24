package main

import "fmt"
import "strings"

func main() {
    text := "Go語言"
    fmt.Printf("%d\n", strings.Index(text, "言"))  // 5
}
