package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) String() string {
    return fmt.Sprintf("Account(id = %s, name = %s, balance = %.2f)",
        ac.id, ac.name, ac.balance)
}

type CheckingAccount struct {
    Account
    overdraftlimit float64
}

func (ac *CheckingAccount) String() string {
    return fmt.Sprintf("CheckingAccount(id = %s, name = %s, balance = %.2f, overdraftlimit = %.2f)",
        ac.id, ac.name, ac.balance, ac.overdraftlimit)
}

func main() {
    account1 := Account{"1234-5678", "Justin Lin", 1000}
    account2 := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}

    // 顯示 Account(id = 1234-5678, name = Justin Lin, balance = 1000.00)
    fmt.Println(&account1)

    // 顯示 CheckingAccount(id = 1234-5678, name = Justin Lin, balance = 1000.00, overdraftlimit = 30000.00)
    fmt.Println(&account2)
}
