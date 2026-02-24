package main

import "fmt"

func main() {
    number := 75
    fmt.Println(number > 70 && number < 80)     // true
    fmt.Println(number > 80 || number < 75)     // false
    fmt.Println(!(number > 80 || number < 75))  // true
}
