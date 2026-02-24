<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 介面入門

</div>

  

在〈[結構組合](StructComposition.html)〉的最後討論到了多型，倘若現在需要有個函式，可以接受 `Account` 與 `CheckingAccount` 實例，或者是有個陣列或 slice，可以收集 `Account` 與 `CheckingAccount`實例，那該怎麼辦呢？

# 介面定義行為

在 Go 語言中，可以使用 `interface` 定義行為，舉例來說，若現在想要定義儲蓄的行為，可以如下：

``` prettyprint
type Savings interface {
    Deposit(amount float64)
    Withdraw(amount float64) error
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

注意，不必使用 `func` 關鍵字，也不用宣告接受者型態，只需要定義行為的名稱、參數與傳回值。接著該怎麼實現這個介面呢？實際上，就〈[結構組合](StructComposition.html)〉，已經實現了這個介面，也就是說，結構上不用任何關鍵字，只要有函式實現這兩個行為就可以了。

因此，現在可以寫個函式，同時接受 `Account` 與 `CheckingAccount` 實例，在提款後顯示餘額：

``` prettyprint
package main

import (
    "errors"
    "fmt"
)

type Savings interface {
    Deposit(amount float64)
    Withdraw(amount float64) error
}

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

type CheckingAccount struct {
    Account
    overdraftlimit float64
}

func (ac *CheckingAccount) Withdraw(amount float64) error {
    if amount > ac.balance+ac.overdraftlimit {
        return errors.New("超出信用額度")
    }
    ac.balance -= amount
    return nil
}

func Withdraw(savings Savings) {
    if err := savings.Withdraw(500); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(savings)
    }
}

func main() {
    account1 := Account{"1234-5678", "Justin Lin", 1000}
    account2 := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}
    Withdraw(&account1) // 顯示 &{1234-5678 Justin Lin 500}
    Withdraw(&account2) // 顯示 &{{1234-5678 Justin Lin 500} 30000}
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

雖然沒有定義接收者為 `*CheckingAccount` 的 `Deposit` 方法，然而，作為內部型態的 `Account` 有定義 `Deposit`（並且沒有使用到 `CheckingAccount` 定義的值域），這個實現被提昇至外部型態，也就滿足了 `Savings` 要求的行為規範。

注意！由於在實作 `Withdraw` 與 `Deposit` 方法時，都是用指標 `(ac *Account)` 或 `(ac *CheckingAccount)` 宣告了接受者型態，因此傳遞實例給 `func Withdraw(savings Savings)` 時，也就必須傳遞指標。

如果在實作`Withdraw` 與 `Deposit` 方法時，是使用 `(ac Account)` 或 `(ac CheckingAccount)` 宣告了接受者型態，那麼傳遞實例給接受 `Savings` 的函式時，就可以不用取指標，例如：

``` prettyprint
package main

import (
    "errors"
    "fmt"
)

type Savings interface {
    Deposit(amount float64)
    Withdraw(amount float64) error
}

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac Account) Deposit(amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    ac.balance += amount
}

func (ac Account) Withdraw(amount float64) error {
    if amount > ac.balance {
        return errors.New("餘額不足")
    }
    ac.balance -= amount
    return nil
}

type CheckingAccount struct {
    Account
    overdraftlimit float64
}

func (ac CheckingAccount) Withdraw(amount float64) error {
    if amount > ac.balance+ac.overdraftlimit {
        return errors.New("超出信用額度")
    }
    ac.balance -= amount
    return nil
}

func Withdraw(savings Savings) {
    if err := savings.Withdraw(500); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(savings)
    }
}

func main() {
    account1 := Account{"1234-5678", "Justin Lin", 1000}
    account2 := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}
    Withdraw(account1) // 顯示 {1234-5678 Justin Lin 1000}
    Withdraw(account2) // 顯示 {{1234-5678 Justin Lin 1000} 30000}
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

當然，就這個例子來說，結果並不是正確的，就算改成 `Withdraw(&account1)` 與 `&Withdraw(account2)`，也不會是正確的結果，因為就 `Withdraw` 與 `Deposit` 的接收者來說，會是複製結構的值域，而不是修改原結構實例的值域，這純綷只是示範。

# 介面實例的型態與值

如果你定義了一個變數：

``` prettyprint
var savings Savings
```

那麼 `savings` 變數儲存了什麼？技術上來說，`savings` 變數儲存兩個資訊：型態與值。就方才的`savings` 被指定為 `nil` 來說，代表著 `savings` 在底層儲存的型態為 `nil`，而值沒有指定，這樣的介面實例稱為 nil interface，因為沒有型態資訊，也就不能透過 nil interface 呼叫方法。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

如果接收者是定義為 `(ac *Account)`，而且有底下的程式，那麼 `savings` 底層儲存的型態會 `*Account`，而值是 `Account` 結構實例的位址值：

``` prettyprint
var savings Savings = &Account{"1234-5678", "Justin Lin", 1000}
```

當接收者是指標時，透過介面比對是否為 `nil` 時要留意，例如以下會是 `true`，這是因為 `savings` 在底層儲存的型態為 `nil`，而值沒有指定，介面宣告的變數只有在這個情況下，跟 `nil` 直接相等比較才會是 `true`：

``` prettyprint
var savings Savings = nil
fmt.Println(savings == nil)
```

然而以下會是 `false`，這是因為 `savings` 在底層儲存的型態為 `*Account`，而值是 `nil`（  
這時透過 `savings` 是可以呼叫方法的，接收者會是 `nil`，就看你要不要在方法中處理 `nil` 了）：

``` prettyprint
var acct *Account = nil
var savings Savings = acct
fmt.Println(savings == nil)
```

這是個 FAQ 了，在〈[Why is my nil error value not equal to nil?](https://go.dev/doc/faq#nil_error)〉就提到了個例子：

``` prettyprint
func returnsError() error {
    var p *MyError = nil
    if bad() {
        p = ErrBad
    }
    return p
}
```

如果對 `returnsError` 傳回值進行 `nil` 比較，結果會是 `false`：

``` prettyprint
fmt.Println(returnsError() == nil) // false
```

因此如果傳回型態是個介面，值會是 `nil`，請記得直接傳 `nil`：

``` prettyprint
func returnsError() error {
    if bad() {
        return ErrBad
    }
    return nil // 直接傳 nil
}
```

如果接收者是定義為 `(ac Account)`，而你有底下的程式：

``` prettyprint
var savings Savings = Account{"1234-5678", "Justin Lin", 1000}
```

這時 `savings` 在底層會儲存型態 `Account`，而值為結構實例，這時透過 `Savings` 來進行實例的指定時，底層也會是結構實例的指定，因此會發生複製：

``` prettyprint
var savings1 Savings = Account{"1234-5678", "Justin Lin", 1000}
var savings2 Savings = savings1

savings2.name = "Monica Huang"
fmt.Println(savings.name) // Justin Lin
```

# Go 1.18+：介面也可作為型別條件

從 Go 1.18 開始，`interface` 除了用來描述物件要有哪些方法，也可以用來描述「型別集合（type set）」，作為泛型的型別條件（constraint）。

例如：

``` prettyprint
type StringKeyed interface {
    ~string
}

func HasKey[K StringKeyed, V any](m map[K]V, key K) bool {
    _, ok := m[key]
    return ok
}
```

上例中的 `StringKeyed` 並不是拿來做一般執行期介面值（例如 `var x StringKeyed`），而是拿來限制型別參數 `K` 的可用型別。

另外，Go 1.18 也新增了兩個常見的預定義識別名稱：

- `any`：`interface{}` 的別名。
- `comparable`：可用 `==`、`!=` 比較的型別集合（只能用在型別條件）。

Go 1.20 之後，像一般介面型別這類「可比較但可能在執行時 panic」的型別，也可以滿足 `comparable` 條件，因此像 `Set[any]` 這類泛型實例化會更容易成立；只是若實際比較到不可比較的動態值（例如內含 slice 的介面值），仍可能在執行時發生 panic。

# 異質陣列或 slice

Go 語言會檢查類型的實例，是否實現了介面中規範的行為，若是的話，就可以使用介面型態來接受不同型態實例的指定，因此，若要建立一個異質陣列或 slice，也是可以的：

``` prettyprint
package main

import (
    "errors"
    "fmt"
)

type Savings interface {
    Deposit(amount float64)
    Withdraw(amount float64) error
}

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

type CheckingAccount struct {
    Account
    overdraftlimit float64
}

func (ac *CheckingAccount) Withdraw(amount float64) error {
    if amount > ac.balance+ac.overdraftlimit {
        return errors.New("超出信用額度")
    }
    ac.balance -= amount
    return nil
}

func main() {
    savingsArray := [...]Savings{
        &Account{"1234-5678", "Justin Lin", 1000},
        &CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000},
    }

    for _, savings := range savingsArray {
        fmt.Println(savings)
    }

    savingsSlice := []Savings{
        &Account{"1234-5678", "Justin Lin", 1000},
        &CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000},
    }

    for _, savings := range savingsSlice {
        fmt.Println(savings)
    }
}
```

在這邊雖然是以 `Account` 及 `CheckingAccount` 為例，不過，只要實現了 `Savings` 的行為，就算是一隻鴨子，也是可以的：

``` prettyprint
package main

import "fmt"

type Savings interface {
    Deposit(amount float64)
    Withdraw(amount float64) error
}

type Duck struct{}

func (d *Duck) Deposit(amount float64) {
    fmt.Println("我是一隻鴨子，我沒帳戶")
}

func (d *Duck) Withdraw(amount float64) error {
    fmt.Println("我是一隻鴨子，我沒錢")
    return nil
}

func main() {
    duckArray := [...]Savings{
        &Duck{},
        &Duck{},
    }

    for _, duck := range duckArray {
        duck.Deposit(1000)
    }

    duckSlice := []Savings{
        &Duck{},
        &Duck{},
    }

    for _, duck := range duckSlice {
        duck.Withdraw(500)
    }
}
```

# 空介面

那麼，如果想要建立一個實例容器，可以收集各種類型的實例，要怎麼做呢？答案就是透過空介面，也就是沒有定義任何行為的 `interface {}`。

``` prettyprint
package main

import "fmt"

type Duck struct{}

func main() {
    instances := [](interface{}){
        &Duck{},
        [...]int{1, 2, 3, 4, 5},
        map[string]int{"caterpillar": 123456, "monica": 54321},
    }

    for _, instance := range instances {
        fmt.Println(instance)
    }
}
```

如果你查看 `fmt.Println` 的文件說明，可以發現，它的參數類型就是 `interface {}`：

``` prettyprint
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```

順便一提的是，就目前來說，在使用 `fmt.Println` 顯示結構時，都是使用預設的字串格式，如果想自訂字串格式，必須實現 `Stringer` 這個介面，這定義在 `fmt` 的 print.go 之中：

``` prettyprint
type Stringer interface {
        String() string
}
```

在需要字串的場合中，會呼叫 `String()` 方法。例如，若你想要帳號顯示時，可以出現 Account 或 CheckingAccount 字樣的話，可以如下實作：

``` prettyprint
package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) String() string {
    return fmt.Sprintf("Account(id = %s, name = %s, balance = %.2f)",
        ac.id, ac.name, ac.balance)
}

type CheckingAccount struct {
    Account
    overdraftlimit float64
}

func (ac *CheckingAccount) String() string {
    return fmt.Sprintf("CheckingAccount(id = %s, name = %s, balance = %.2f, overdraftlimit = %.2f)",
        ac.id, ac.name, ac.balance, ac.overdraftlimit)
}

func main() {
    account1 := Account{"1234-5678", "Justin Lin", 1000}
    account2 := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}

    // 顯示 Account(id = 1234-5678, name = Justin Lin, balance = 1000.00)
    fmt.Println(&account1)

    // 顯示 CheckingAccount(id = 1234-5678, name = Justin Lin, balance = 1000.00, overdraftlimit = 30000.00)
    fmt.Println(&account2)
}
```

# 實作某介面的型態有哪些？

來自 Java 之類語言的開發者，在認識 Go 的 `interface` 後可能會有些疑問，像是「如何知道某個介面的實現型態有哪些？」、「這個型態實現了哪些介面？」…並且會想在文件上尋找這類資訊，因為 Java 的文件中，會記錄某介面的實現類別有哪些。

這是因為 Java 中，介面型態與行為是結合在一起的。

在 Go 中不需要記錄這些，當開發者看到某 API 上定義可以接收某介面型態的值時，應該看看該介面定義了哪些行為，接著看看要傳入的值是否有實作這些行為，這樣就可以了，因為 Go 的介面重點是「行為」，不管 API 上定義的介面型態是什麼，只要行為符合都可以傳入。

也就是說 Go 中，介面型態與行為是分開的，應該重視的只有行為本身，本質上與動態定型語言中只重行為而非型態相同，因此「如何知道某個介面的實現型態有哪些？」、「這個型態實現了哪些介面？」這類問題也就不重要了！

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
