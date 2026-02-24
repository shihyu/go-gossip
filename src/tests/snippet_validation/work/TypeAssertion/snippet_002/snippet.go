package main

import "fmt"

type ATester interface {
    testA()
}

type BTester interface {
    testB()
}

type Subject struct {
    name string
}

func (s *Subject) testA() {
    fmt.Println(s)
}

func (s *Subject) testB() {
    fmt.Println(s)
}

func main() {
    var testerA ATester = &Subject{"Test"}
    var testerB BTester = testerA // 錯誤：ATester does not implement BTester
    testerA.testA()
    testerB.testB()
}
