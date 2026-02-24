package main

import "fmt"

func main() {
    passwords := map[string]int{
        "caterpillar": 123456,
        "monica":      54321,
    }

    for name, passwd := range passwords {
        fmt.Printf("%s : %d\n", name, passwd)
    }
}
