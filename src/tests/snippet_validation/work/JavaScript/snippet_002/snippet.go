package main

import "syscall/js"

func main() {
    window := js.Global()                       // 取得全域的 window
    doc := window.Get("document")               // 相當於 window.document
    c1 := doc.Call("getElementById", "c1")      // 相當於 document.getElementById('c1')
    innerHTML := c1.Get("innerHTML").String()   // 相當於 c1.innerHTML
    println(innerHTML)
}
