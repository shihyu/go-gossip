package main

import "fmt"

type Predicate = func(int) bool

func filter(origin []int, predicate Predicate) []int {
    filtered := []int{}
    for _, elem := range origin {
        if predicate(elem) {
            filtered = append(filtered, elem)
        }
    }
    return filtered
}

func greaterThan7(n int) bool {
    return n > 7
}

func lessThan5(n int) bool {
    return n < 5
}

func main() {
    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println(filter(data, greaterThan7))
    fmt.Println(filter(data, lessThan5))    
}
