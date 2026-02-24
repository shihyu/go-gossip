package main

import (
    "fmt"
    "reflect"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func main() {
    x := 10
    vx := reflect.ValueOf(&x).Elem()
    fmt.Printf("x = %d\n", vx.Int()) // x = 10

    vx.SetInt(20)
    fmt.Printf("x = %d\n", x) // x = 20
}
