package main

import "fmt"

func main() {
    passwds := map[string]int{"caterpillar": 123456}
    name := "caterpillar"
    _, exists := passwds[name]
    if exists {
        fmt.Printf("%s's password is %d\n", name, passwds[name])
    } else {
        fmt.Printf("No password for %s\n", name)
    }
}
