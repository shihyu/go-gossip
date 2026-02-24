package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Fprintln(os.Stderr, "輸出至標準錯誤")
}
