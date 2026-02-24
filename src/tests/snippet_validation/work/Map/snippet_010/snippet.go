package main

import "fmt"

func main() {
    passwords := map[string]int{
        "caterpillar": 123456,
        "monica":      54321,
    }

    for _, passwd := range passwords {
        fmt.Printf("%d\n", passwd)
    }
}
