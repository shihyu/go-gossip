package main

import (
    "golang.org/x/text/encoding/traditionalchinese"
    "golang.org/x/text/transform"
    "fmt"
)

func main() {
    utf8ToBig5 := traditionalchinese.Big5.NewEncoder()
    big5, _, _ := transform.String(utf8ToBig5, "測試")
    fmt.Printf("%q", big5)  // 顯示 "\xb4\xfa\xb8\xd5"
}
