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

func Deposit(account *Account, amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    account.balance += amount
}

func Withdraw(account *Account, amount float64) error {
    if amount > account.balance {
        return errors.New("餘額不足")
    }
    account.balance -= amount
    return nil
}

func String(account *Account) string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        account.id, account.name, account.balance)
}

func main() {
    account := &Account{"1234-5678", "Justin Lin", 1000}
    Deposit(account, 500)
    Withdraw(account, 200)
    fmt.Println(String(account)) // Account{1234-5678 Justin Lin 1300.00}
}
