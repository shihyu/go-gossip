package main

import "fmt"

func Sum(numbers ...int) int {
    var sum int
    for _, number := range numbers {
        sum += number
    }
    return sum
}

func main() {
    fmt.Println(Sum(1, 2))          // 3
    fmt.Println(Sum(1, 2, 3))       // 6
    fmt.Println(Sum(1, 2, 3, 4))    // 10
    fmt.Println(Sum(1, 2, 3, 4, 5)) // 15
}
