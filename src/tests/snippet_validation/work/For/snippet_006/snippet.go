package main

import (
    "fmt"
    "math/rand"
    "time"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max-min) + min
}

func main() {
    for {
        number := random(1, 10)
        fmt.Println(number)
        if number == 5 {
            break
        }
        time.Sleep(time.Second)
    }
    fmt.Println("I hit 5....Orz")
}
