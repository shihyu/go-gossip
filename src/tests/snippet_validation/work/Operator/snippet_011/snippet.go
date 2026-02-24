package main

import "fmt"

func main() {
    var i *int
    j := 1

    i = &j
    fmt.Println(i)  // 0x104382e0 之類的值
    fmt.Println(*i) // 1

    j = 10
    fmt.Println(*i) // 10

    *i = 20
    fmt.Println(j) // 20
}
