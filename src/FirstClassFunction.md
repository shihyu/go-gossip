<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 一級函式

</div>

  

作為一門現代語言，Go 的特色之一是函式為一級函式（First-class function），可以作為值來進行傳遞。

# 函式作為值

例如你定義一個取最大值的函式 `max`，你可以將此函式作為值傳遞給 `maximum`：

``` go
package main

import "fmt"

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    maximum := max
    fmt.Println(max(10, 5))     // 10
    fmt.Println(maximum(10, 5)) // 10
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

可以看到，被 `max` 參考的函式，也被 `maximum` 參考著，因而，現在透過 `max` 或者 `maximum`，都可以呼叫函式。

因為 Go 型態推斷能力的關係，上頭的 `maximum` 並不用宣告型態，而可以直接參考 `max` 函式的型態，那麼，`max` 或者是 `maximum` 的型態是什麼呢？

``` go
package main

import "fmt"
import "reflect"

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    maximum := max
    fmt.Println(reflect.TypeOf(max))     // func(int, int) int
    fmt.Println(reflect.TypeOf(maximum)) // func(int, int) int
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

可以看到，函式的型態包括了 `func`、參數型態與傳回值型態，但不用宣告函式、參數與傳回值的名稱。

# 宣告函式變數

你可以僅宣告一個變數可用來參考特定型態的函式，例如：

``` go
package main

import "fmt"

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    var maximum func(int, int) int
    fmt.Println(maximum) // nil

    maximum = max
    fmt.Println(maximum(10, 5)) // 10
}
```

若想先宣告一個 `maximum` 變數，可以在之後參考 `max` 函式，可以使用型態 `func(int, int) int` 來宣告，通常，宣告函式變數時，若想免於冗長的函式型態宣告，可以使用 `type` 來定義一個新的型態名稱：

``` go
package main

import "fmt"

type BiFunc func(int, int) int // 定義了新型態

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    var maximum BiFunc
    fmt.Println(maximum) // nil

    maximum = max
    fmt.Println(maximum(10, 5)) // 10
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

在上例中，`BiFunc` 是個新的定義型態（defined type），底層型態（underlying type）為 `func(int, int) int`，Go 會認定兩者屬於不同型態，因為新的型態會擁有新的名稱，在 Go 1.9 前，這是避免冗長函式型態宣告的唯一方式。

不過，就這邊來說，實際上只是想要 `func(int, int) int` 能有個簡短一點的名稱，從 Go 1.9 開始，可以為型態取別名，別名就只是同一型態的另一個名稱，：

``` go
package main

import "fmt"

type BiFunc = func(int, int) int // 型態別名宣告

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    var maximum BiFunc
    fmt.Println(maximum) // nil

    maximum = max
    fmt.Println(maximum(10, 5)) // 10
}
```

在這邊，`BiFunc` 只是 `func(int, int) int` 的另一個名稱，而不是新的型態。

函式變數既然是個變數，也就可以對它取指標，例如：

``` go
package main

import "fmt"

type BiFunc = func(int, int) int

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    var maximum BiFunc
    fmt.Println(&maximum) // 0x1040a130
    // fmt.Println(&max)
}
```

如上，你可以對 `maximum` 取指標，得到變數位址，不過，你不能對宣告的 `max` 取指標，去除程式中最後一個註解的話，會發生 cannot take the address of max 的錯誤。

# 回呼應用

因為函式可以當作值傳遞，因此，對於函式中流程幾乎相同，只有少數操作不同的情況，就可以將操作不同的部份以回呼（Callback）函式取代。例如，可以設計一個 `filter` 函式，用來過濾出符合特定條件的值：

``` go
package main

import "fmt"

type Predicate = func(int) bool

func filter(origin []int, predicate Predicate) []int {
    filtered := []int{}
    for _, elem := range origin {
        if predicate(elem) {
            filtered = append(filtered, elem)
        }
    }
    return filtered
}

func greaterThan7(n int) bool {
    return n > 7
}

func lessThan5(n int) bool {
    return n < 5
}

func main() {
    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println(filter(data, greaterThan7))
    fmt.Println(filter(data, lessThan5))    
}
```

在這個例子中，`filter` 函式重用了 `for range` 與 `if` 等流程，只要傳入過濾用的函式，就可以讓 `filter` 具有各種的過濾用途。

除了作為值傳遞之外，Go 的函式還可以是匿名函式，且具有閉包（Closure）的特性，這將在下一篇文件加以說明。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
