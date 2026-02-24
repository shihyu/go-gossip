package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) String() string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        ac.id, ac.name, ac.balance)
}

type Point struct {
    x, y int
}

func (p *Point) String() string {
    return fmt.Sprintf("Point{%d %d}", p.x, p.y)
}

func main() {
    account := &Account{"1234-5678", "Justin Lin", 1000}
    point := &Point{10, 20}
    fmt.Println(account.String()) // Account{1234-5678 Justin Lin 1000.00}
    fmt.Println(point.String())   // Point{10 20}
}
