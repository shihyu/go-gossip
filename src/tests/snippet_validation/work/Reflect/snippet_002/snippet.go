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
    account := Account{"X123", "Justin Lin", 1000}
    t := reflect.TypeOf(account)
    fmt.Println(t.Kind())   // struct
    fmt.Println(t.String()) // main.Account
    /*
       底下顯示
       id string
       name string
       balance float64
    */
    for i, n := 0, t.NumField(); i < n; i++ {
        f := t.Field(i)
        fmt.Println(f.Name, f.Type)
    }
}
