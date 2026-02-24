package main

import (
    "fmt"
    "container/ring"
)

type Person struct {
    Number int
}

func main() {
    persons := ring.New(41)
    // 給每個人編號
    for i := 1; i <= persons.Len(); i++ {
        persons.Value = &Person{i}
        persons = persons.Next()    
    }

    persons = persons.Prev()

    // 最後只留下兩人
    for persons.Len() > 2 {
        for i := 1; i <= 2; i++ {
            persons = persons.Next()
        }
        // 報數 3 Out
        persons.Unlink(1)
    }

    fmt.Print("安全位置：")
    persons.Do(func(p interface{}) {
        person := p.(*Person)
        fmt.Printf("%d ", person.Number)
    })
}
