package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    data := `Justin 45
             Monica 42
             Irene 12`
    r := strings.NewReader(data)
    var name string
    var age int
    for {
        if _, err := fmt.Fscanln(r, &name, &age); err == io.EOF {
            break
        }
        fmt.Printf("%s: %d\n", name, age)
    }
}  
