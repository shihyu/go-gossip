package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    var filename string
    fmt.Print("檔案名稱：")
    fmt.Scanf("%s", &filename);

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    bufio.NewReader(f).WriteTo(os.Stdout)
}
