package main

import (
    "fmt"
    "os"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    f, err := os.Open("/tmp/dat")
    check(err)

    defer func() {
        if f != nil {
            f.Close()
        }
    }()

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)

    fmt.Printf("%d bytes: %s\n", n1, string(b1))
}
