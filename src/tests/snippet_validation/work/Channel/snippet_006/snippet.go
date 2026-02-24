package main

import "fmt"

func main() {
    ch := make(chan int, 1)

    ch <- 1
    select {
    case <-ch:
        fmt.Println("隨機任務 1")
    case <-ch:
        fmt.Println("隨機任務 2")
    case <-ch:
        fmt.Println("隨機任務 3")        
    }
}
