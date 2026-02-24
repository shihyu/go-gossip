package main

import "fmt"

func main() {
    var scores [10]int
    scores[0] = 90
    scores[1] = 89
    fmt.Println(scores)      // [90 89 0 0 0 0 0 0 0 0]
    fmt.Println(len(scores)) // 10
}
