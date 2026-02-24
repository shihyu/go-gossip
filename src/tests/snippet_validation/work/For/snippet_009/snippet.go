package main

import "fmt"

func Counter(n int) func(func(int) bool) {
    return func(yield func(int) bool) {
        for i := range n {
            if !yield(i) {
                return
            }
        }
    }
}

func main() {
    for v := range Counter(3) {
        fmt.Println(v)
    }
}
