<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 泛型入門（Go 1.18+）

</div>

  

從 Go 1.18 開始，Go 支援泛型（Generics，也就是 type parameters）。這讓你可以用同一份程式碼，處理多種型態，而不用到處寫重複函式或依賴 `interface{}` 加型態斷言。

# 泛型函式

函式可以在名稱後面宣告型別參數，例如：

``` go
package main

import "fmt"

func First[T any](elems []T) (T, bool) {
    if len(elems) == 0 {
        var zero T
        return zero, false
    }
    return elems[0], true
}

func main() {
    if v, ok := First([]int{10, 20, 30}); ok {
        fmt.Println(v) // 10
    }
    if v, ok := First([]string{"Go", "Rust"}); ok {
        fmt.Println(v) // Go
    }
}
```

`T any` 表示 `T` 可以是任何型態；`any` 是 `interface{}` 的別名（Go 1.18 新增）。

通常編譯器可以從引數推斷型別參數，因此呼叫時常不必寫成 `First[int](...)`。

# 型別條件（constraints）

如果泛型函式或泛型型別需要使用某些運算（例如 `==`、`<`），就要限制型別參數可接受的型態範圍，這就是型別條件。

``` go
package main

import "fmt"

func IndexOf[T comparable](elems []T, target T) int {
    for i, elem := range elems {
        if elem == target {
            return i
        }
    }
    return -1
}

func main() {
    fmt.Println(IndexOf([]int{1, 2, 3}, 2))           // 1
    fmt.Println(IndexOf([]string{"Go", "C"}, "Rust")) // -1
}
```

`comparable` 也是 Go 1.18 新增的預定義識別名稱，只能用在型別條件中。從 Go 1.20 起，可比較（但可能在執行時比較時 panic）的型別，也能滿足 `comparable` 條件。

# 介面作為型別集合（type set）

Go 1.18 之後，介面不只描述方法集合，也可以描述型別集合（用於 constraint）。例如：

``` go
type Integer interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}
```

`~int` 表示「底層型態是 `int` 的型態」，這讓自訂型別（例如 `type MyInt int`）也能符合條件。

這類帶有 union（`|`）或具體型別元素的介面，通常只能用在型別條件，不會拿來當一般執行期介面值使用。

# 泛型型別

不只函式可以用泛型，型別宣告也可以：

``` go
package main

import "fmt"

type Stack[T any] struct {
    elems []T
}

func (s *Stack[T]) Push(v T) {
    s.elems = append(s.elems, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elems) == 0 {
        var zero T
        return zero, false
    }
    i := len(s.elems) - 1
    v := s.elems[i]
    s.elems = s.elems[:i]
    return v, true
}

func main() {
    var s Stack[string]
    s.Push("Go")
    s.Push("1.26")
    v, _ := s.Pop()
    fmt.Println(v) // 1.26
}
```

# Go 1.24：generic type aliases

Go 1.24 起，generic type aliases 完整支援，因此可像下面這樣寫：

``` go
type Set[T comparable] = map[T]struct{}
```

這在整理常用資料結構別名時很方便，例如 `Set[string]`、`Set[int]`。

# Go 1.26：型別參數列表中的自我參照

Go 1.26 放寬了限制，泛型型別可在自己的型別參數列表中參照自己。這對某些泛型介面（特別是要求「輸入輸出都是自身型態」）會比較自然：

``` go
package main

import "fmt"

type Adder[A Adder[A]] interface {
    Add(A) A
}

type MyInt int

func (x MyInt) Add(y MyInt) MyInt {
    return x + y
}

func Algo[A Adder[A]](x, y A) A {
    return x.Add(y)
}

func main() {
    fmt.Println(Algo(MyInt(10), MyInt(20))) // 30
}
```

# 何時該用泛型？

泛型適合用在：

- 演算法或容器邏輯完全相同，只差元素型態（例如 stack、set、搜尋、走訪）
- 想避免 `interface{}` 與型態斷言造成的樣板程式碼與執行期錯誤
- 需要在編譯期保留型態資訊，提高可讀性與型別安全

但如果某個邏輯本質上就是依賴行為（方法）而不是型態，傳統介面仍然常常是更直接的做法。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
