<div id="main" role="main" style="height: auto !important;">

<div class="header">

# defer、panic、recover

</div>

  

就許多現代語言而言，例外處理機制是基本特性之一，然而，例外處理是好是壞，一直以來存在著各種不同的意見，在 Go 語言中，沒有例外處理機制，取而代之的，是運用 `defer`、`panic`、`recover` 來滿足類似的處理需求。

# defer 延遲執行

在 Go 語言中，可以使用 `defer` 指定某個函式延遲執行，那麼延遲到哪個時機？簡單來說，在函式 `return` 之前，例如：

``` prettyprint
package main

import "fmt"

func deferredFunc() {
    fmt.Println("deferredFunc")    
}

func main() {
    defer deferredFunc()
    fmt.Println("Hello, 世界")    
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

這個範例執行時，`deferredFunc()` 前加上了 `defer`，因此，會在 `main()` 函式 `return` 前執行，結果就是先顯示了 `"Hello, 世界"`，才顯示 `"deferredFunc"`。

如果有多個函式被 `defer`，那麼在函式 `return` 前，會依 `defer` 的相反順序執行，也就是 LIFO，例如：

``` prettyprint
package main

import "fmt"

func deferredFunc1() {
    fmt.Println("deferredFunc1")
}

func deferredFunc2() {
    fmt.Println("deferredFunc2")
}

func main() {
    defer deferredFunc1()
    defer deferredFunc2()
    fmt.Println("Hello, 世界")
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

由於先 `defer` 了 `deferredFunc1()`，才 `defer` 了 `deferredFunc2()`，因此執行結果會是 `"Hello, 世界"`、`"deferredFunc2"`、`"deferredFunc1"` 的顯示順序。

上頭是為了清楚表示出 `defer` 與函式的關係，實際上，你也可以寫成這樣就好：

``` prettyprint
package main

import "fmt"

func main() {
    defer fmt.Println("deffered 1")    
    defer fmt.Println("deffered 2")
    fmt.Println("Hello, 世界")    
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

執行結果會是 `"Hello, 世界"`、`"deferred 2"`、`"deferred 1"` 的顯示順序。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

有趣的一點是，被 `defer` 的函式若有接受某變數作為引數，那麼會是被 `defer` 當時的變數值，例如：

``` prettyprint
package main

import "fmt"

func main() {
    i := 10
    defer fmt.Println(i)    
    i++
    fmt.Println(i) 
}
```

在上面的例子中，會顯示 11 與 10，這是因為第一個 `fmt.Println(i)` 被 `defer` 時，保有 `i` 當時的值 10。

# 使用 defer 清除資源

那麼要用在何處？記得 `defer` 的特性是在函式 `return` 前執行，而且一定會被執行，因此，對於以下的這個程式：

``` prettyprint
package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("/tmp/dat")
    if err != nil {
        fmt.Println(err)
    } else {
        b1 := make([]byte, 5)
        n1, err := f.Read(b1)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Printf("%d bytes: %s\n", n1, string(b1))
            // 處理讀取的內容....
            f.Close()
        }
    }
}
```

這是一個讀取檔案的例子，`os.Open` 與 `f.Read` 的風格是傳回兩個值，第二個值代表著有無錯誤發生，因此程式中進行了錯誤的檢查，在沒有錯誤的情況下才進行檔案的讀取與內容處理，而最後透過 `f.Close()` 關閉檔案。

基本上，這個範例的問題在於，`f.Close()` 不一定會被執行，因為 Go 語言中還有其他展現錯誤的方式，例如使用 `panic` 函式。假設在「處理讀取的內容」過程中因為呼叫了 `panic` 來表示有錯誤發生，那麼會立即中斷函式的執行（在這個例子就是直接離開 `main` 函式），這時 `f.Close()` 就不會被執行。

你可以使用 `defer` 來執行函式的關閉：

``` prettyprint
package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("/tmp/dat")
    if err != nil {
        fmt.Println(err)
        return;
    }

    defer func() { // 延遲執行，而且函式 return 前一定會執行
        if f != nil {
            f.Close()
        }
    }()

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    if err != nil {
        fmt.Printf("%d bytes: %s\n", n1, string(b1))
        // 處理讀取的內容....
    }
}
```

這麼一來，若 `Read` 發生錯誤，最後一定會執行被 `defer` 的函式，從而保證了 `f.Close()` 一定會關閉檔案。

（就某些意義來說，`defer` 的角色類似於例外處理機制中 `finally` 的機制，將資源清除的函式，藉由 `defer` 來處理，一方面大概也是為了在程式碼閱讀上，強調出資源清除的重要性吧！）

# panic 恐慌中斷

方才稍微提過，如果在函式中執行 `panic`，那麼函式的流程就會中斷，若 A 函式呼叫了 B 函式，而 B 函式中呼叫了 `panic`，那麼 B 函式會從呼叫了 `panic` 的地方中斷，而 A 函式也會從呼叫了 B 函式的地方中斷，若有更深層的呼叫鏈，`panic` 的效應也會一路往回傳播。

（如果你有例外處理的經驗，這就相當於被拋出的例外都沒有處理的情況。）

可以將方才的範例改寫為以下：

``` prettyprint
package main

import (
    "fmt"
    "os"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    f, err := os.Open("/tmp/dat")
    check(err)

    defer func() {
        if f != nil {
            f.Close()
        }
    }()

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)

    fmt.Printf("%d bytes: %s\n", n1, string(b1))
}
```

如果在開啟檔案時，就發生了錯誤，假設這是在一個很深的呼叫層次中發生，若你直接想撰寫程式，將 `os.Open` 的 `error` 逐層傳回，那會是一件很麻煩的事，此時直接發出 `panic`，就可以達到想要的目的。

# recover 恢復流程

如果發生了 `panic`，而你必須做一些處理，可以使用 `recover`，這個函式必須在被 `defer` 的函式中執行才有效果，若在被 `defer` 的函式外執行，`recover` 一定是傳回 `nil`。

如果有設置 `defer` 函式，在發生了 `panic` 的情況下，被 `defer` 的函式一定會被執行，若當中執行了 `recover`，那麼 `panic` 就會被捕捉並作為 `recover` 的傳回值，那麼 `panic` 就不會一路往回傳播，除非你又呼叫了 `panic`。

因此，雖然 Go 語言中沒有例外處理機制，也可使用 `defer`、`panic` 與 `recover` 來進行類似的錯誤處理。例如，將上頭的範例，再修改為：

``` prettyprint
package main

import (
    "fmt"
    "os"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    f, err := os.Open("/tmp/dat")
    check(err)

    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err) // 這已經是頂層的 UI 介面了，想以自己的方式呈現錯誤
        }

        if f != nil {
            if err := f.Close(); err != nil {
                panic(err) // 示範再拋出 panic
            }
        }
    }()

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)

    fmt.Printf("%d bytes: %s\n", n1, string(b1))
}
```

在這個例子中，假設已經是最頂層的 UI 介面了，因此使用 `recover` 嘗試捕捉 `panic`，並以自己的方式呈現錯誤，附帶一題的是，關閉檔案也有可能發生錯誤，程式中也檢查了 `f.Close()`，視需求而定，你可以像這邊重新拋出 `panic`，或者也可以單純地設計一個 UI 介面來呈現錯誤。

什麼時候該用 `error`？什麼時候該用 `panic`？在 Go 的慣例中，鼓勵你使用 `error`，明確地進行錯誤檢查，然而，就如方才所言，巢狀且深層的呼叫時，使用 `panic` 會比較便於傳播錯誤，就 Go 的慣例來說，是以套件為界限，於套件之中，必要時可以使用 `panic`，而套件公開的函式，建議以 `error` 來回報錯誤，若套件公開的函式可能會收到 `panic`，建議使用 `recover` 捕捉，並轉換為 `error`。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
