<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 錯誤的比對

</div>

  

如果函式或方法傳回錯誤，要比對的不單只是 `nil` 與否，例如，讀取檔案時，會需要判斷傳回的錯誤是否為 `io.EOF`，那麼 `io.EOF` 這些錯誤是什麼呢？在 `io` 套件的 [io.go](https://go.dev/src/io/io.go) 原始碼中可以看到，它們就是個 `errors.New` 建出的值罷了：

``` prettyprint
var ErrShortWrite = errors.New("short write")
var ErrShortBuffer = errors.New("short buffer")
var EOF = errors.New("EOF")
var ErrUnexpectedEOF = errors.New("unexpected EOF")
var ErrNoProgress = errors.New("multiple Read calls return no data or error")
```

在 `errors` 套件的 [errors.go](https://go.dev/src/errors/errors.go) 可以看到，`errors.New` 建立的是個結構值，只有一個 `string` 欄位，並且實作了 `Error` 方法：

``` prettyprint
func New(text string) error {
    return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

字串是可以比較的（Comparable），`errorString` 結構也是個可以比較的，因此可以直接使用 `==` 來比較錯誤是否為 `io.EOF` 等，在開發自己的應用程式或程式庫時，對於通用、簡單的錯誤，也可以如上定義。

`errors.New` 建立的實例，能攜帶的資訊就只是字串罷了，如果錯誤發生時，需要傳遞更多的環境資訊，怎麼辦呢？

在方法宣告傳回錯誤時的 `error` 其實是個內建的介面，定義的正是 `Error` 方法：

``` prettyprint
type error interface {
    Error() string
}
```

也就是說，只要有實作 `Error` 方法，都可以作為 `error` 實例傳回，例如，`os.PathError` 在 `os` 套件的 [error.go](https://go.dev/src/os/error.go) 是這麼定義的：

``` prettyprint
type PathError struct {
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

func (e *PathError) Unwrap() error { return e.Err }

func (e *PathError) Timeout() bool {
    t, ok := e.Err.(timeout)
    return ok && t.Timeout()
}
```

也就是說，若錯誤是 `PathError` 實例，可以有透過欄位或者是方法來取得更多資訊，例如：

``` prettyprint
if e, ok := err.(*PathError); ok {
    // 透過 e 取得欄位或呼叫方法
}
```

若要多種類型要判斷，可以使用型態 `switch` 語法，例如 `os` 套件的 [error.go](https://go.dev/src/os/error.go) 內部實作就有個例子：

``` prettyprint
func underlyingError(err error) error {
    switch err := err.(type) {
    case *PathError:
        return err.Err
    case *LinkError:
        return err.Err
    case *SyscallError:
        return err.Err
    }
    return err
}
```

像 `PathError` 中還包含了 `Err` 欄位，這並非必要，其應用的情境是在呼叫某函式時檢查到錯誤，除了建立另一個錯誤實例收集當時的環境資訊之外，你可能會想要包裹來源的錯誤實例，以便後續呼叫者可以進一步檢視錯誤根源。

然而，當某個錯誤包裹了另一個錯誤，也就表示後續呼叫者得知道該錯誤的細節，如果這些細節來自另一個底層，而你不想曝露，就不要直接包裹它，這時在目前應用程式或程式庫的抽象層面中，抽取出來源錯誤中的資訊，包裝為目前層次的錯誤就可以了。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
