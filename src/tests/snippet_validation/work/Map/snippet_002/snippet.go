package main

import "fmt"

func main() {
    passwords := map[string]int{
        "caterpillar": 123456,
        "monica":      54321,
    }

    fmt.Println(passwords)                // map[monica:54321 caterpillar:123456]
    fmt.Println(len(passwords))           // 2
    fmt.Println(passwords["caterpillar"]) // 12345
    fmt.Println(passwords["monica"])      // 54321
}
