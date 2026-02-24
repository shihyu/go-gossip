<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 匿名函式與閉包

</div>

  

除了作為值傳遞之外，Go 的函式還可以是匿名函式，且具有閉包（Closure）的特性，由於 Go 具有指標，在理解閉包時反而容易一些了。

# 匿名函式

在〈[一級函式](http://openhome.cc/Gossip/Go/FirstClassFunction.html)〉中，我們看過函式可作為值傳遞的一個應用是，可將函式傳入另一函式作為回呼（Callback），除了傳遞具名的函式之外，有時會想要臨時建立一個函式進行傳遞，例如：

``` prettyprint
package main

import "fmt"

type Predicate = func(int) bool

func filter(origin []int, predicate Predicate) []int {
    filtered := []int{}
    for _, elem := range origin {
        if predicate(elem) {
            filtered = append(filtered, elem)
        }
    }
    return filtered
}

func main() {
    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println(filter(data, func(elem int) bool {
        return elem > 5
    }))
    fmt.Println(filter(data, func(elem int) bool {
        return elem <= 6
    }))
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

這個函式與〈[一級函式](http://openhome.cc/Gossip/Go/FirstClassFunction.html)〉中最後一個範例的作用相同，不過這次傳遞了匿名函式給 `filter`，可以看到，匿名函式可使用 `func` 建立，同樣必須指定參數與傳回值型態。

在 Go 中，不允許在函式中又宣告函式，例如，以下是不允許的：

``` prettyprint
func funcA() {
    func funcB() {
        ...
    }
    ...
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

這會出現 “nested func not allowed” 的錯誤，然而，你可以建立匿名函式，然後將之指定給某個變數：

``` prettyprint
func funcA() {
    funcB := func() {
       ...
    }
    ...
}
```

你也可以在函式中建立匿名函式，並將之傳回：

``` prettyprint
package main

import "fmt"

type Func1 = func(int) int

func funcA() Func1 {
    x := 10
    return func(n int) int {
        return x + n
    }
}

func main() {
    fmt.Println(funcA()(2)) // 12
}
```

在上面的範例中，執行 `funcA` 會傳回一個函式，這個傳回的函式會將接受的引數指定給參數 `n`，並與 `x` 的值進行相加，因此最後顯示結果為 12。

# 閉包

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

可以在函式中建立匿名函式，引發了一個有趣的事實，先來看個例子：

``` prettyprint
package main

import "fmt"

type Consumer = func(int)

func forEach(elems []int, consumer Consumer) {
    for _, elem := range elems {
        consumer(elem)
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    sum := 0
    forEach(numbers, func(elem int) {
        sum += elem
    })
    fmt.Println(sum) // 15
}
```

乍看之下，似乎有點像是：

``` prettyprint
package main

import "fmt"

type Consumer = func(int)

func forEach(elems []int, consumer Consumer) {
    for _, elem := range elems {
        consumer(elem)
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    sum := 0
    for _, elem := range numbers {
        sum += elem
    }
    fmt.Println(sum) // 15
}
```

然而意義完全不同。在使用 `forEach` 函式的範例中，`sum` 變數被匿名函式包覆並傳入 `forEach` 之中，在 `forEach` 執行迴圈的過程中，每次呼叫傳入的函式（被 `consumer` 參考），就會改變 `sum` 的值，因此，最後得到的是加總後的值 15。

實際上，使用 `forEach` 函式的範例中，建立了一個閉包，閉包本質上就是一個匿名函式，`sum` 變數被閉包包覆，讓 `sum` 變數可以存活於閉包的範疇中，其實，更之前從 `funcA` 傳回函式的範例中，也建立了閉包，`funcA` 的 `x` 區域變數被閉包包覆，因此，你執行傳回的函式時，即使 `funcA` 已執行完畢，`x` 變數依然是存活著在傳回的閉包範疇中，所以，你指定的引數總是會與 `x` 的值進行相加。

重點在於，閉包將**變數本身**關閉在自己的範疇中，而不是變數的值，可以用以下這個範例來做個示範：

``` prettyprint
package main

import "fmt"

type Getter = func() int
type Setter = func(int)

func x_getter_setter(x int) (Getter, Setter) {
    getter := func() int {
        return x
    }
    setter := func(n int) {
        x = n
    }
    return getter, setter
}

func main() {
    getX, setX := x_getter_setter(10)

    fmt.Println(getX()) // 10
    setX(20)
    fmt.Println(getX()) // 20
}
```

對 `x_getter_setter` 來說，`x` 參數也是變數，`x_getter_setter` 傳回了兩個匿名函式，這兩個匿名函式都形成了閉包，將 `x` 變數關閉在自己的範疇中，因此，你使用了 `setX(20)` 改變了 `x` 的值，使用 `getX()` 時取得的值，就會是修改後的值。

# 閉包與指標

如果你寫過 JavaScript，對於方才的範例，應該不會陌生，也因為 JavaScript 的普及，現在開發者多半對閉包不會覺得神秘難解了，而對於「閉包將變數本身關閉在自己的範疇中，而不是變數的值」，也比較瞭解其應用所在。

由於 Go 語言有指標，我們可以將指標的值顯示出來，這代表著變數的位址值，來看看被閉包關閉的變數，到底是怎麼一回事好了：

``` prettyprint
package main

import "fmt"

type Getter = func() int
type Setter = func(int)

func x_getter_setter(x int) (Getter, Setter) {
    fmt.Printf("the parameter :\tx (%p) = %d\n", &x, x)

    getter := func() int {
        fmt.Printf("getter invoked:\tx (%p) = %d\n", &x, x)
        return x
    }
    setter := func(n int) {
        x = n
        fmt.Printf("setter invoked:\tx (%p) = %d\n", &x, x)
    }
    return getter, setter
}

func main() {
    getX, setX := x_getter_setter(10)

    fmt.Println(getX())
    setX(20)
    fmt.Println(getX())
}
```

這個範例與前一個範例類似，只不過呼叫函式時，都會顯示 `x` 變數的位址值與儲存值，一個執行結果是：

``` prettyprint
the parameter : x (0x104382e0) = 10
getter invoked: x (0x104382e0) = 10
10
setter invoked: x (0x104382e0) = 20
getter invoked: x (0x104382e0) = 20
20
```

看到了嗎？顯示的變數的位址值都是相同的，閉包將**變數本身**關閉在自己的範疇中，而不是變數的值，就是這麼一回事。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
