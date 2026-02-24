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

func guess(n int, ch chan int) {
    for {
        number := random(1, 10)
        ch <- number
        if number == n {
            close(ch)
        }
        time.Sleep(time.Second)
    }
}

func main() {
    ch := make(chan int)

    go guess(5, ch)

    for i := range ch {
        fmt.Println(i)
    }

    fmt.Println("I hit 5....Orz")
}
