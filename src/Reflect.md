<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 反射入門

</div>

  

反射（Reflection）是探知資料自身結構的一種能力，不同的語言提供不同的反射機制，在 Go 語言中，反射的能力主要由 `reflect` 套件提供。

# 資料的 Type

在先前的文件中，有時會用到 `reflect.TypeOf()` 來顯示資料的型態名稱，實際上，`reflect.TypeOf()` 傳回 `Type` 的實例，`Type` 是個介面定義，目前包含了以下的方法定義：

``` prettyprint
type Type interface {
    Align() int
    FieldAlign() int
    Method(int) Method
    MethodByName(string) (Method, bool)
    NumMethod() int
    Name() string
    PkgPath() string
    Size() uintptr
    String() string
    Kind() Kind
    Implements(u Type) bool
    AssignableTo(u Type) bool
    ConvertibleTo(u Type) bool
    Comparable() bool
    Bits() int
    ChanDir() ChanDir
    IsVariadic() bool
    Elem() Type
    Field(i int) StructField
    FieldByIndex(index []int) StructField
    FieldByName(name string) (StructField, bool)
    FieldByNameFunc(match func(string) bool) (StructField, bool)
    In(i int) Type
    Key() Type
    Len() int
    NumField() int
    NumIn() int
    NumOut() int
    Out(i int) Type
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

因此，你可以透過 `Type` 的方法定義，取得某個型態的相關結構資訊，舉例來說：

``` prettyprint
package main

import (
    "fmt"
    "reflect"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func main() {
    account := Account{"X123", "Justin Lin", 1000}
    t := reflect.TypeOf(account)
    fmt.Println(t.Kind())   // struct
    fmt.Println(t.String()) // main.Account
    /*
       底下顯示
       id string
       name string
       balance float64
    */
    for i, n := 0, t.NumField(); i < n; i++ {
        f := t.Field(i)
        fmt.Println(f.Name, f.Type)
    }
}
```

如果 `reflect.TypeOf()` 接受的是個指標，因為指標實際上只是個位址值，必須要透過 `Type` 的 `Elem` 方法取得指標的目標 `Type`，才能取得型態的相關成員：

``` prettyprint
package main

import (
    "errors"
    "fmt"
    "reflect"
)

type Savings interface {
    Deposit(amount float64) error
    Withdraw(amount float64) error
}

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("必須存入正數")
    }
    ac.balance += amount
    return nil
}

func (ac *Account) Withdraw(amount float64) error {
    if amount > ac.balance {
        return errors.New("餘額不足")
    }
    ac.balance -= amount
    return nil
}

func main() {
    var savings Savings = &Account{"X123", "Justin Lin", 1000}
    t := reflect.TypeOf(savings)

    for i, n := 0, t.NumMethod(); i < n; i++ {
        f := t.Method(i)
        fmt.Println(f.Name, f.Type)
    }

    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }

    fmt.Println(t.Kind())
    fmt.Println(t.String())
    for i, n := 0, t.NumField(); i < n; i++ {
        f := t.Field(i)
        fmt.Println(f.Name, f.Type)
    }
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

有上面的範例中，也示範了如何取得介面定義的方法資訊，這個範例會顯示以下的結果：

``` prettyprint
Deposit func(*main.Account, float64) error
Withdraw func(*main.Account, float64) error
struct
main.Account
id string
name string
```

# 資料的 Kind

上面的範例中，使用了 `Type` 的 `Kind()` 方法，這會傳回 `Kind` 列舉值：

``` prettyprint
type Kind uint

const (
    Invalid Kind = iota
    Bool
    Int
    Int8
    Int16
    Int32
    Int64
    Uint
    Uint8
    Uint16
    Uint32
    Uint64
    Uintptr
    Float32
    Float64
    Complex64
    Complex128
    Array
    Chan
    Func
    Interface
    Map
    Ptr
    Slice
    String
    Struct
    UnsafePointer
)
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

以下是個簡單的型態測試：

``` prettyprint
package main

import (
    "fmt"
    "reflect"
)

type Duck struct {
    name string
}

func main() {
    values := [...](interface{}){
        Duck{"Justin"},
        Duck{"Monica"},
        [...]int{1, 2, 3, 4, 5},
        map[string]int{"caterpillar": 123456, "monica": 54321},
        10,
    }

    for _, value := range values {
        switch t := reflect.TypeOf(value); t.Kind() {
        case reflect.Struct:
            fmt.Println("it's a struct.")
        case reflect.Array:
            fmt.Println("it's a array.")
        case reflect.Map:
            fmt.Println("it's a map.")
        case reflect.Int:
            fmt.Println("it's a integer.")
        default:
            fmt.Println("非預期之型態")
        }
    }
}
```

# 資料的 Value

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

如果想實際獲得資料的值，可以使用 `reflect.ValueOf()` 函式，這會傳回 `Value` 實例，`Value` 是個結構，定義了一些方法可以使用，可用來取得實際的值，例如：

``` prettyprint
package main

import (
    "fmt"
    "reflect"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func main() {
    x := 10
    vx := reflect.ValueOf(x)
    fmt.Printf("x = %d\n", vx.Int())

    account := Account{"X123", "Justin Lin", 1000}
    vacct := reflect.ValueOf(account)
    fmt.Printf("id = %s\n", vacct.FieldByName("id").String())
    fmt.Printf("name = %s\n", vacct.FieldByName("name").String())
    fmt.Printf("balance = %.2f\n", vacct.FieldByName("balance").Float())
}
```

如果是個指標，一樣也是要透過 `Elem()` 方法取得目標值，例如：

``` prettyprint
package main

import (
    "fmt"
    "reflect"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func main() {
    x := 10
    vx := reflect.ValueOf(&x)
    fmt.Printf("x = %d\n", vx.Elem().Int())

    account := &Account{"X123", "Justin Lin", 1000}
    vacct := reflect.ValueOf(account).Elem()
    fmt.Printf("id = %s\n", vacct.FieldByName("id").String())
    fmt.Printf("name = %s\n", vacct.FieldByName("name").String())
    fmt.Printf("balance = %.2f\n", vacct.FieldByName("balance").Float())
}
```

可以透過 `Value` 對值進行變動，不過，`Value` 必須是可定址的，具體來說，就是 `reflect.ValueOf()` 必須接受指標：

``` prettyprint
package main

import (
    "fmt"
    "reflect"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func main() {
    x := 10
    vx := reflect.ValueOf(&x).Elem()
    fmt.Printf("x = %d\n", vx.Int()) // x = 10

    vx.SetInt(20)
    fmt.Printf("x = %d\n", x) // x = 20
}
```

上面的例子若改成以下，就會出現錯誤：

``` prettyprint
package main

import (
    "fmt"
    "reflect"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func main() {
    x := 10
    vx := reflect.ValueOf(x)
    fmt.Printf("x = %d\n", vx.Int())

    vx.SetInt(20) // panic: reflect: reflect.Value.SetInt using unaddressable value
    fmt.Printf("x = %d\n", x)
}
```

技術上來說，上面的例子，只是傳了 `x` 的值複本給 `reflect.ValueOf()`，因此，對其設值並無意義。

若對反射想進一步研究，可以參考〈[The Laws of Reflection](https://go.dev/blog/laws-of-reflection)〉。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
