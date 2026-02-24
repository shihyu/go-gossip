<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 型態斷言

</div>

  

宣告介面時使用的名稱，只是一個方便取用及閱讀的標示，最重要的是介面中定義的行為，以及實際的接收者型態。因此，若你打算從一個介面轉換至另一個介面，只要行為符合就可以了。例如以下是可行的：

``` prettyprint
package main

import "fmt"

type ATester interface {
    test()
}

type BTester interface {
    test()
}

type Subject struct {
    name string
}

func (s *Subject) test() {
    fmt.Println(s)
}

func main() {
    var testerA ATester = &Subject{"Test"}
    var testerB BTester = testerA
    testerA.test()
    testerB.test()
}
```

在第二個指定時，編譯器會檢查 `testerA` 的型態定義，也就是介面中，是否定義了 `test()` 行為，若是則可通過編譯，若否就編譯錯誤。例如以下的情況：

``` prettyprint
package main

import "fmt"

type ATester interface {
    testA()
}

type BTester interface {
    testB()
}

type Subject struct {
    name string
}

func (s *Subject) testA() {
    fmt.Println(s)
}

func (s *Subject) testB() {
    fmt.Println(s)
}

func main() {
    var testerA ATester = &Subject{"Test"}
    var testerB BTester = testerA // 錯誤：ATester does not implement BTester
    testerA.testA()
    testerB.testB()
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

就算 `testerA` 儲存的結構實例，確實有實作`testB()` 這個方法，然而從編譯器的角度來看，`testerA` 的行為只有 `testA()`，而看不到它有 `testB()` 的行為，因此上面這個範例會編譯錯誤。

# Comma-ok 型態斷言

如果真的要通過編譯，可以使用[型態斷言（Type assertion）](https://golang.org/ref/spec#Type_assertions)：

``` prettyprint
...同前…略

func main() {
    var testerA ATester = &Subject{"Test"}
    var testerB BTester = testerA.(BTester) 
    testerA.testA()
    testerB.testB()
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

`x.(T)` 這個語法，`x` 的型態是某介面，而 `T` 是預期的型態，或者是值實作的另一個介面名稱，在〈[介面入門](Interface.html)〉中談過，介面底層儲存了型態與值的資訊，`x.(T)` 是在告知編譯器，在執行時期再來斷言型態，也就是執行時期再來判斷 `x` 底層儲存的值，型態是否為 `T`，若是就傳回底層儲存的值。

型態斷言與型態轉換不同，型態轉換是將值的型態轉換為另一型態，編譯器會檢查兩個型態的資料結構是否相同，若否會發生編譯錯誤。

斷言是執行時期進行的，在底下的範例中，執行時期會斷言 `value` 底層儲存的值，其型態為 `Duck`：

``` prettyprint
package main

import "fmt"

type Duck struct {
    name string
}

func main() {
    values := [...](interface{}){
        Duck{"Justin"},
        Duck{"Monica"},
    }

    for _, value := range values {
        duck := value.(Duck)
        fmt.Println(duck.name)
    }
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

如果 `value` 底層儲存的值，其型態為實際上不是 `Duck` 型態，那麼操作 `duck` 時會發生執行時期錯誤，為了避免這類錯誤發生，可以進行 Comma-ok 型態斷言，例如：

``` prettyprint
package main

import "fmt"

type Duck struct {
    name string
}

func main() {
    values := [...](interface{}){
        Duck{"Justin"},
        Duck{"Monica"},
        [...]int{1, 2, 3, 4, 5},
        map[string]int{"caterpillar": 123456, "monica": 54321},
    }

    for _, value := range values {
        if duck, ok := value.(Duck); ok {
            fmt.Println(duck.name)
        }
    }
} 
```

第一個 `duck` 變數是 `Duck` 型態，若 `value` 底層儲存的值確實是 `Duck` 型態，`ok` 變數會是 `true`，否則 `ok` 會是 `false`，因此，在上面的例子中，只會針對 `Duck` 顯示其 `name` 的值。

在〈[介面入門](Interface.html)〉中談過，底下的範例會是 `false`：

``` prettyprint
var acct *Account = nil
var savings Savings = acct
fmt.Println(savings == nil) // false
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

實際上 `savings` 底層儲存的值確實是 `nil`，透過型態斷言的話可以取出。例如：

``` prettyprint
var acct *Account = nil
var savings Savings = acct
fmt.Println(savings.(*Account) == nil) // true
```

# 型態 switch 測試

依照上面的說明，如果想測試多個型態，可以用多個 `if...else if`，例如：

``` prettyprint
package main

import "fmt"

type Duck struct {
    name string
}

func main() {
    values := [...](interface{}){
        Duck{"Justin"},
        Duck{"Monica"},
        [...]int{1, 2, 3, 4, 5},
        map[string]int{"caterpillar": 123456, "monica": 54321},
        10,
    }

    for _, value := range values {
        if duck, ok := value.(Duck); ok {
            fmt.Println(duck.name)
        } else if arr, ok := value.([5]int); ok {
            fmt.Println(arr)
        } else if passwds, ok := value.(map[string]int); ok {
            fmt.Println(passwds)
        } else if i, ok := value.(int); ok {
            fmt.Println(i)
        } else {
            fmt.Println("非預期之型態")
        }
    }
}
```

不過，針對這個情況，使用型態 `switch` 測試會更為適合：

``` prettyprint
package main

import "fmt"

type Duck struct {
    name string
}

func main() {
    values := [...](interface{}){
        Duck{"Justin"},
        Duck{"Monica"},
        [...]int{1, 2, 3, 4, 5},
        map[string]int{"caterpillar": 123456, "monica": 54321},
        10,
    }

    for _, value := range values {
        switch v := value.(type) {
        case Duck:
            fmt.Println(v.name)
        case [5]int:
            fmt.Println(v[0])
        case map[string]int:
            fmt.Println(v["caterpillar"])
        case int:
            fmt.Println(v)
        default:
            fmt.Println("非預期之型態")
        }
    }
}
```

`value.(type)` 這樣的語法，只能用在 `switch` 之中。

來看個實際的應用，在 Go 的 `fmt` 中，有個 print.go 的原始碼，其中有一段是針對傳入的引數，是實作了 `Error` 介面或 `Stringer` 介面，若實作了 `Error` 介面，則呼叫其 `Error()` 方法，若實作了 `Stringer` 介面，就呼叫其 `String()` 方法：

``` prettyprint
720             switch v := p.arg.(type) {
721             case error:
722                 handled = true
723                 defer p.catchPanic(p.arg, verb)
724                 p.printArg(v.Error(), verb, depth)
725                 return
726 
727             case Stringer:
728                 handled = true
729                 defer p.catchPanic(p.arg, verb)
730                 p.printArg(v.String(), verb, depth)
731                 return
732             }
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
