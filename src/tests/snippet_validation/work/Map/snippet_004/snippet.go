package main

import "fmt"

func main() {
    passwds1 := map[string]int{"caterpillar": 123456}
    passwds2 := passwds1

    fmt.Println(passwds1) // map[caterpillar:123456]

    passwds2["monica"] = 54321
    fmt.Println(passwds1) // map[monica:54321 caterpillar:123456]
}
