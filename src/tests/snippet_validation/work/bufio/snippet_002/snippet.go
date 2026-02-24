package main

import (
    "bufio"
    "os"
    "fmt"
    "io"
)

func printFile(f *os.File) (err error){
    var (
        r = bufio.NewReader(f)
        line string
    )
    for err == nil {
        line, err = r.ReadString('\n')
        fmt.Println(line)
    }
    if err == io.EOF {
        err = nil
    }
    return
}

func main() {
    var filename string
    fmt.Print("檔案名稱：")
    fmt.Scanf("%s", &filename);

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    printFile(f)
}
