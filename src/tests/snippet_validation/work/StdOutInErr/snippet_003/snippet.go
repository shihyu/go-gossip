package main

import "fmt"

func main() {
    fmt.Print("輸入空白分隔的文字")
    var text1, text2 string
    fmt.Scanln(&text1, &text2)
    fmt.Println(text1)
    fmt.Println(text2)
}
