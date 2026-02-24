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
    }

    for _, value := range values {
        if duck, ok := value.(Duck); ok {
            fmt.Println(duck.name)
        }
    }
} 
