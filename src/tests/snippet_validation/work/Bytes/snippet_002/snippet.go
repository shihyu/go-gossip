package main

import (
    "fmt"
    "bytes"
    "strings"
)

func encodeURI(s string) string {
    buf := bytes.NewBufferString(s)

    var builder strings.Builder
    for {
        b, e := buf.ReadByte()
        if e != nil {
            break
        }
        builder.WriteString(fmt.Sprintf("%%%X", b))
    }

    return builder.String()
}

func main() {
    fmt.Println(encodeURI("良葛格")) // %E8%89%AF%E8%91%9B%E6%A0%BC
}
