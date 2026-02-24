package main

import (
    "fmt"
    "container/list"
)

type Person struct {
    Name string
    Age  int
}

func printAllPerson(persons *list.List) {
    for e := persons.Front(); e != nil; e = e.Next() {
        p := e.Value.(*Person)
        fmt.Printf("姓名：%s\t年齡：%d\n", p.Name, p.Age)
    }
}

func main() {
    persons := list.New()

    persons.PushBack(&Person{"Irene", 12})
    persons.PushBack(&Person{"Justin", 45})
    persons.PushBack(&Person{"Monica", 42})

    printAllPerson(persons)
}
