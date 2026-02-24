package main

import (
    "fmt"
    "unicode/utf8"
    "unicode/utf16"
)

func main() {    
    b := []byte("Hello, 世界")

    for len(b) > 0 {
        r, size := utf8.DecodeRune(b)
        u16 := utf16.Encode([]rune{r})
        fmt.Printf("%#U:\n  Code unit %04X\n", r, u16)
        b = b[size:]
    }
}
