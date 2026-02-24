package main

import "fmt"

func main() {
    nums := []int{3, 1, 2}
    m := map[string]int{"a": 1, "b": 2}

    fmt.Println(min(3, 1, 2)) // 1
    fmt.Println(max(3, 1, 2)) // 3

    clear(nums)
    clear(m)
    fmt.Println(nums) // [0 0 0]
    fmt.Println(m)    // map[]
}
