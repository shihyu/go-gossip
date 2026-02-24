package main

import "fmt"

type Savings interface {
    Deposit(amount float64)
    Withdraw(amount float64) error
}

type Duck struct{}

func (d *Duck) Deposit(amount float64) {
    fmt.Println("我是一隻鴨子，我沒帳戶")
}

func (d *Duck) Withdraw(amount float64) error {
    fmt.Println("我是一隻鴨子，我沒錢")
    return nil
}

func main() {
    duckArray := [...]Savings{
        &Duck{},
        &Duck{},
    }

    for _, duck := range duckArray {
        duck.Deposit(1000)
    }

    duckSlice := []Savings{
        &Duck{},
        &Duck{},
    }

    for _, duck := range duckSlice {
        duck.Withdraw(500)
    }
}
