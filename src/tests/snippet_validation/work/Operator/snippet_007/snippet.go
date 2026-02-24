package main

import "fmt"

func main() {
    fmt.Printf("10 >  5 結果 %t\n", 10 > 5)   // true
    fmt.Printf("10 >= 5 結果 %t\n", 10 >= 5)  // true
    fmt.Printf("10 <  5 結果 %t\n", 10 < 5)   // false
    fmt.Printf("10 <= 5 結果 %t\n", 10 <= 5)  // false
    fmt.Printf("10 == 5 結果 %t\n", 10 == 5)  // false
    fmt.Printf("10 != 5 結果 %t\n", 10 != 5)  // true
}
