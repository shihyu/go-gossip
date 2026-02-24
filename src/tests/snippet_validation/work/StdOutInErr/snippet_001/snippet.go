package main

import "fmt"

func main() {
    fmt.Print("輸入名稱 年齡：")
    var name string
    var age int
    fmt.Scanf("%s %d", &name, &age)
    fmt.Printf("嗨！%s！今年 %d 歲了啊？", name, age)
}
