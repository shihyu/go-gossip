package main

import "fmt"

func Gcd(m, n int) (gcd int) {
    if n == 0 {
        gcd = m
    } else {
        gcd = Gcd(n, m%n)
    }
    return
}

func main() {
    fmt.Printf("Gcd of 10 and 4: %d\n", Gcd(10, 4)) // 2
}
