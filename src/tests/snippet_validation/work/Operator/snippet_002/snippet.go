package main

import "fmt"

func main() {
    fmt.Println("AND運算：")
    fmt.Printf("0 AND 0 %5d\n", 0&1)
    fmt.Printf("0 AND 1 %5d\n", 0&1)
    fmt.Printf("1 AND 0 %5d\n", 1&0)
    fmt.Printf("1 AND 1 %5d\n", 1&1)

    fmt.Println("\nOR運算：")
    fmt.Printf("0 OR 0 %6d\n", 0|0)
    fmt.Printf("0 OR 1 %6d\n", 0|1)
    fmt.Printf("1 OR 0 %6d\n", 1|0)
    fmt.Printf("1 OR 1 %6d\n", 1|1)

    fmt.Println("\nXOR運算：")
    fmt.Printf("0 XOR 0 %5d\n", 0^0)
    fmt.Printf("0 XOR 1 %5d\n", 0^1)
    fmt.Printf("1 XOR 0 %5d\n", 1^0)
    fmt.Printf("1 XOR 1 %5d\n", 1^1)

    fmt.Println("\nAND NOT運算：")
    fmt.Printf("0 AND NOT 0 %5d\n", 0&^0)
    fmt.Printf("0 AND NOT 1 %5d\n", 0&^1)
    fmt.Printf("1 AND NOT 0 %5d\n", 1&^0)
    fmt.Printf("1 AND NOT 1 %5d\n", 1&^1)
}
