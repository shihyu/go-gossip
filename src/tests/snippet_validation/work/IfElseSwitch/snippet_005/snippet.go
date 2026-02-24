package main

import "fmt"

func main() {
    var level rune
    score := 88

    switch score / 10 {
    case 10, 9:
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
