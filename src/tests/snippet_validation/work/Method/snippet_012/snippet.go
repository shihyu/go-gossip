package main

import (
    "fmt"
)

type Int int
type FuncInt func(Int)

func (n Int) Times(f FuncInt) {
    if n < 0 {
        panic("必須是正數")
    }

    var i Int
    for i = 0; i < n; i++ {
        f(i)
    }
}

func main() {
    var x Int = 10
    x.Times(func(n Int) {
        fmt.Println(n)
    })
}
