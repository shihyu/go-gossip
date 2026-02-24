package main

import (
    "fmt"
    "io"
    "os"
)

func printUTF8TC(r io.Reader) (err error) {
    var (
        buf = make([]byte, 3)
        n int
    )

    for err == nil {
        n, err = r.Read(buf)
        fmt.Print(string(buf[:n]))
    }
    if err == io.EOF {
        err = nil
    }
    return
}

func main() {
    fmt.Print("檔案來源：")
    var filename string
    fmt.Scanf("%s", &filename)

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    printUTF8TC(f)
}
