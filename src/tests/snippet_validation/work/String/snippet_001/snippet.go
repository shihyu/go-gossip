package main

import "fmt"

func main() {
    text1 := "Go語言"
    text2 := "Cool"
    var text3 string
    fmt.Println(text1 + text2) // Go語言Cool
    fmt.Printf("%q\n", text3)  // ""
    fmt.Println(text1 > text2) // true
}
