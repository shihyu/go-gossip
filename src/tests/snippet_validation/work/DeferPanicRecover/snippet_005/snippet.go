package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("/tmp/dat")
    if err != nil {
        fmt.Println(err)
    } else {
        b1 := make([]byte, 5)
        n1, err := f.Read(b1)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Printf("%d bytes: %s\n", n1, string(b1))
            // 處理讀取的內容....
            f.Close()
        }
    }
}
