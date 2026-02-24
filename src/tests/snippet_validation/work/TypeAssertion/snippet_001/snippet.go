package main

import "fmt"

type ATester interface {
    test()
}

type BTester interface {
    test()
}

type Subject struct {
    name string
}

func (s *Subject) test() {
    fmt.Println(s)
}

func main() {
    var testerA ATester = &Subject{"Test"}
    var testerB BTester = testerA
    testerA.test()
    testerB.test()
}
