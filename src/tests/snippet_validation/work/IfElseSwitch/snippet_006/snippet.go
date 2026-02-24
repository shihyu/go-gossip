package main

import "fmt"

func main() {
    var level rune

    switch score := 100; score / 10 {
    case 10:
        fmt.Println("滿分喔！")
        fallthrough
    case 9:
        level = 'A'
    case 8:

        level = 'B'
    case 7:
        level = 'C'
    case 6:
        level = 'D'
    default:
        level = 'E'
    }
    fmt.Printf("得分等級：%c\n", level)
}
