package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) String() string {
    if ac == nil {
        return "<nil>"
    }
    return fmt.Sprintf("Account{%s %s %.2f}",
        ac.id, ac.name, ac.balance)
}

func findById(id string) *Account {
    accts := []*Account{&Account{"123", "Justin Lin", 10000}, &Account{"456", "Monica", 10000}}
    for i := 0; i < len(accts); i++ {
        if accts[i].id == id {
            return accts[i]
        }
    }
    return nil
}

func main() {
    fmt.Println(findById("123").String())
    fmt.Println(findById("789").String())
}
