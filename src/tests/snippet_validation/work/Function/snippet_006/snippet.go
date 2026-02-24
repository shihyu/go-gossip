package main

import "fmt"
import "errors"

func Div(x, y int) (int, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
}

func main() {
    if result, err := Div(10, 5); err == nil {
        fmt.Printf("10 / 5 = %d\n", result)
    } else {
        fmt.Println(err)
    }
}
