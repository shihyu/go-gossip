<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 成對鍵值的 map

</div>

  

許多語言中都會有的成對鍵值資料結構，在 Go 中是以內建型態 `map` 來實作，格式為 `map[keyType]valueType`。

# 建立與初始 map

想要建立例一個 `map` 實例，但尚無任何鍵值對，可以使用 `make` 函式，例如：

``` prettyprint
package main

import "fmt"

func main() {
    passwords := make(map[string]int)
    fmt.Println(passwords)      // map[]
    fmt.Println(len(passwords)) // 0

    passwords["caterpillar"] = 123456
    passwords["monica"] = 54321
    fmt.Println(passwords)                // map[caterpillar:123456 monica:54321]
    fmt.Println(len(passwords))           // 2
    fmt.Println(passwords["caterpillar"]) // 123456
    fmt.Println(passwords["monica"])      // 54321
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

在上例中，`passwords` 是個參考，指向 `make(map[string]int)` 建立的 map 實例。

類似一些語言（例如 Python），要設定一個鍵值對應時，使用 `[]` 與 `=` 指定，要取得鍵對應的值時，使用 `[]` 指定鍵，這會傳回對應的值，想知道 `map` 中的鍵數，可以使用 `len` 函式。

要注意的是，可用來做為鍵的值，必須是 [comparable](https://go.dev/blog/go-maps-in-action)，就目前來說，你要先知道的 comparable 型態有布林、數字、字串、指標（pointer）、channel、interface、struct，或者含有這這些型態的陣列，這些是都可以使用 `==` 來比較的值；而 slice、map 與函式，就不能用來做為鍵。

如果已知 map 中會有的鍵值對，則可以如下建立 map：

``` prettyprint
package main

import "fmt"

func main() {
    passwords := map[string]int{
        "caterpillar": 123456,
        "monica":      54321,
    }

    fmt.Println(passwords)                // map[monica:54321 caterpillar:123456]
    fmt.Println(len(passwords))           // 2
    fmt.Println(passwords["caterpillar"]) // 12345
    fmt.Println(passwords["monica"])      // 54321
}
```

如果 `passwords` 建立時，最後一個鍵值項目後不換行，那麼最後一個逗號就不需要，例如：

``` prettyprint
passwords := map[string]int {"caterpillar" : 123456, "monica" : 54321}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

實際上，你也可以寫 `passwords := map[string]int {}`，來建立一個沒有任何鍵值對的 `map`，這相當於寫 `passwords := make(map[string]int)`，不過，若是 `var passwords map[string]int` 的話，只是建立一個參考名稱 `passwords`，預設零值是 `nil`，也就是相當於 `var passwords map[string]int = nil` 的意思。

也就是說，`var passwords map[string]int` 宣告了一個參考型態，兩個 `map` 型態的名稱，可以指向同一個 `map` 實例，透過其中一個名稱來改變 `map` 內容，從另一個名稱就可以獲得對應的修改：

``` prettyprint
package main

import "fmt"

func main() {
    passwds1 := map[string]int{"caterpillar": 123456}
    passwds2 := passwds1

    fmt.Println(passwds1) // map[caterpillar:123456]

    passwds2["monica"] = 54321
    fmt.Println(passwds1) // map[monica:54321 caterpillar:123456]
}
```

# 鍵值存取、刪除

如方才所看到的，要設定一個鍵值對應時，使用 `[]` 與 `=` 指定，要取得鍵對應的值時，使用 `[]` 指定鍵，這會傳回對應的值，如果指定的鍵不存在，那麼會傳回值型態對應的零值，例如，若 `passwords := map[string]int {"caterpillar" : 123456}`，那麼 `passwords["monica"]` 會傳回 0。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

不過，更精確地說，使用 `mapName[key]` 時，會傳回兩個值（Go 中允許同時傳回多值），第一個是鍵對應的值，若沒有該鍵就傳回值型態的零值，第二個是布林值，指出鍵是否存在。例如：

``` prettyprint
package main

import "fmt"

func main() {
    passwds := map[string]int{"caterpillar": 123456}

    v, exists := passwds["monica"]
    fmt.Printf("%d %t\n", v, exists) // 0 false

    passwds["monica"] = 54321
    v, exists = passwds["monica"]
    fmt.Printf("%d %t\n", v, exists) // 54321 true
}
```

因此，若只是單純想測試鍵是否存在，只要用底線 `_` 忽略傳回的值就可以了，例如：

``` prettyprint
package main

import "fmt"

func main() {
    passwds := map[string]int{"caterpillar": 123456}
    name := "caterpillar"
    _, exists := passwds[name]
    if exists {
        fmt.Printf("%s's password is %d\n", name, passwds[name])
    } else {
        fmt.Printf("No password for %s\n", name)
    }
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

`exists` 的指定與 `if` 的判斷也可以寫在同一行：

``` prettyprint
if _, exists := passwds[name]; exists {
    fmt.Printf("%s's password is %d\n", name, passwds[name])
} else {
    fmt.Printf("No password for %s\n", name)
}
```

如果想刪除某個鍵值，可以使用 `delete` 函式，例如 `delete(passwds, "caterpillar")`。

# 迭代鍵值

如果要迭代 `map` 的鍵值，可以使用 `for range`，例如：

``` prettyprint
package main

import "fmt"

func main() {
    passwords := map[string]int{
        "caterpillar": 123456,
        "monica":      54321,
    }

    for name, passwd := range passwords {
        fmt.Printf("%s : %d\n", name, passwd)
    }
}
```

如果只是想迭代 `map` 的鍵，可以如下：

``` prettyprint
package main

import "fmt"

func main() {
    passwords := map[string]int{
        "caterpillar": 123456,
        "monica":      54321,
    }

    for name := range passwords {
        fmt.Printf("%s\n", name)
    }
}
```

如果只想迭代 `map` 的值，可以如下：

``` prettyprint
package main

import "fmt"

func main() {
    passwords := map[string]int{
        "caterpillar": 123456,
        "monica":      54321,
    }

    for _, passwd := range passwords {
        fmt.Printf("%d\n", passwd)
    }
}
```

如果想取得 `map` 中的鍵清單或者是值清單，方式之一是使用 slice 進行收集，例如：

``` prettyprint
package main

import "fmt"

func keys(m map[string]int) []string {
    ks := make([]string, 0, len(m))
    for k := range m {
        ks = append(ks, k)
    }
    return ks
}

func values(m map[string]int) []int {
    vs := make([]int, 0, len(m))
    for _, v := range m {
        vs = append(vs, v)
    }
    return vs
}

func main() {
    passwords := map[string]int{
        "caterpillar": 123456,
        "monica":      54321,
    }

    fmt.Println(keys(passwords))   // [caterpillar monica]
    fmt.Println(values(passwords)) // [123456 54321]
}
```

Go 的 `map` 在迭代時沒有一定的順序，如果想要有排序結果，必須自行處理，例如，針對鍵排序來進行迭代：

``` prettyprint
package main

import "sort"
import "fmt"

func main() {
    passwords := map[string]int{
        "caterpillar": 123456,
        "monica":      54321,
        "hamimi":      13579,
    }

    keys := make([]string, 0, len(passwords))
    for key := range passwords {
        keys = append(keys, key)
    }
    sort.Strings(keys)

    for _, key := range keys {
        fmt.Printf("%s : %d\n", key, passwords[key])
    }
}
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
