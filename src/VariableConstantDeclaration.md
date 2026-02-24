<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 變數宣告、常數宣告

</div>

  

變數（Variable）是儲存值的位置，變數宣告可以給予識別名稱（Identifier）、型態與初始值，在 Go 中寫下的 `10`、`3.14`、`true`、`"Justin"` 等稱之為常數（Constant），常數宣告可給予這些常數識別名稱。

# 基本變數宣告

要在 Go 中進行變數宣告有多種形式，使用 `var` 是最基本的方式。例如，宣告一個 `x` 變數，型態為 `int`，初始值為 `10`：

``` go
var x int = 10
```

這麼一來，從 `x` 這個位置開始，儲存了 `int` 長度的值 10，在宣告變數時，型態是寫在名稱之後。你也可以同時建立多個變數並指定初值：

``` go
var x, y, z int = 10, 20, 30
```

這樣的話，`x`、`y`、`z` 的型態都是 `int`，值分別是 `10`、`20`、`30`。如果宣告多個變數時，想要指定不同的型態，可以使用批量宣告：

``` go
var (
    x int = 10
    y string = "Justin"
    z bool = true
)
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

如果宣告變數時指定了型態，但未指定初值，那麼編輯器會提供預設初值，例如：

``` go
var (
    a bool        
    b int32
    c float32
    d string
    e complex128
)
```

在上面的宣告中，`a`、`b`、`c`、`d`、`e` 的值分別會是 `false`、`0`、`0.0`、`""` 與 `0 + 0i`。在 Go 中，宣告了變數，程式中卻沒有取用的動作，那麼會發生 declared and not used 的編譯錯誤。

# 自動推斷型態

在 Go 中宣告變數並指定值時，可以不用提供型態，由編譯器自動推斷型態，例如：

``` go
var x = 10
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

上頭的 `x` 型態會是 `int`，而底下的宣告：

``` go
var x, y, z = 10, 3.14, "Justin"
```

`x`、`y`、`z` 的型態分別會是 `int`、`float64` 與 `string`，批量宣告時也可以自動推斷型態，例如：

``` go
var (
    x = 10        // int 型態
    y = 3.14      // float64 型態
    z = "Justin"  // string 型態
)
```

# 短變數宣告

在函式中，想要定義變數值的場合，可以使用短變數宣告，例如：

``` go
x := 10
y := 3.14
z := "Justin"
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

如果 `x` 是首次定義，就等於是宣告變數並指定值。上例也可以寫成一行：

``` go
x, y, z := 10, 3.14, "Justin"
```

由於 Go 的函式外，每個語法都必須以關鍵字開始，因此短變數宣告不能在函式外使用。

`var` 宣告的變數名稱不可重複，然而，短變數宣告時，若同一行內有新宣告了另一變數，就可以重複宣告已存在的變數，例如，以下是合法的，因為使用 `:=` 時有一個新的 `y` 變數：

``` go
var x = 10
x, y :=  20, 30
```

此時，並沒有建立一個新的 `x` 變數，只是將新值指定給 `x` 而已。

由於短變數宣告可以同時宣告變數並指定值，因此對於底下這類需求：

``` go
package main

import "fmt"

func main() {
    var x = true
    if x {
        fmt.Println(x)
    }
}
```

在上例，`x` 的範圍是整個 `main`，若改為底下，範圍就只會是 `if` 區塊：

``` go
package main

import "fmt"

func main() {
    if x := true; x {
        fmt.Println(x)
    }
}
```

類似地，`for` 之類的語法，也常運用短變數宣告。

（在數學上 `A := B` 的寫法，涵義是藉由 B 來定義 A，例如數學上若已經定義 `x` 以及 `f(x)`，`x := f(x)` 表示用舊的 `x` 定義新的 `x`，這反而像是程式語言中的 `x = f(x)` 指定的概念，當然，數學上的符號與程式語言中的符號是有出入的，Go 在這邊只是借用了 `:=` 來作為另一種變數宣告符號。）

# 調換變數值

在 Go 中，要調換兩變數的值很簡單，例如底下的程式執行過後，`x` 會是 `20`，而 `y` 會是 `10`：

``` go
var x = 10
var y = 20
x, y = y, x
```

# 基本常數宣告

如一開始談到的，在 Go 中寫下的 `10`、`3.14`、`true`、`"Justin"` 等稱之為常數（Constant），常數宣告可給予這些常數識別名稱，要給予名稱時使用的是 `const` 關鍵字，例如：

``` go
const x = 10
```

正如〈[認識預定義型態](http://openhome.cc/Gossip/Go/PreDeclaredType.html)〉中談過的，10 會是一個整數常數，不過型態未定，如果要定義一個常數的型態，可以使用 `int32()`、`int64()` 之類的函式進行型態轉換，或者是在使用 `const` 宣告常數名稱時指定型態，例如：

``` go
const x int32 = 10
```

這邊的 `10` 就是 `int32` 型態了，注意，這邊的 `x` 並不是一個變數，而是一個識別名稱罷了，因此，會說 `x` 常數的型態是 `int32`，而不能說 `x` 變數的型態是 `int32`。

如果有多個常數要宣告，也可以批量宣告，例如：

``` go
const (
    x = 10
    y = 3.14
    z = "Justin"
)
```

再次地，`x`、`y`、`z` 分別是未定型態的整數、浮點數與字串常數（而不是 `int`、`float64`、`string` 這三個 Go 中定義的型態），如果你想要讓他們為已定義型態的整數、浮點數與字串常數，可以在宣告時指定型態：

``` go
const (
    x int = 10
    y float32 = 3.14
    z string = "Justin"
)
```

由於常數並非變數，因此，宣告了常數並不一定要用到，底下的程式不會發生錯誤：

``` go
package main

import "fmt"

func main() {
    const (
        x = 10
        y = 3.14
        z = "Justin"
    )

    fmt.Println(x)
    fmt.Println(y)
}
```

# 常數運算式

由於常數可以是未定型態，因此一個有趣的地方就是，像 `2 + 3.0`、`15 / 4`、`15 / 4.0` 這樣的常數運算式，該怎麼在編譯時期決定它們的值？答案是根據運算式中的常數運算元是整數、`rune`（單引號括住的常數）、浮點數或複數來決定，如果運算式中包括了越後面的常數，就會用它來決定。

因此，`2 + 3.0` 會是未定型態的浮點數 `5.0`，`15 / 4` 會是未定型態的整數 `3`，然而，`15 / 4.0`，會是浮點數型態的 `3.75`，在規格書的〈[Constant expressions](https://go.dev/ref/spec#Constant_expressions)〉中，列出了說明以及一些範例，例如：

``` go
const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 << 3.0         // d == 8     (untyped integer constant)
const e = 1.0 << 3         // e == 8     (untyped integer constant)
const f = int32(1) << 33   // illegal    (constant 8589934592 overflows int32)
const g = float64(2) >> 1  // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" > "bar"    // h == true  (untyped boolean constant)
const j = true             // j == true  (untyped boolean constant)
const k = 'w' + 1          // k == 'x'   (untyped rune constant)
const l = "hi"             // l == "hi"  (untyped string constant)
const m = string(k)        // m == "x"   (type string)
const Σ = 1 - 0.707i       //            (untyped complex constant)
const Δ = Σ + 2.0e-4       //            (untyped complex constant)
const Φ = iota*1i - 1/1i   //            (untyped complex constant)
```

現在，應該能明白，〈[認識預定義型態](http://openhome.cc/Gossip/Go/PreDeclaredType.html)〉中 `math.MaxInt64` 若不加上 `int64`，何以會 overflow 的錯誤了。

附帶一提的是，在 Go 中，模組中定義的名稱若要能在模組外可見，必須是首字大寫，而對於像 `math.MaxInt64` 這類的公用常數，可以定義在一個 .go 檔案之中，例如 `math.MaxInt64`，就是定義在一個 [const.go](https://go.dev/src/math/const.go) 之中。

# 使用 iota 列舉

如果要需要列舉一些常數時，可以使用 `iota`，它每遇到一次 `const`，就會重置為 `0`，若它在批量常數宣告中使用時，第一次出現時的預設值是 `0`，每出現一次就遞增 `1`，例如：

``` go
const (
    x = iota   // 0
    y = iota   // 1
    z = iota   // 2
 )
```

因為 `const` 批量宣告時，若後面的值沒寫出，會使用前一個值設定，例如：

``` go
const (
    x = 1
    y      // 1
    z      // 1
 )
```

因此，如果是連續的列舉，只要寫一次 `iota` 就可以了，這表示後續的值，也都使用 `iota`，結果就是：

``` go
const (
    x = iota   // 0
    y          // 1
    z          // 2
 )
```

其實也可以這麼寫來列舉常數，只是比較麻煩：

``` go
const x, y, z = iota, iota, iota
```

# Go 1.20+ / 1.21+ / 1.26 補充

本章談變數、常數與型別轉換，這裡補充幾個較新的語法與內建函式。

Go 1.17 起可將 slice 轉成陣列指標，Go 1.20 起也可直接轉成陣列（長度不足時會 panic）：

``` go
package main

import "fmt"

func main() {
    s := []int{10, 20, 30}
    ap := (*[3]int)(s) // Go 1.17+
    a := [3]int(s)     // Go 1.20+

    ap[0] = 99
    fmt.Println(s) // [99 20 30]
    fmt.Println(a) // [10 20 30]（a 是轉換當下的值複製）
}
```

Go 1.21 新增 `min`、`max`、`clear` 三個內建函式：

``` go
package main

import "fmt"

func main() {
    nums := []int{3, 1, 2}
    m := map[string]int{"a": 1, "b": 2}

    fmt.Println(min(3, 1, 2)) // 1
    fmt.Println(max(3, 1, 2)) // 3

    clear(nums)
    clear(m)
    fmt.Println(nums) // [0 0 0]
    fmt.Println(m)    // map[]
}
```

Go 1.26 起，`new` 的運算元可以是運算式，能直接建立並初始化指標值：

``` go
package main

import "fmt"

func main() {
    p := new(42)
    q := new(int64(300))
    fmt.Println(*p, *q) // 42 300
}
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
