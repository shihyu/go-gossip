package main

import (
    "golang.org/x/text/encoding/traditionalchinese"
    "golang.org/x/text/transform"
    "fmt"
)

func main() {
    big5ToUTF8 := traditionalchinese.Big5.NewDecoder()
    big5Test := "\xb4\xfa\xb8\xd5" // 測試的 Big5 編碼
    utf8, _, _ := transform.String(big5ToUTF8, big5Test)
    fmt.Println(utf8) // 顯示「測試」
}
