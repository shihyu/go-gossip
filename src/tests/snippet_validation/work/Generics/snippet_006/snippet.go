package main

import "fmt"

type Adder[A Adder[A]] interface {
    Add(A) A
}

type MyInt int

func (x MyInt) Add(y MyInt) MyInt {
    return x + y
}

func Algo[A Adder[A]](x, y A) A {
    return x.Add(y)
}

func main() {
    fmt.Println(Algo(MyInt(10), MyInt(20))) // 30
}
