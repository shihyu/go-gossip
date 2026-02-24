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
        switch v := value.(type) {
        case Duck:
            fmt.Println(v.name)
        case [5]int:
            fmt.Println(v[0])
        case map[string]int:
            fmt.Println(v["caterpillar"])
        case int:
            fmt.Println(v)
        default:
            fmt.Println("非預期之型態")
        }
    }
}
