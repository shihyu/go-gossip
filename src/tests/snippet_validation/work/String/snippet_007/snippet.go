package main

import "fmt"

func main() {
    text1 := "Go語言"
    bs := []byte(text1)
    bs[0] = 103
    text2 := string(bs)
    fmt.Println(text1) // Go語言
    fmt.Println(text2) // go語言
}
