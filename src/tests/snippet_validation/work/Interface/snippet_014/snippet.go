package main

import (
    "errors"
    "fmt"
)

type Savings interface {
    Deposit(amount float64)
    Withdraw(amount float64) error
}

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

type CheckingAccount struct {
    Account
    overdraftlimit float64
}

func (ac *CheckingAccount) Withdraw(amount float64) error {
    if amount > ac.balance+ac.overdraftlimit {
        return errors.New("超出信用額度")
    }
    ac.balance -= amount
    return nil
}

func main() {
    savingsArray := [...]Savings{
        &Account{"1234-5678", "Justin Lin", 1000},
        &CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000},
    }

    for _, savings := range savingsArray {
        fmt.Println(savings)
    }

    savingsSlice := []Savings{
        &Account{"1234-5678", "Justin Lin", 1000},
        &CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000},
    }

    for _, savings := range savingsSlice {
        fmt.Println(savings)
    }
}
