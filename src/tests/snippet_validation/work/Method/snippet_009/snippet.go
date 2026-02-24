package main

import "fmt"

type IntList []int
type Funcint func(int)

func (lt IntList) ForEach(f Funcint) {
    for _, ele := range lt {
        f(ele)
    }
}

func main() {
    var lt IntList = []int{10, 20, 30, 40, 50}
    lt.ForEach(func(ele int) {
        fmt.Println(ele)
    })
}
