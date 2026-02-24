package main

import (
    "golang.org/x/text/encoding/traditionalchinese"
    "golang.org/x/text/transform"
    "fmt"
    "io"
    "os"
)

func printBig5(r io.Reader) error {
    var big5R = transform.NewReader(r, traditionalchinese.Big5.NewDecoder())

    b, err := io.ReadAll(big5R)
    fmt.Println(string(b))

    return err
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

    printBig5(f)
}
