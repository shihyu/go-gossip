<div id="main" role="main" style="height: auto !important;">

<div class="header">

# errors 套件

</div>

  

在 Go 1.13 之前，`errors` 套件只公開了 `New` 函式，從 Go 1.13 之後，增加了 `Is`、`As`、與 `Unwrap` 函式。

`Is` 函式是用於取代 `==` 判斷錯誤的場合，例如以下的程式片段：

``` go
if err == io.EOF {
    ...
}
```

可以改用 `Is` 函式：

``` go
if errors.Is(err, io.EOF) {
    ...
}
```

`Is` 也可以用於判斷 `nil`，`err` 若有實作 `Is` 方法，也可以使用 `Is` 函式來判斷，因為 `Is` 函式的原始碼是這麼實作的：

``` go
func Is(err, target error) bool {
    if target == nil {
        return err == target
    }

    isComparable := reflectlite.TypeOf(target).Comparable()
    for {
        if isComparable && err == target {
            return true
        }
        if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
            return true
        }
        // TODO: consider supporing target.Is(err). This would allow
        // user-definable predicates, but also may allow for coping with sloppy
        // APIs, thereby making it easier to get away with them.
        if err = Unwrap(err); err == nil {
            return false
        }
    }
}
```

（從原始碼中的註解可以看到，未來可能進一步支援 `target` 實作 `Is` 方法的情況。）

`As` 函式是用於取代型態斷言判斷錯誤類型的場合，例如以下的程式片段：

``` go
if e, ok := err.(*PathError); ok {
    ...
}
```

可以改用 `As` 函式：

``` go
var e *PathError
if errors.As(err, &e) {
    ...
}
```

來看看 `As` 函式的實作：

``` go
func As(err error, target interface{}) bool {
    if target == nil {
        panic("errors: target cannot be nil")
    }
    val := reflectlite.ValueOf(target)
    typ := val.Type()
    if typ.Kind() != reflectlite.Ptr || val.IsNil() {
        panic("errors: target must be a non-nil pointer")
    }
    if e := typ.Elem(); e.Kind() != reflectlite.Interface && !e.Implements(errorType) {
        panic("errors: *target must be interface or implement error")
    }
    targetType := typ.Elem()
    for err != nil {
        if reflectlite.TypeOf(err).AssignableTo(targetType) {
            val.Elem().Set(reflectlite.ValueOf(err))
            return true
        }
        if x, ok := err.(interface{ As(interface{}) bool }); ok && x.As(target) {
            return true
        }
        err = Unwrap(err)
    }
    return false
}
```

`target` 若不是指標就會 `panic`；另外，`err` 可以是個實作 `As` 方法的實例。

在 `Is` 與 `As` 的實作中，都看到了 `Unwrap` 函式：

``` go
func Unwrap(err error) error {
    u, ok := err.(interface {
        Unwrap() error
    })
    if !ok {
        return nil
    }
    return u.Unwrap()
} 
```

從 Go 1.13 開始，錯誤可以實作 `Unwrap` 方法，如果 `e1.Unwrap()` 可以得到 `e2`，那麼 `e1` 實例包裹了 `e2`，因此，對於需要包含根源錯誤的情況，保存根源錯誤的欄位不需要是公開的，可以透過 `Unwrap` 來傳回，`Unwrap` 為取得包裹的錯誤提供了統一的名稱。

`fmt` 套件有個 `Errorf` 函式，可以格式化字串並傳回 `error` 實例，在 Go 1.13 之前的版本，就只是將格式化後的字串傳給 `errors.New`；從 Go 1.13 開始，`Errorf` 支援 `%w`，這會令傳回的 `error` 實例會包裹指定的錯誤，並具有 `Unwrap` 方法。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
