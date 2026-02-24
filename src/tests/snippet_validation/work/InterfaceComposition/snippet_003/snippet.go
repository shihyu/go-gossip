package main

import "fmt"

type SuperTester interface {
    stest()
}

type ParentTester interface {
    ptest()
}

type ChildTester interface {
    SuperTester
    ParentTester
    ctest()
}

type Tester interface {
    stest()
    ptest()
    ctest()
}

type Subject struct {
    name string
}

func (s *Subject) stest() {
    fmt.Printf("stest %s\n", s)
}

func (s *Subject) ptest() {
    fmt.Printf("ptest %s\n", s)
}

func (s *Subject) ctest() {
    fmt.Printf("ctest %s\n", s)
}

func main() {
    var ctester ChildTester = &Subject{"Test"}
    var tester Tester = ctester
    tester.stest()
    tester.ptest()
    tester.ctest()
}
