package main

import "fmt"

type CheckingAccount struct {
    account struct {
        id      string
        name    string
        balance float64
    }
    overdraftlimit float64
}

func main() {
    checking := CheckingAccount{}
    checking.account = struct {
        id      string
        name    string
        balance float64
    }{"1234-5678", "Justin Lin", 1000}
    checking.overdraftlimit = 30000

    fmt.Println(checking)                // {{1234-5678 Justin Lin 1000} 30000}
    fmt.Println(checking.account)        // {1234-5678 Justin Lin 1000}
    fmt.Println(checking.account.name)   // Justin Lin
    fmt.Println(checking.overdraftlimit) // 30000
}
