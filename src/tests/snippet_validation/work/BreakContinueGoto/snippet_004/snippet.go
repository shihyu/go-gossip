package main

import "fmt"

func main() {
BACK:
    for j := 1; j < 10; j++ {
        for i := 1; i < 10; i++ {
            if i == 5 {
                continue BACK
            }
            fmt.Printf("i = %d, j = %d\n", i, j)
        }
        fmt.Println("test")
    }
}
