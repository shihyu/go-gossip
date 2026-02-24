package main

import "fmt"

type ParentTester interface {
    ptest()
}

type ChildTester interface {
    ParentTester
    ctest()
}

type Subject struct {
    name string
}

func (s *Subject) ptest() {
    fmt.Printf("ptest %s\n", s)
}

func (s *Subject) ctest() {
    fmt.Printf("ctest %s\n", s)
}

func main() {
    var tester ChildTester = &Subject{"Test"}
    tester.ptest()
    tester.ctest()
}
