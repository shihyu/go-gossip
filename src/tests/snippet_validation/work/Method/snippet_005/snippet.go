package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

func String(account *Account) string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        account.id, account.name, account.balance)
}

type Point struct {
    x, y int
}

func String(point *Point) string { // String redeclared in this block 的編譯錯誤
    return fmt.Sprintf("Point{%d %d}", point.x, point.y)
}

func main() {
    account := &Account{"1234-5678", "Justin Lin", 1000}
    point := &Point{10, 20}
    fmt.Println(account.String())
    fmt.Println(point.String())
}
