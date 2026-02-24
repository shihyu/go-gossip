package main

import "fmt"

func main() {
    passwds := map[string]int{"caterpillar": 123456}

    v, exists := passwds["monica"]
    fmt.Printf("%d %t\n", v, exists) // 0 false

    passwds["monica"] = 54321
    v, exists = passwds["monica"]
    fmt.Printf("%d %t\n", v, exists) // 54321 true
}
