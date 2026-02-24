<div id="main" role="main" style="height: auto !important;">

<div class="header">

# break、continue、goto

</div>

  

`break` 可以離開目前 `switch`、`for` 以及 `select`（之後介紹）；`continue` 只用於 `for` 迴圈，略過之後陳述句，並回到迴圈開頭進行下一次迴圈，而不是離開迴圈。`goto` 可以在函式中，讓流程直接跳至指定標籤；實際上，`break`、`continue` 在迴圈中，也可以搭配標籤來使用。

# break

在〈[if … else、switch 條件式](http://openhome.cc/Gossip/Go/IfSElsewitch.html)〉中說明過，`switch` 的 `case` 中不必特別使用 `break`，因為 `switch` 中預設不會 `fallthrough`，但 `case` 中若必要，還是可以使用 `break`，中斷 `break` 之後與下個 `case` 前的流程。

`break` 使用於 `for` 迴圈時，會結束迴圈，例如：

``` prettyprint
package main

import "fmt"

func main() {
    for i := 1; i < 10; i++ {
        if i == 5 {
            break
        }
        fmt.Printf("i = %d\n", i)
    }
}
```

這段程式會顯示 i = 1 到 i = 4，因為在 i 等於 5 時就會執行 `break` 而離開 `for` 迴圈。

`break` 可以配合標籤使用，例如本來 `break` 只會離開一層 `for` 迴圈，若設定標籤，並於 `break` 時指定標籤，就可以直接離開多層 `for` 迴圈：

``` prettyprint
package main

import "fmt"

func main() {

BACK:
    for j := 1; j < 10; j++ {
        for i := 1; i < 10; i++ {
            if i == 5 {
                break BACK
            }
            fmt.Printf("i = %d, j = %d\n", i, j)
        }
        fmt.Println("test")
    }
}
```

你可以執行看看上面的範例，之後將 `BACK:` 與 `BACK` 拿掉看看，前者 `break BACK` 時會離開兩層 `for` 迴圈，後者只會離開內層 `for` 迴圈。

# continue

`continue` 只用於 `for` 迴圈，略過之後陳述句，並回到迴圈開頭進行下一次迴圈，例如將先前第一個範例程式的 `break` 改成 `continue`：

``` prettyprint
package main

import "fmt"

func main() {
    for i := 1; i < 10; i++ {
        if i == 5 {
            continue
        }
        fmt.Printf("i = %d\n", i)
    }
}
```

這段程式會顯示 i = 1 到 4，以及 6 到 9，當 `i` 等於 5 時，會執行 `continue` 直接略過之後陳述句，也就是該次的 `fmt.Printf()` 該行並沒有被執行，直接從 `for` 開頭執行下一次迴圈，所以 i = 5 沒有被顯示。

`continue` 也有搭配標籤的用法：

``` prettyprint
package main

import "fmt"

func main() {
BACK:
    for j := 1; j < 10; j++ {
        for i := 1; i < 10; i++ {
            if i == 5 {
                continue BACK
            }
            fmt.Printf("i = %d, j = %d\n", i, j)
        }
        fmt.Println("test")
    }
}
```

# goto

在 C/C++ 中，`goto` 是一個很方便，但是常不建議使用的語法，因為濫用它的話，經常會破壞程式的架構、使得程式的邏輯混亂，然而，在 Go 中，亦有提供有 `goto` 語法。

相對於 `break` 與 `continue` 跳躍時，只能前往 `for` 迴圈開頭處設定的標籤，`goto` 可以在函式中，從某區塊內跳躍至區塊外任何位置，一個簡單的例子如下：

``` prettyprint
package main

import "fmt"

func main() {
    var input int

RETRY:
    fmt.Printf("輸入數字")
    fmt.Scanf("%d", &input)

    if input == 0 {
        fmt.Println("除數不得為 0")
        goto RETRY
    }
    fmt.Printf("100 / %d = %f\n", input, 100/float32(input))
}
```

如果你輸入 0，程式會顯示錯誤訊息後跳至 `RETRY:`，再執行一次提示與輸入。

注意，`goto` 可以在函式中，從某區塊內跳躍至區塊外任何位置，但不能從某區塊跳入另一區塊內，例如，以下是錯誤的，會發生 goto TEST jumps into block 的錯誤：

``` prettyprint
package main

import "fmt"

func main() {
    var input int

RETRY:
    fmt.Printf("輸入數字")
    fmt.Scanf("%d", &input)

    if input == 0 {
    TEST:
        fmt.Println("除數不得為 0")
        goto RETRY
    }
    fmt.Printf("100 / %d = %f\n", input, 100/float32(input))
    goto TEST
}
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
