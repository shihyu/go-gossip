package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    fmt.Print("檔案來源：")
    var filename string
    fmt.Scanf("%s", &filename)

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    io.Copy(os.Stdout, f)
}
