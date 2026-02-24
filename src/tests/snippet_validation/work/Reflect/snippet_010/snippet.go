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
    vx := reflect.ValueOf(x)
    fmt.Printf("x = %d\n", vx.Int())

    vx.SetInt(20) // panic: reflect: reflect.Value.SetInt using unaddressable value
    fmt.Printf("x = %d\n", x)
}
