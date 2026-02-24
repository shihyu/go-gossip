<div id="main" role="main" style="height: auto !important;">

<div class="header">

# for 迴圈

</div>

  

在 Go 中唯一的迴圈語法是 `for`，然而，它也擔任了一些語言中 `while` 的功能，並可搭配 `range` 來使用。

# 有分號的 for

`for` 最基本的使用形式，與 C/C++、Java 等語言類似，具有初始式、條件式、後置式三個部份，中間使用分號加以區隔，不必使用 `()` 括號包住這三個式子，同樣地，`for` 迴圈本體一定要使用 `{}`。

初始式只執行一次，通常用來宣告或初始變數，若是宣告變數，可見範圍僅在 `for` 中。第一個分號後是每次執行迴圈本體前會執行一次，且必須是 `true` 或 `false` 的結果，`true` 就會執行迴圈本體，`false` 就會結束迴圈，第二個分號後，則是每次執行完迴圈本體後會執行一次。

實際來看個 `for` 迴圈範例，在文字模式下從 1 顯示到 10：

``` go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}
```

這個程式白話讀來，就是從 `i` 等於 1，只要 `i` 小於等於 10 就執行迴圈本體（顯示 `i`），然後遞增 `i`。在介紹 `for` 迴圈時，許多書籍或文件很喜歡用的範例就是顯示九九乘法表，這邊也用這個例子來示範巢狀迴圈：

``` go
package main

import "fmt"

func main() {
    for i, j := 0, 0; i < 10; i, j = i+1, j+1 {
        fmt.Printf("%d * %d = %2d\n", i, j, i*j)
    }
}
```

`for` 中的各陳述是以分號區隔，若當中想寫兩個陳述則使用逗號區隔，例如：

``` go
package main

import "fmt"

func main() {
    for i, j := 0, 0; i < 10; i, j = i+1, j+1 {
        fmt.Printf("%d * %d = %2d\n", i, j, i*j)
    }
}
```

初始式、後置式都可以省略，不過，分號必須保留，例如：

``` go
package main

import "fmt"

func foo(i int) {
    for ; i < 10; i++ {
        fmt.Println(i)
    }
}

func multiplication_table() {
    for i, j := 2, 1; j < 10; {
        fmt.Printf("%d * %d = %2d ", i, j, i*j)
        if i == 9 {
            fmt.Println()
            j++
            i = (j+1)/j + 1
        } else {
            i++
        }
    }
}

func main() {
    foo(1)
    multiplication_table()
}
```

# 無分號的 for

在沒有初始式、後置式，只有條件式的情況，也就是 `for ; cond;` 的時候，可以只寫 `for cond`，這就是 C/C++、Java 中 `while` 迴圈的作用了：

``` go
package main

import "fmt"

func main() {
    i := 1
    for i < 10 {
        fmt.Println(i)
        i++
    }
}
```

如果想製造個無窮迴圈，在 C/C++、Java 等語言中常見寫成 `for(;;)`，在 Go 中是也可以寫 `for ;;`，因為條件式不寫預設就是 `true`，不過，可以只寫個 `for` 就可以了，底下是個很無聊的遊戲，看誰可以最久不撞到這個數字 5：

``` go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max-min) + min
}

func main() {
    for {
        number := random(1, 10)
        fmt.Println(number)
        if number == 5 {
            break
        }
        time.Sleep(time.Second)
    }
    fmt.Println("I hit 5....Orz")
}
```

在 `for` 迴圈中如果執行到 `break`，會離開迴圈本體。

# for range

Go 的 `for` 可以搭配 `range`，對 slice、陣列、`string`、`map` 和 channel（之後說明）進行迭代，`range` 視給定的形態不同，會有不同的傳回值。

對於 slice、陣列、`string`、`map`，在之前的〈[位元組構成的字串](http://openhome.cc/Gossip/Go/String.html)〉、〈[身為複合值的陣列](http://openhome.cc/Gossip/Go/Array.html)〉、〈[底層為陣列的 slice](http://openhome.cc/Gossip/Go/Slice.html)〉與〈[成對鍵值的 map](http://openhome.cc/Gossip/Go/Map.html)〉中，都有相關範例示範，這邊不再贅述。

# Go 1.22+ 與 1.23+ 的 for 補充

從 Go 1.22 開始，`for` 迴圈中宣告的變數，會在每次迭代建立新的變數實例，這修正了過去常見的閉包捕捉問題。例如：

``` go
package main

import "fmt"

func main() {
    funcs := []func(){}
    for _, v := range []string{"a", "b", "c"} {
        funcs = append(funcs, func() {
            fmt.Println(v)
        })
    }
    for _, f := range funcs {
        f()
    }
}
```

在 Go 1.22 之後，上例會依序印出 `a`、`b`、`c`（舊版本常會看到重複最後一個值）。

Go 1.22 也支援對整數直接使用 `range`，這相當於從 `0` 迭代到 `n-1`：

``` go
package main

import "fmt"

func main() {
    for i := range 5 {
        fmt.Println(i)
    }
}
```

Go 1.23 更進一步支援對 iterator function 使用 `range`，常見型態之一是 `func(func(T) bool)`：

``` go
package main

import "fmt"

func Counter(n int) func(func(int) bool) {
    return func(yield func(int) bool) {
        for i := range n {
            if !yield(i) {
                return
            }
        }
    }
}

func main() {
    for v := range Counter(3) {
        fmt.Println(v)
    }
}
```

這讓自訂容器或序列，也能自然地配合 `for range` 語法使用。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
