package main

import "fmt"

func producer(clerk chan<- int) {
    fmt.Println("生產者開始生產整數......")
    for product := 1; product <= 10; product++ {
        clerk <- product
        fmt.Printf("生產了 (%d)\n", product)
    }
}

func consumer(clerk <-chan int) {
    fmt.Println("消費者開始消耗整數......")
    for i := 1; i <= 10; i++ {
        fmt.Printf("消費了 (%d)\n", <-clerk)
    }
}

func main() {
    clerk := make(chan int, 2)

    go producer(clerk)
    consumer(clerk)
}   
