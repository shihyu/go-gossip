package main

import (
    "fmt"
    "reflect"
)

type Duck struct {
    name string
}

func main() {
    values := [...](interface{}){
        Duck{"Justin"},
        Duck{"Monica"},
        [...]int{1, 2, 3, 4, 5},
        map[string]int{"caterpillar": 123456, "monica": 54321},
        10,
    }

    for _, value := range values {
        switch t := reflect.TypeOf(value); t.Kind() {
        case reflect.Struct:
            fmt.Println("it's a struct.")
        case reflect.Array:
            fmt.Println("it's a array.")
        case reflect.Map:
            fmt.Println("it's a map.")
        case reflect.Int:
            fmt.Println("it's a integer.")
        default:
            fmt.Println("非預期之型態")
        }
    }
}
