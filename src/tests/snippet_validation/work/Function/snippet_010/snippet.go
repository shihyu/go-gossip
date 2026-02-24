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
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(Sum(numbers...)) // 15
}
