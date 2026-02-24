package main

import (
    "fmt"
    "container/ring"
)

func main() {
    numbers := ring.New(10)
    for i := 0; i < numbers.Len(); i++ {
        numbers.Value = i
        numbers = numbers.Next()
    }

    numbers.Do(func(n interface{}) {
        fmt.Printf("%d ", n.(int))
    })
}
