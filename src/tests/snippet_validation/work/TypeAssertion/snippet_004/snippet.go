package main

import "fmt"

type Duck struct {
    name string
}

func main() {
    values := [...](interface{}){
        Duck{"Justin"},
        Duck{"Monica"},
    }

    for _, value := range values {
        duck := value.(Duck)
        fmt.Println(duck.name)
    }
}
