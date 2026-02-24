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
    account1 := &Account{"1234-5678", "Justin Lin", 1000}
    acct1Deposit := account1.Deposit
    acct1Withdraw := account1.Withdraw
    acct1String := account1.String

    acct1Deposit(500)
    acct1Withdraw(200)
    fmt.Println(acct1String()) // Account{1234-5678 Justin Lin 1300.00}

    account2 := &Account{"5678-1234", "Monica Huang", 500}
    acct2Deposit := account2.Deposit
    acct2Withdraw := account2.Withdraw
    acct2String := account2.String

    acct2Deposit(250)
    acct2Withdraw(150)
    fmt.Println(acct2String()) // Account{5678-1234 Monica Huang 600.00}
}
