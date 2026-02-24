package main

import "fmt"

type Duck struct{}

func main() {
    instances := [](interface{}){
        &Duck{},
        [...]int{1, 2, 3, 4, 5},
        map[string]int{"caterpillar": 123456, "monica": 54321},
    }

    for _, instance := range instances {
        fmt.Println(instance)
    }
}
