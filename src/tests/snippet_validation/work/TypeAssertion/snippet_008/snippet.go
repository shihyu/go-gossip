package main

import "fmt"

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
        if duck, ok := value.(Duck); ok {
            fmt.Println(duck.name)
        } else if arr, ok := value.([5]int); ok {
            fmt.Println(arr)
        } else if passwds, ok := value.(map[string]int); ok {
            fmt.Println(passwds)
        } else if i, ok := value.(int); ok {
            fmt.Println(i)
        } else {
            fmt.Println("非預期之型態")
        }
    }
}
