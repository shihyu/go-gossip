package main

import "fmt"

func main() {
    passwords := make(map[string]int)
    fmt.Println(passwords)      // map[]
    fmt.Println(len(passwords)) // 0

    passwords["caterpillar"] = 123456
    passwords["monica"] = 54321
    fmt.Println(passwords)                // map[caterpillar:123456 monica:54321]
    fmt.Println(len(passwords))           // 2
    fmt.Println(passwords["caterpillar"]) // 123456
    fmt.Println(passwords["monica"])      // 54321
}
