<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 結構與方法

</div>

  

在〈[結構入門](Struct.html)〉中看過，有些資料會有相關性，相關聯的資料組織在一起，對於資料本身的可用性或者是程式碼的可讀性，都會有所幫助，實際上，有些資料與可處理它的函式也會有相關性，將相關聯的資料與函式組織在一起，對資料與函式本身的可用性或者是程式碼的可讀性，也有著極大的幫助。

# 建立方法

假設可能原本有如下的程式內容，負責銀行帳戶的建立、存款與提款：

``` prettyprint
package main

import (
    "errors"
    "fmt"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func Deposit(account *Account, amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    account.balance += amount
}

func Withdraw(account *Account, amount float64) error {
    if amount > account.balance {
        return errors.New("餘額不足")
    }
    account.balance -= amount
    return nil
}

func String(account *Account) string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        account.id, account.name, account.balance)
}

func main() {
    account := &Account{"1234-5678", "Justin Lin", 1000}
    Deposit(account, 500)
    Withdraw(account, 200)
    fmt.Println(String(account)) // Account{1234-5678 Justin Lin 1300.00}
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

實際上，`Desposit`、`Withdraw`、`String` 的函式操作，都是與傳入的 `Account` 實例有關，何不將它們組織在一起呢？這樣比較容易使用些，在 Go 語言中，你可以重新修改函式如下：

``` prettyprint
package main

import (
    "errors"
    "fmt"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) Deposit(amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    ac.balance += amount
}

func (ac *Account) Withdraw(amount float64) error {
    if amount > ac.balance {
        return errors.New("餘額不足")
    }
    ac.balance -= amount
    return nil
}

func (ac *Account) String() string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        ac.id, ac.name, ac.balance)
}

func main() {
    account := &Account{"1234-5678", "Justin Lin", 1000}
    account.Deposit(500)
    account.Withdraw(200)
    fmt.Println(account.String()) // Account{1234-5678 Justin Lin 1300.00}
}
```

簡單來說，只是將函式的第一個參數，移至方法名稱之前成為函式呼叫的接收者（Receiver），這麼一來，就可以使用 `account.Deposit(500)`、`account.Withdraw(200)`、`account.String()` 這樣的方式來呼叫函式，就像是物件導向程式語言中的方法（Method）。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

注意到，在這邊使用的是 `(ac *Account)`，也就是指標，如果你是如下使用 `(ac Account)`：

``` prettyprint
func (ac Account) Deposit(amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    ac.balance += amount
}
```

那麼執行像是 `account.Deposit(500)`，就像是以 `Deposit(*account, 500)` 呼叫以下函式：

``` prettyprint
func Deposit(account Account, amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    account.balance += amount
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

也就是，相當於將 `Account` 實例以傳值方式複製給 `Deposit` 函式的參數。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

某些程度上，可以將接收者想成是其他語言中的 `this` 或 `self`，[Go 建議為接收者適當命名](https://github.com/golang/go/wiki/CodeReviewComments#receiver-names)，而不是用 `this`、`self` 之類的名稱。接收者並沒有文件上記載的作用，命名時不用其他參數具有一定的描述性，只要能表達程式意圖就可以了，Go 建議是個一或兩個字母的名稱（某些程度上，也可以用來與其他參數區別）。

# 名稱相同的方法

之前談過，Go 語言中不允許方法重載（Overload），因此，對於以下的程式，是會發生 `String` 重複宣告的編譯錯誤：

``` prettyprint
package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

func String(account *Account) string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        account.id, account.name, account.balance)
}

type Point struct {
    x, y int
}

func String(point *Point) string { // String redeclared in this block 的編譯錯誤
    return fmt.Sprintf("Point{%d %d}", point.x, point.y)
}

func main() {
    account := &Account{"1234-5678", "Justin Lin", 1000}
    point := &Point{10, 20}
    fmt.Println(account.String())
    fmt.Println(point.String())
}
```

然而，若是將函式定義為方法，就不會有這個問題，Go 可以從方法的接收者辨別，該使用哪個 `String` 方法：

``` prettyprint
package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) String() string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        ac.id, ac.name, ac.balance)
}

type Point struct {
    x, y int
}

func (p *Point) String() string {
    return fmt.Sprintf("Point{%d %d}", p.x, p.y)
}

func main() {
    account := &Account{"1234-5678", "Justin Lin", 1000}
    point := &Point{10, 20}
    fmt.Println(account.String()) // Account{1234-5678 Justin Lin 1000.00}
    fmt.Println(point.String())   // Point{10 20}
}
```

# 方法作為值

在 Go 語言中，函式也可以作為值傳遞，那麼就產生了一個問題，方法呢？既然方法本質上也是個函式，那麼是否也可以作為值傳遞，答案是可以的，不過，以上面的程式為例，你不能直接以 `String := String` 這樣的方式傳遞，而必須使用方法運算式（Method expression）。例如：

``` prettyprint
package main

import (
    "errors"
    "fmt"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) Deposit(amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    ac.balance += amount
}

func (ac *Account) Withdraw(amount float64) error {
    if amount > ac.balance {
        return errors.New("餘額不足")
    }
    ac.balance -= amount
    return nil
}

func (ac *Account) String() string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        ac.id, ac.name, ac.balance)
}

func main() {
    deposit := (*Account).Deposit
    withdraw := (*Account).Withdraw
    String := (*Account).String

    account1 := &Account{"1234-5678", "Justin Lin", 1000}
    deposit(account1, 500)
    withdraw(account1, 200)
    fmt.Println(String(account1)) // Account{1234-5678 Justin Lin 1300.00}

    account2 := &Account{"5678-1234", "Monica Huang", 500}
    deposit(account2, 250)
    withdraw(account2, 150)
    fmt.Println(String(account2)) // Account{5678-1234 Monica Huang 600.00}
}
```

可以看到，這樣取得的函式，就像是本文一開始的範例那樣，你可以傳入任何的 `Account` 實例。另一個取得方法的方式是方法值（Method value），這會保有取得方法當時的接收者：

``` prettyprint
package main

import (
    "errors"
    "fmt"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) Deposit(amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    ac.balance += amount
}

func (ac *Account) Withdraw(amount float64) error {
    if amount > ac.balance {
        return errors.New("餘額不足")
    }
    ac.balance -= amount
    return nil
}

func (ac *Account) String() string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        ac.id, ac.name, ac.balance)
}

func main() {
    account1 := &Account{"1234-5678", "Justin Lin", 1000}
    acct1Deposit := account1.Deposit
    acct1Withdraw := account1.Withdraw
    acct1String := account1.String

    acct1Deposit(500)
    acct1Withdraw(200)
    fmt.Println(acct1String()) // Account{1234-5678 Justin Lin 1300.00}

    account2 := &Account{"5678-1234", "Monica Huang", 500}
    acct2Deposit := account2.Deposit
    acct2Withdraw := account2.Withdraw
    acct2String := account2.String

    acct2Deposit(250)
    acct2Withdraw(150)
    fmt.Println(acct2String()) // Account{5678-1234 Monica Huang 600.00}
}
```

# 值都能有方法

實際上，不只是結構的實例可以定義方法，在 Go 語言中，只要是值，就可以定義方法，條件是必須是定義的型態（defined type），具體而言，就是使用 `type` 定義的新型態。

例如，以下的範例為 `[]int` 定義了一個新的型態名稱，並定義了一個 `ForEach` 方法：

``` prettyprint
package main

import "fmt"

type IntList []int
type Funcint func(int)

func (lt IntList) ForEach(f Funcint) {
    for _, ele := range lt {
        f(ele)
    }
}

func main() {
    var lt IntList = []int{10, 20, 30, 40, 50}
    lt.ForEach(func(ele int) {
        fmt.Println(ele)
    })
}
```

這個範例會顯示 10 到 50 作為結果，必須留意的是，`type` 定義了新型態 `Funcint`，因為 `ForEach` 是針對 `Funcint` 定義，而不是針對 `[]int`，因此底下是行不通的：

``` prettyprint
lt2 := []int {10, 20, 30, 40, 50}

// lt2.ForEach undefined (type []int has no field or method ForEach)
lt2.ForEach(func(ele int) {
    fmt.Println(ele)
})
```

編譯器認為 `[]int` 並沒有定義 `ForEach`，因此發生錯誤，想要通過編譯的話，可以進行型態轉換：

``` prettyprint
lt2 := IntList([]int {10, 20, 30, 40, 50})
lt2.ForEach(func(ele int) {
    fmt.Println(ele)
})
```

你甚至可以基於 `int` 等基本型態定義方法，同樣地，必須定義一個新的型態名稱：

``` prettyprint
package main

import (
    "fmt"
)

type Int int
type FuncInt func(Int)

func (n Int) Times(f FuncInt) {
    if n < 0 {
        panic("必須是正數")
    }

    var i Int
    for i = 0; i < n; i++ {
        f(i)
    }
}

func main() {
    var x Int = 10
    x.Times(func(n Int) {
        fmt.Println(n)
    })
}
```

像這樣基於某個基本型態定義新型態，並為其定義更多高階特性，在 Go 的領域是常見的做法。這個範例會顯示 0 到 9，看起來就像是指定函式，要求執行 x 次吧！…XD

# nil 接收者

在 Go 中，接收者可以是 `nil`，這讓你有機會在方法中處理接收者為 `nil` 的情況，例如：

``` prettyprint
package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) String() string {
    if ac == nil {
        return "<nil>"
    }
    return fmt.Sprintf("Account{%s %s %.2f}",
        ac.id, ac.name, ac.balance)
}

func findById(id string) *Account {
    accts := []*Account{&Account{"123", "Justin Lin", 10000}, &Account{"456", "Monica", 10000}}
    for i := 0; i < len(accts); i++ {
        if accts[i].id == id {
            return accts[i]
        }
    }
    return nil
}

func main() {
    fmt.Println(findById("123").String())
    fmt.Println(findById("789").String())
}
```

如果是其他語言，例如 Java 的話，在 `findById("789").String()` 的地方會 `NullPointerException`，不過在 Go 中，可以針對接收者是否為 `nil`，來決定如何處理，例如這邊就實作了 nil safety 的概念。

# 模擬建構式、初始式

Go 沒有物件導向語言中建構式或初始式之類的概念，然而可以自行模擬，例如在 [container/list](https://pkg.go.dev/container/list/) 的[原始碼](https://go.dev/src/container/list/list.go)可以看到 `New` 作為一個工廠函式，用來建立新的 `List`，初始的流程寫在 `Init` 方法之中：

``` prettyprint
...
// Init initializes or clears list l.
func (l *List) Init() *List {
    l.root.next = &l.root
    l.root.prev = &l.root
    l.len = 0
    return l
}

// New returns an initialized list.
func New() *List { return new(List).Init() }
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
