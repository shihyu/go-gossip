package main

import "fmt"

func main() {
    funcs := []func(){}
    for _, v := range []string{"a", "b", "c"} {
        funcs = append(funcs, func() {
            fmt.Println(v)
        })
    }
    for _, f := range funcs {
        f()
    }
}
