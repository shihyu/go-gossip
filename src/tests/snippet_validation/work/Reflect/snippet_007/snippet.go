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

    account := Account{"X123", "Justin Lin", 1000}
    vacct := reflect.ValueOf(account)
    fmt.Printf("id = %s\n", vacct.FieldByName("id").String())
    fmt.Printf("name = %s\n", vacct.FieldByName("name").String())
    fmt.Printf("balance = %.2f\n", vacct.FieldByName("balance").Float())
}
