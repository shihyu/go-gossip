package main

import (
    "errors"
    "fmt"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) Deposit(amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    ac.balance += amount
}

func (ac *Account) Withdraw(amount float64) error {
    if amount > ac.balance {
        return errors.New("餘額不足")
    }
    ac.balance -= amount
    return nil
}

func (ac *Account) String() string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        ac.id, ac.name, ac.balance)
}

func main() {
    account := &Account{"1234-5678", "Justin Lin", 1000}
    account.Deposit(500)
    account.Withdraw(200)
    fmt.Println(account.String()) // Account{1234-5678 Justin Lin 1300.00}
}
