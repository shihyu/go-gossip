package main

import "syscall/js"

func main() {
    alert := js.Global().Get("alert")
    // 相當於 alert('Hello, WebAssembly!')
    alert.Invoke("Hello, WebAssembly!")
}
