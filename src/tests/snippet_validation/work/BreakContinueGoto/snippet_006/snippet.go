package main

import "fmt"

func main() {
    var input int

RETRY:
    fmt.Printf("輸入數字")
    fmt.Scanf("%d", &input)

    if input == 0 {
    TEST:
        fmt.Println("除數不得為 0")
        goto RETRY
    }
    fmt.Printf("100 / %d = %f\n", input, 100/float32(input))
    goto TEST
}
