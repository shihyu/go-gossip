package main

import "fmt"

func Gcd(m, n int) int {
    if n == 0 {
        return m
    } else {
        return Gcd(n, m%n)
    }
}

func main() {
    fmt.Printf("Gcd of 10 and 4: %d\n", Gcd(10, 4)) // 2
}
