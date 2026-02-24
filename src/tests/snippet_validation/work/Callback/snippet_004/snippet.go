package main

import (
    "strconv"
    "syscall/js"
)

func main() {
    // 註冊按鈕事件
    dom("runButton").Call("addEventListener", "click", js.FuncOf(cal))
    select {}
}

// 根據 id 取得 DOM 物件
func dom(id string) js.Value {
    return js.Global().Get("document").Call("getElementById", id)
}

// 按下 Run 的事件處理器
func cal(this js.Value, args []js.Value) any {
    n1, _ := inputValue("n1")
    n2, _ := inputValue("n2")
    dom("r").Set("innerHTML", n1+n2)
    return nil
}

// 取得輸入欄位值
func inputValue(id string) (int, error) {
    return strconv.Atoi(dom(id).Get("value").String())
}
