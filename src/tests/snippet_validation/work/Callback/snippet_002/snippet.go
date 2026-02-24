package main

import "syscall/js"

func main() {
    // 註冊在 JavaScript 全域
    js.Global().Set("printSumTo", js.FuncOf(printSum))
    // 阻斷 main 流程
    select {}   
}

func printSum(this js.Value, args []js.Value) any {
    c1 := args[0]         // 結果顯示到這個 div 
    numbers := args[1:]   // 接下來是要加總的數字
    c1.Set("innerHTML", sum(numbers))
    return nil
}

func sum(numbers []js.Value) int {
    var sum int
    for _, val := range numbers {
        sum += val.Int()
    }
    return sum
}
