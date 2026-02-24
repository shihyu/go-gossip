package main

import "fmt"

func foo(i int) {
    for ; i < 10; i++ {
        fmt.Println(i)
    }
}

func multiplication_table() {
    for i, j := 2, 1; j < 10; {
        fmt.Printf("%d * %d = %2d ", i, j, i*j)
        if i == 9 {
            fmt.Println()
            j++
            i = (j+1)/j + 1
        } else {
            i++
        }
    }
}

func main() {
    foo(1)
    multiplication_table()
}
