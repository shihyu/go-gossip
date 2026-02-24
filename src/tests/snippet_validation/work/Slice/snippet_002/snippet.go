package main

import "fmt"

func main() {
    var s1 []int = make([]int, 5)
    var s2 []int    // s2 這時是 nil
    s2 = s1         // 將 s1 的參考對象指定給 s2
    fmt.Println(s1) // [0 0 0 0 0]
    fmt.Println(s2) // [0 0 0 0 0]
    s1[0] = 1
    fmt.Println(s1) // [1 0 0 0 0]
    fmt.Println(s2) // [1 0 0 0 0]
    s2[1] = 2
    fmt.Println(s1) // [1 2 0 0 0]
    fmt.Println(s2) // [1 2 0 0 0]
}
