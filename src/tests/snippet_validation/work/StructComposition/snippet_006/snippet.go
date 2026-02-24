package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

type CheckingAccount struct {
    Account
    balance        float64
    overdraftlimit float64
}

func main() {
    account := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 2000, 30000}

    fmt.Println(account.balance)         // 2000
    fmt.Println(account.Account.balance) // 1000
}
