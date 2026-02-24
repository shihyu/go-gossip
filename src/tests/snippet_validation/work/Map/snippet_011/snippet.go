package main

import "fmt"

func keys(m map[string]int) []string {
    ks := make([]string, 0, len(m))
    for k := range m {
        ks = append(ks, k)
    }
    return ks
}

func values(m map[string]int) []int {
    vs := make([]int, 0, len(m))
    for _, v := range m {
        vs = append(vs, v)
    }
    return vs
}

func main() {
    passwords := map[string]int{
        "caterpillar": 123456,
        "monica":      54321,
    }

    fmt.Println(keys(passwords))   // [caterpillar monica]
    fmt.Println(values(passwords)) // [123456 54321]
}
