<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 函式入門

</div>

  

在 Go 中要定義函式，是使用 `func` 來定義，其基本格式如下：

``` prettyprint
func funcName(param1 type1, param2 type2) (return1 type1, return2 type2) {
    // 一些程式碼...
    return value1, value2
}
```

# 定義函式

可以看到，Go 定義函式時，參數的型態宣告同樣地是放在名稱之後，如果多個參數有同樣的型態，那麼只要最右邊同型態的名稱右方加上型態就可以了，比較特別的地方在於，可以有兩個以上的傳回值，且傳回值可以設定名稱。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

來看個簡單的函式定義，以下是個求最大公因數的函式定義：

``` prettyprint
package main

import "fmt"

func Gcd(m, n int) int {
    if n == 0 {
        return m
    } else {
        return Gcd(n, m%n)
    }
}

func main() {
    fmt.Printf("Gcd of 10 and 4: %d\n", Gcd(10, 4)) // 2
}
```

當只有一個傳回值且沒有宣告名稱時，傳回值的宣告可以不用使用 `()`，傳回值的名稱可以在函式中使用，傳回值名稱設定的值，會自動於函式 `return` 時傳回，例如：

``` prettyprint
package main

import "fmt"

func Gcd(m, n int) (gcd int) {
    if n == 0 {
        gcd = m
    } else {
        gcd = Gcd(n, m%n)
    }
    return
}

func main() {
    fmt.Printf("Gcd of 10 and 4: %d\n", Gcd(10, 4)) // 2
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

官方的建議是要宣告傳回值名稱，令程式可讀性更高（當然程式會變得囉嗦一些），對那些公開給套件外使用的函式（也就是首字大寫的函式），最好是宣告傳回值名稱。

# 多個傳回值

Go 中允許多個傳回值，例如，定義一個函式，可搜尋 slice 的元素中是否指定的子字串，若有就傳回元素索引位置與字串，若無就傳回 -1 與空字串：

``` prettyprint
package main

import "fmt"
import "strings"

func FirstMatch(elems []string, substr string) (int, string) {
    for index, elem := range elems {
        if strings.Contains(elem, substr) {
            return index, elem
        }
    }
    return -1, ""
}

func main() {
    names := []string{"Justin Lin", "Monica Huang", "Irene Lin"}
    if index, name := FirstMatch(names, "Huang"); index == -1 {
        fmt.Println("找不到任何東西")
    } else {
        fmt.Printf("在索引 %d 找到 \"%s\"\n", index, name)
    }
}
```

傳回多值時，指定給變數時必須依順序，若不需要某個傳回值，可以使用 `_` 略過：

``` prettyprint
_, name := FirstMatch(names, "Huang")
```

另一種多值傳回的場合之一是錯誤處理，例如：

``` prettyprint
package main

import "fmt"
import "errors"

func Div(x, y int) (int, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
}

func main() {
    if result, err := Div(10, 5); err == nil {
        fmt.Printf("10 / 5 = %d\n", result)
    } else {
        fmt.Println(err)
    }
}
```

若函式簽署上有傳回 `error`，應透過檢查其是否為 `nil` 來確認執行時是否有錯誤發生，這是 Go 的錯誤處理風格之一，例如，`os.Open` 的函式簽署是：

``` prettyprint
func Open(name string) (file *File, err error)
```

透過 `os.Open` 開啟檔案時的一個基本範例就是：

``` prettyprint
file, err := os.Open("file.go")
if err != nil {
    log.Fatal(err)
}
```

# 可變參數

在呼叫方法時，若方法的引數個數事先無法決定該如何處理？在 Go 中支援不定長度引數（Variable-length Argument），可以輕鬆的解決這個問題。直接來看示範：

``` prettyprint
package main

import "fmt"

func Sum(numbers ...int) int {
    var sum int
    for _, number := range numbers {
        sum += number
    }
    return sum
}

func main() {
    fmt.Println(Sum(1, 2))          // 3
    fmt.Println(Sum(1, 2, 3))       // 6
    fmt.Println(Sum(1, 2, 3, 4))    // 10
    fmt.Println(Sum(1, 2, 3, 4, 5)) // 15
}
```

可以看到，要使用不定長度引數，宣告參數時要於型態關鍵字前加上 `...`，此參數本質上是個 slice，因此可以使用 `for range` 來走訪元素，可接受可變長度的參數只能有一個，而必須是最後一個參數。

雖然可接受可變長度引數的參數，本質上是個 slice，然而，若已經有個 slice，並不能直接傳遞給它，而必須使用 `...` 展開，否則會發生錯誤：

``` prettyprint
package main

import "fmt"

func Sum(numbers ...int) int {
    var sum int
    for _, number := range numbers {
        sum += number
    }
    return sum
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(Sum(numbers...)) // 15
}
```

# 函式與指標

Go 語言有指標，因此，在變數傳遞就多了一種選擇，直接來看個例子，以下的執行結果會顯示 1：

``` prettyprint
package main

import "fmt"

func add1To(n int) {
    n = n + 1
}

func main() {
    number := 1
    add1To(number)
    fmt.Println(number) // 1
}
```

這應該沒有問題，因為傳遞的是**變數值**給 `n`，函式中 `n` 的值加上 1 之後，再指定回給 `n`，這對 `main` 中的 `number` 變數毫無影響，因此函式結束後，顯示 `number` 的值，仍舊是 1。

那麼來看下面這個例子：

``` prettyprint
package main

import "fmt"

func add1To(n *int) {
    *n = *n + 1
}

func main() {
    number := 1
    add1To(&number)
    fmt.Println(number) // 2
}
```

這次使用了 `&number` 取得 `number` 的位址值再傳遞給 `n`，也就是傳遞了**變數位址值**給 `n`，函式中使用 `*n` 取得位址處的值，加上 1 後再將值存回原位址處，因此，透過 `main` 函式中的 `number` 取得的值，也會是加 1 後的值。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
