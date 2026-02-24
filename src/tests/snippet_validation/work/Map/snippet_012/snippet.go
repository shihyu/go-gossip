package main

import "sort"
import "fmt"

func main() {
    passwords := map[string]int{
        "caterpillar": 123456,
        "monica":      54321,
        "hamimi":      13579,
    }

    keys := make([]string, 0, len(passwords))
    for key := range passwords {
        keys = append(keys, key)
    }
    sort.Strings(keys)

    for _, key := range keys {
        fmt.Printf("%s : %d\n", key, passwords[key])
    }
}
