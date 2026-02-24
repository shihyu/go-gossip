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

func producer(clerk chan int) {
    fmt.Println("生產者開始生產整數......")
    for product := 1; product <= 10; product++ {
        time.After(time.Duration(random(1, 5)) * time.Second)
        clerk <- product
        fmt.Printf("生產了 (%d)\n", product)
    }
}

func consumer(clerk1 chan int, clerk2 chan int) {
    fmt.Println("消費者開始消耗整數......")
    for i := 1; i <= 20; i++ {
        select {
        case p1 := <-clerk1:
            fmt.Printf("消費了生產者一的 (%d)\n", p1)
        case p2 := <-clerk2:
            fmt.Printf("消費了生產者二的 (%d)\n", p2)
        case <-time.After(3 * time.Second):
            fmt.Printf("消費者抱怨中…XD")
        }

    }
}

func main() {
    clerk1 := make(chan int)
    clerk2 := make(chan int)

    go producer(clerk1)
    go producer(clerk2)

    consumer(clerk1, clerk2)
}
