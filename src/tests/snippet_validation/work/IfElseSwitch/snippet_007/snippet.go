package main

import "fmt"

func main() {
    var level rune
    score := 88
    switch {
    case score >= 90:
        level = 'A'
    case score >= 80 && score < 90:
        level = 'B'
    case score >= 70 && score < 80:
        level = 'C'
    case score >= 60 && score < 70:
        level = 'D'
    default:
        level = 'E'
    }
    fmt.Printf("得分等級：%c\n", level)
}
