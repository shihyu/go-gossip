<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 結構組合

</div>

  

結構本身用來組織相關資料，可以將處理結構的相關函式定義為方法，類似物件導向程式語言中，使用類別定義值域與方法，那麼繼承呢？Go 語言並非以物件導向為主要典範的語言，沒有繼承的概念，不過可以使用組合代替繼承。

# 在組告之前

在〈[結構與方法](http://openhome.cc/Gossip/Go/Method.html)〉中使用 `struct` 定義了 `Account`，如果今天你想定義一個支票帳戶，方式之一是…

``` go
type CheckingAccount struct {
    id string
    name string
    balance float64
    overdraftlimit float64
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

這是個很尋常的作法，也許你想將 `id`、`name` 與 `balance` 組織在一起：

``` go
package main

import "fmt"

type CheckingAccount struct {
    account struct {
        id      string
        name    string
        balance float64
    }
    overdraftlimit float64
}

func main() {
    checking := CheckingAccount{}
    checking.account = struct {
        id      string
        name    string
        balance float64
    }{"1234-5678", "Justin Lin", 1000}
    checking.overdraftlimit = 30000

    fmt.Println(checking)                // {{1234-5678 Justin Lin 1000} 30000}
    fmt.Println(checking.account)        // {1234-5678 Justin Lin 1000}
    fmt.Println(checking.account.name)   // Justin Lin
    fmt.Println(checking.overdraftlimit) // 30000
}
```

這是一種方式，不過使用起來麻煩，或許你可以這麼做：

``` go
package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

type CheckingAccount struct {
    account        Account
    overdraftlimit float64
}

func main() {
    checking := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}

    fmt.Println(checking)                // {{1234-5678 Justin Lin 1000} 30000}
    fmt.Println(checking.account)        // {1234-5678 Justin Lin 1000}
    fmt.Println(checking.account.name)   // Justin Lin
    fmt.Println(checking.overdraftlimit) // 300000
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

看來還不錯，不過，如果想要 `fmt.Println(checking.name)` 就能取得名稱的話，這種寫法行不通！

# 結構值域的查找

在定義結構時，可以將另一已定義的結構直接內嵌：

``` go
package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

type CheckingAccount struct {
    Account
    overdraftlimit float64
}

func main() {
    account := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}

    fmt.Println(account)                // {{1234-5678 Justin Lin 1000} 30000}
    fmt.Println(account.id)             // 1234-5678
    fmt.Println(account.name)           // Justin
    fmt.Println(account.balance)        // 1000
    fmt.Println(account.overdraftlimit) // 30000
}
```

這稱為型態內嵌（type embedding），`Account` 被稱為 `CheckingAccount` 的內部型態，反之，`CheckingAccount` 是 `Account` 的外部型態，雖然是透過 `account.id`、`account.name`、`account.balance` 來存取，不過內部型態提昇，令內部型態定義的值域為可見。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

那麼，如果想要明確地透過 `Account` 的結構來存取呢？也是可以的：

``` go
package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

type CheckingAccount struct {
    Account
    overdraftlimit float64
}

func main() {
    account := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}

    fmt.Println(account)                 // {{1234-5678 Justin Lin 1000} 30000}
    fmt.Println(account.Account.id)      // 1234-5678
    fmt.Println(account.Account.name)    // Justin
    fmt.Println(account.Account.balance) // 1000
    fmt.Println(account.overdraftlimit)  // 30000
}
```

雖然內部型態會提昇，然而，若外部型態中定義了同名值域，就會直接取得外部型態的值域，因此，如果 `CheckingAccount` 定義了相同的值域 `balance`，如果透過 `account.balance`，結果會是找到 `CheckingAccount` 定義的 `balance`，如果想明確找到 `Account` 的 `balance`，可以指定 `Account` 作為前置：

``` go
package main

import "fmt"

type Account struct {
    id      string
    name    string
    balance float64
}

type CheckingAccount struct {
    Account
    balance        float64
    overdraftlimit float64
}

func main() {
    account := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 2000, 30000}

    fmt.Println(account.balance)         // 2000
    fmt.Println(account.Account.balance) // 1000
}
```

無論是結構值域或是方法，若來自兩個結構的值域或方法產生了同名衝突，Go 會有 ambiguous selector 的錯誤提示，此時你必須明確指定結構名稱，指定使用來自哪個結構的值域或方法。

# 方法的查找

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

如果內部型態原本定義了方法，這些方法也是查找時的對象：

``` go
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

type CheckingAccount struct {
    Account
    overdraftlimit float64
}

func main() {
    account := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}
    account.Deposit(2000)
    account.Withdraw(500)
    fmt.Println(account) // {{1234-5678 Justin Lin 2500} 30000}
}
```

類似地，若外部型態中定義了同名的方法，那麼就會使用該方法，這類似重新定義（Override）的概念：

``` go
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
    account := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}
    account.Deposit(2000)
    if err := account.Withdraw(50000); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(account)
    }
}
```

在上面的範例中，會顯示「超出信用額度」的訊息，拿掉 `func (account *CheckingAccount) Withdraw(amount float64)` 該函式的定義，則會顯示「餘額不足」的訊息。

如果想指定使用 `Account` 的 `Withdraw` 函式，也還是可以的：

``` go
func main() {
    account := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}
    account.Deposit(2000)
    if err := account.Account.Withdraw(50000); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(account)
    }
}
```

雖然可以實現方法重新定義的概念，不過，單純只是如上定義的話，並不支援多型的概念，因為一開始這麼指定就會出錯了：

``` go
// cannot use CheckingAccount literal (type CheckingAccount) as type Account in assignment
var account Account = CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}
```

若想實作出多型的概念，必須使用 `interface`，這在之後的文件會加以說明。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
