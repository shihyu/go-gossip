package main

import "syscall/js"

func main() {
    innerHTML :=
        js.Global().
            Get("document").
            Call("getElementById", "c1").
            Get("innerHTML").
            String()
    println(innerHTML)
}
