package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("/tmp/dat")
    if err != nil {
        fmt.Println(err)
        return;
    }

    defer func() { // 延遲執行，而且函式 return 前一定會執行
        if f != nil {
            f.Close()
        }
    }()

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    if err != nil {
        fmt.Printf("%d bytes: %s\n", n1, string(b1))
        // 處理讀取的內容....
    }
}
