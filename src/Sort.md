<div id="main" role="main" style="height: auto !important;">

<div class="header">

# sort 套件

</div>

  

Go 提供了 `sort` 套件來協助排序、搜尋任務，對於 `[]int`、`[]float64` 與 `[]string`，可以透過 `Ints`、`Float64s`、`Strings` 來由小而大排序，可以使用 `IntsAreSorted`、`Float64sAreSorted`、`StringsAreSorted` 來看看是否已經排序。

若想在已由小而大排序的 `[]int`、`[]float64` 與 `[]string` 中進行搜尋，可以使用 `SearchInts`、`SearchFloat64s`、`SearchStrings` 函式，搜尋結果將傳回找到搜尋值的索引位置，**沒有搜尋到的話，傳回的會是可以安插搜尋值的索引位置**。例如：

``` prettyprint
package main

import (
    "fmt"
    "sort"
)

func main() {
    s := []int{8, 2, 6, 3, 1, 4} 
    sort.Ints(s)
    fmt.Println(sort.IntsAreSorted(s)) // true
    fmt.Println(s)                     // [1 2 3 4 6 8]
    fmt.Println(sort.SearchInts(s, 7)) // 5
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

如果想要由大而小排序呢？可以透過 `Slice`、`SliceStable`，指定一個 `less` 函式，該函式接受兩個索引，你要傳回布林值表示 `i` 處的值順序上是否小於 `j`：

``` prettyprint
package main

import (
    "fmt"
    "sort"
)

func main() {
    s := []int{8, 2, 6, 3, 1, 4} 
    sort.Slice(s, func(i, j int) bool {
        return s[i] > s[j]
    })
    fmt.Println(s)  // [8 6 4 3 2 1]
}
```

實際上，`Slice`、`SliceStable` 可用於任意的結構，例如：

``` prettyprint
package main

import (
    "fmt"
    "sort"
)

func main() {
    family := []struct {
        Name string
        Age  int
    } {{"Irene", 12}, {"Justin", 45}, {"Monica", 42}}

    // 依年齡由小而大排序
    sort.SliceStable(family, func(i, j int) bool {
        return family[i].Age < family[j].Age
    })

    fmt.Println(family) // [{Irene 12} {Monica 42} {Justin 45}]
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

那麼怎麼搜尋上面的 `family` 呢？例如，找出年齡 45 歲的資料？這可以用 `Search`，例如：

``` prettyprint
package main

import (
    "fmt"
    "sort"
)

func main() {
    family := []struct {
        Name string
        Age  int
    } {{"Irene", 12}, {"Justin", 45}, {"Monica", 42}}

    // 依年齡由小而大排序
    sort.SliceStable(family, func(i, j int) bool {
        return family[i].Age < family[j].Age
    })

    fmt.Println(family) // [{Irene 12} {Monica 42} {Justin 45}]

    idx := sort.Search(len(family), func (i int) bool {
        return family[i].Age == 45
    })
    fmt.Println(idx)
}
```

`Search` 會使用二分搜尋，第二個參數指定的函式要傳回布林值，表示是否符合搜尋條件，若找到第一個符合的話傳回索引位置，否則傳回第一個參數指定的值。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

在 [Search](https://pkg.go.dev/sort/#Search) 說明中，還有個猜數字的有趣範例，由程式猜出你心中想的數字：

``` prettyprint
func GuessingGame() {
    var s string
    fmt.Printf("Pick an integer from 0 to 100.\n")
    answer := sort.Search(100, func(i int) bool {
        fmt.Printf("Is your number <= %d? ", i)
        fmt.Scanf("%s", &s)
        return s != "" && s[0] == 'y'
    })
    fmt.Printf("Your number is %d.\n", answer)
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

`sort` 還提供了 `Sort`、`Stable` 函式，乍看很奇怪：

``` prettyprint
func Sort(data Interface)
func Stable(data Interface)
```

`Interface` 的定義是：

``` prettyprint
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

這是給有序、具索引的資料結構實現的行為，任何具有 `Interface` 行為的資料結構，都可以透過 `Sort`、`Stable` 函式排序，`sort` 套件提供的實作有 `IntSlice`、`Float64Slice`、`StringSlice`，以 `IntSlice` 的原始碼實現為例：

``` prettyprint
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
```

因此，若要對整數排序，也可以如下：

``` prettyprint
package main

import (
    "fmt"
    "sort"
)

func main() {
    s := sort.IntSlice([]int{8, 2, 6, 3, 1, 4})
    sort.Sort(s)
    fmt.Println(s)                     // [1 2 3 4 6 8]
}
```

實際上，`Ints`、`Float64s`、`Strings` 函式，內部也只是轉換為 `IntSlice`、`Float64Slice`、`StringSlice`，然後呼叫 `Sort` 罷了：

``` prettyprint
func Ints(a []int) { Sort(IntSlice(a)) }
func Float64s(a []float64) { Sort(Float64Slice(a)) }
func Strings(a []string) { Sort(StringSlice(a)) }
```

對於一個實現了 `Interface` 的資料結構，除了可以使用 `Sort`、`Stable` 函式外，若需要反向排序，可以有個簡單方式，透過 `Reverse` 來包裹。例如：

``` prettyprint
package main

import (
    "fmt"
    "sort"
)

func main() {
    s := sort.IntSlice([]int{8, 2, 6, 3, 1, 4})
    sort.Sort(sort.Reverse(s))
    fmt.Println(s)                     // [8 6 4 3 2 1]
}
```

有趣的是 `Reverse` 的實作，它不過就是將給原本資料結構 `Less` 方法的 `i`、`j` 對調罷了：

``` prettyprint
type reverse struct {
    Interface
}

func (r reverse) Less(i, j int) bool {
    return r.Interface.Less(j, i)
}

func Reverse(data Interface) Interface {
    return &reverse{data}
} 
```

來自己實現一下 `Interface`，使用家人的年齡來排序：

``` prettyprint
package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

type Family []Person

func (f Family) Len() int {
    return len(f)
}

func (f Family) Less(i, j int) bool {
    return f[i].Age < f[j].Age
}

func (f Family) Swap(i, j int) {
    f[i], f[j] = f[j], f[i]
}

func main() {
    family := Family{{"Irene", 12}, {"Justin", 45}, {"Monica", 42}}

    sort.Sort(family)
    fmt.Println(family)  // [{Irene 12} {Monica 42} {Justin 45}]

    sort.Sort(sort.Reverse(family))
    fmt.Println(family)  // [{Justin 45} {Monica 42} {Irene 12}]
}
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
