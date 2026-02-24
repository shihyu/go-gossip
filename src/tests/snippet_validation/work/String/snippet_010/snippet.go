package main

import "fmt"

func main() {
    text := "Go語言"
    cs := []rune(text)
    fmt.Printf("%c\n", cs[2]) // 語
    fmt.Println(len(cs))      // 4
}
