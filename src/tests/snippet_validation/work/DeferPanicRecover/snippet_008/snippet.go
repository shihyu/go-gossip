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
        if err := recover(); err != nil {
            fmt.Println(err) // 這已經是頂層的 UI 介面了，想以自己的方式呈現錯誤
        }

        if f != nil {
            if err := f.Close(); err != nil {
                panic(err) // 示範再拋出 panic
            }
        }
    }()

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)

    fmt.Printf("%d bytes: %s\n", n1, string(b1))
}
