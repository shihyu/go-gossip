package main

import (
    "fmt"
    "container/list"
)

func printAll(lt *list.List) {
    for e := lt.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
}

func main() {
    lt := list.New()
    for i := 1; i <= 10; i++ {
        lt.PushBack(i)
    }

    printAll(lt)
}
