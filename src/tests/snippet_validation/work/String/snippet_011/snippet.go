package main

import "fmt"

func main() {
    text := "Go語言"
    for index, runeValue := range text {
        fmt.Printf("%#U 位元起始位置 %d\n", runeValue, index)
    }
}
