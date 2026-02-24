package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

type CheckingAccount struct {
    Account
    overdraftlimit float64
}

func main() {
    account := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}

    fmt.Println(account)                // {{1234-5678 Justin Lin 1000} 30000}
    fmt.Println(account.id)             // 1234-5678
    fmt.Println(account.name)           // Justin
    fmt.Println(account.balance)        // 1000
    fmt.Println(account.overdraftlimit) // 30000
}
