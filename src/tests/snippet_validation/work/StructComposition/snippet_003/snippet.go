package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

type CheckingAccount struct {
    account        Account
    overdraftlimit float64
}

func main() {
    checking := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}

    fmt.Println(checking)                // {{1234-5678 Justin Lin 1000} 30000}
    fmt.Println(checking.account)        // {1234-5678 Justin Lin 1000}
    fmt.Println(checking.account.name)   // Justin Lin
    fmt.Println(checking.overdraftlimit) // 300000
}
