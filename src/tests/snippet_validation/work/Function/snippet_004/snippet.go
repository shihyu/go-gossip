package main

import "fmt"
import "strings"

func FirstMatch(elems []string, substr string) (int, string) {
    for index, elem := range elems {
        if strings.Contains(elem, substr) {
            return index, elem
        }
    }
    return -1, ""
}

func main() {
    names := []string{"Justin Lin", "Monica Huang", "Irene Lin"}
    if index, name := FirstMatch(names, "Huang"); index == -1 {
        fmt.Println("找不到任何東西")
    } else {
        fmt.Printf("在索引 %d 找到 \"%s\"\n", index, name)
    }
}
