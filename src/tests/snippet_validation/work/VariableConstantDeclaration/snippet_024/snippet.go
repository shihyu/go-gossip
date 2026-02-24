package main

import "fmt"

func main() {
    s := []int{10, 20, 30}
    ap := (*[3]int)(s) // Go 1.17+
    a := [3]int(s)     // Go 1.20+

    ap[0] = 99
    fmt.Println(s) // [99 20 30]
    fmt.Println(a) // [10 20 30]（a 是轉換當下的值複製）
}
