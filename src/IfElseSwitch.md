<div id="main" role="main" style="height: auto !important;">

<div class="header">

# if..else、switch 條件式

</div>

  

在分支判斷的控制上，Go 提供了 `if...else`、`switch` 語法，相較於其他提供類似語法的語言，在 Go 中 `if...else`、`switch` 語法的相似性更高。

# if..else 語法

直接來看個 `if..else` 的實例：

``` go
package main

import "fmt"

func main() {
    input := 10
    remain := input % 2
    if remain == 1 {
        fmt.Printf("%d 為奇數\n", input)
    } else {
        fmt.Printf("%d 為偶數\n", input)
    }
}
```

在 Go 中，`if` 之後直接寫判斷式可以不用使用 `()` 括號，而 `{}` 是必要的，這樣應該是比較能避免 [Apple 曾經發生某個函式中有兩個連續縮排而引發的問題](http://support.apple.com/kb/HT6147)：

``` go
...       
if ((err = SSLHashSHA1.update(&hashCtx, &signedParams)) != 0)
        goto fail;
        goto fail;
if ((err = SSLHashSHA1.final(&hashCtx, &hashOut)) != 0)
        goto fail;
...
```

Go 的 `if` 可以使用 `:=` 宣告與指定變數值，與判斷式之間以分號區隔，因此方才的範例也可以寫成：

``` go
package main

import "fmt"

func main() {
    input := 10
    if remain := input % 2; remain == 1 {
        fmt.Printf("%d 為奇數\n", input)
    } else {
        fmt.Printf("%d 為偶數\n", input)
    }
}
```

這麼一來，`remain` 變數就只在 `if..else` 的區塊中有作用。如果要使用 `:=` 宣告與指定多個變數值，可以寫成 `if var1, var2 := 10, 20; cond` 的形式。`if...else` 可以組成 `if...else if...else` 形式，例如：

``` go
package main

import "fmt"

func main() {
    var level rune
    if score := 88; score >= 90 {
        level = 'A'
    } else if score >= 80 && score < 90 {
        level = 'B'
    } else if score >= 70 && score < 80 {
        level = 'C'
    } else if score >= 60 && score < 70 {
        level = 'D'
    } else {
        level = 'E'
    }
    fmt.Printf("得分等級：%c\n", level)
}
```

# switch 語法

實際上，對於上頭的範例，可以改用 `switch` 來撰寫，程式會更為簡潔：

``` go
package main

import "fmt"

func main() {
    var level rune
    score := 88

    switch score / 10 {
    case 10, 9:
        level = 'A'
    case 8:

        level = 'B'
    case 7:
        level = 'C'
    case 6:
        level = 'D'
    default:
        level = 'E'
    }
    fmt.Printf("得分等級：%c\n", level)
}
```

注意，與 C/C++ 或 Java 等語言不同的是，Go 的 `switch` 比對成功後，不會自動往下執行，因而不用撰寫 `break`，有多個條件想符合時，在同一 `case` 中使用逗號區隔。如果真的想在比對成功後，往下一個 `case` 中的陳述執行，可以使用 `fallthrough`，例如：

``` go
package main

import "fmt"

func main() {
    var level rune

    switch score := 100; score / 10 {
    case 10:
        fmt.Println("滿分喔！")
        fallthrough
    case 9:
        level = 'A'
    case 8:

        level = 'B'
    case 7:
        level = 'C'
    case 6:
        level = 'D'
    default:
        level = 'E'
    }
    fmt.Printf("得分等級：%c\n", level)
}
```

在上面的例子中，如果沒有 `fallthrough`，那麼只會顯示 “滿分喔！“，而不會執行 `case 9` 中的 `level = 'A'`，因此最後顯示得分等級時，不會有 A 的字眼。在上頭也可以看到，`switch` 中也可以使用 `:=` 宣告與初始變數。

實際上，Go 的 `switch` 中， `case` 不用是常數，只要 `switch` 的值型態與 `case` 比對的型態符合，也可以是個變數或運算式，甚至還可以接受布林運算式，例如：

``` go
package main

import "fmt"

func main() {
    var level rune
    score := 88
    switch {
    case score >= 90:
        level = 'A'
    case score >= 80 && score < 90:
        level = 'B'
    case score >= 70 && score < 80:
        level = 'C'
    case score >= 60 && score < 70:
        level = 'D'
    default:
        level = 'E'
    }
    fmt.Printf("得分等級：%c\n", level)
}
```

在上面的例子中，`switch` 沒有指定任何變數，此時等同於 `switch true`，這時的 `case` 可以接受布林運算式，可用來取代 `if...else if...else` 的風格。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
