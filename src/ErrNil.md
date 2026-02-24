<div id="main" role="main" style="height: auto !important;">

<div class="header">

# err 是否 nil？

</div>

  

對於錯誤，Go 不採取例外處理機制，而是透過傳回 `error` 值來表示是否發生了什麼錯誤，最基本的做法就是：

``` prettyprint
if err != nil {
    // 做些什麼
}
```

然而，接觸 Go 不用多久就會發現，若要認真地檢查、處理錯誤，`if err != nil` 之類的程式碼就會到處充斥，特別是在進行 IO 之類的操作時更是如此，單純地 `if err != nil` 寫法最後會寫到懷疑人生，這麼寫真的是對的嗎？

這時可能會做的選擇之一是：就別檢查了吧！如果寫的是特定目的之程式、不太需要考慮太多狀況、不用考慮過多的穩固性、想要很快地寫出原型之類的，這個選擇可能是正確的，畢竟真要認真寫 Go 中的錯誤檢查，某些程度上就像 Java 中常被人嫌的受檢例外（Checked exception）一樣囉嗦，還好 Go 可以選擇不檢查…XD

只不過，如果想寫出較通用、具有穩固性的程式，錯誤檢查就是必需的，Go 也鼓勵開發者積極地檢查錯誤；那麼…乾脆全 `panic` 好了？

``` prettyprint
func check(err) {
    if err != nil {
        panic(err)
    }
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

這麼一來，遇到要檢查錯誤時，就呼叫 `check` 來檢查，這樣就能少寫些 `if err != nil` 了吧！這種做法其實並不建議，因為 `panic` 是 `panic`，`error` 是 `error`，`panic` 的場合，應該用在適用 `panic` 的場合，也就是那些實際上真的無法處理的錯誤，發生這類錯誤最重要的引發開發者恐慌，讓開發者知道要修改程式的演算，避免發生 `panic`。

`panic` 就像 Java 中發生 `RuntimeException`，其實不建議捕捉，而是停下程式，修正演算上的錯誤。

不過，可以想想為什麼會有人想在發生錯誤時，一律引發 `panic`，因為可以從目前的執行處中斷，就像例外處理機制中例外發生時，後續程式碼就不會執行那樣。

這就是以檢查是否有錯誤的方式，沒辦法直接做到的事，因為不在檢查出錯誤的時候進行 `return`、`break` 之類的動作，程式碼就會往下執行。

為了能在錯誤發生時中斷流程，就有可能寫出這類的程式碼：

``` prettyprint
_, err = fd.Write(p0[a:b])
if err != nil {
    return err
}
_, err = fd.Write(p1[c:d])
if err != nil {
    return err
}
_, err = fd.Write(p2[e:f])
if err != nil {
    return err
}
// 諸如此類
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

這段程式碼摘自〈[Errors are values](https://go.dev/blog/errors-are-values)〉，該文章中提到一個解決的方式是：

``` prettyprint
var err error
write := func(buf []byte) {
    if err != nil {
        return
    }
    _, err = w.Write(buf)
}
write(p0[a:b])
write(p1[c:d])
write(p2[e:f])
// 諸如此類
if err != nil {
    return err
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

這麼一來，每一次 `write` 呼叫時，就都會檢查 `err` 是否為 `nil`，如果不是 `nil` 就 `return`，實際上也就不會執行 `w.Write`，雖然程式碼上呼叫了 `write` 多次；然而，某次呼叫若發生了錯誤，後續的 `write` 並不會真正執行寫出的動作，而透過這個方式，可以將發生錯誤時要進行的動作，統整在最後檢查並執行。

匿名函式的方式建立了 Closure，捕捉了 `err` 變數，這麼一來就得做些迴避同名變數的問題，另外匿名函式的寫法也不是那麼簡明，因此文章中定義了：

``` prettyprint
type errWriter struct {
    w   io.Writer
    err error
}

func (ew *errWriter) write(buf []byte) {
    if ew.err != nil {
        return
    }
    _, ew.err = ew.w.Write(buf)
}
```

這麼一來，每個 `io.Writer` 可以有個別的 `err` 可以使用，而原本的程式就可以改寫為：

``` prettyprint
ew := &errWriter{w: fd}
ew.write(p0[a:b])
ew.write(p1[c:d])
ew.write(p2[e:f])
// 諸如此類
if ew.err != nil {
    return ew.err
}
```

在〈[bufio 套件](bufio.html)〉中看過的 `bufio.Writer` 就是這類的設計：

``` prettyprint
type Writer struct {
    err error
    buf []byte
    n   int
    wr  io.Writer
}

...略

func (b *Writer) Write(p []byte) (nn int, err error) {
    for len(p) > b.Available() && b.err == nil {
        var n int
        if b.Buffered() == 0 {
            // Large write, empty buffer.
            // Write directly from p to avoid copy.
            n, b.err = b.wr.Write(p)
        } else {
            n = copy(b.buf[b.n:], p)
            b.n += n
            b.Flush()
        }
        nn += n
        p = p[n:]
    }
    if b.err != nil {
        return nn, b.err
    }
    n := copy(b.buf[b.n:], p)
    b.n += n
    nn += n
    return nn, nil
}

... 略

func (b *Writer) Flush() error {
    if b.err != nil {
        return b.err
    }
    if b.n == 0 {
        return nil
    }
    n, err := b.wr.Write(b.buf[0:b.n])
    if n < b.n && err == nil {
        err = io.ErrShortWrite
    }
    if err != nil {
        if n > 0 && n < b.n {
            copy(b.buf[0:b.n-n], b.buf[n:b.n])
        }
        b.n -= n
        b.err = err
        return err
    }
    b.n = 0
    return nil
}
```

在 `b.err` 不為 `nil` 的情況下，實際上不會有實際的寫出，而 `Flush` 時，若 `b.err` 不為 `nil` 就會被 `return`，因此在使用 `bufio.Writer` 時，可以如下撰寫，在最後檢查

``` prettyprint
b := bufio.NewWriter(fd)
b.Write(p0[a:b])
b.Write(p1[c:d])
b.Write(p2[e:f])
// 諸如此類
if b.Flush() != nil {
    return b.Flush()
}
```

這個模式可以進一步應用，例如在〈[bufio 套件](bufio.html)〉中看過 `bufio.Scanner` 的使用，語意上比較高階：

``` prettyprint
scanner := bufio.NewScanner(f)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
if err := scanner.Err(); err != nil {
    panic(err)
}
```

`scanner.Scan()` 傳回布林值，表示是否掃描到下一行，沒有下一行或中途發生錯誤，就會傳回 `false`；然而迴圈檢查就只在乎有沒有下一行，離開迴圈後再來檢查錯誤，兩個程式區塊各司其職。

`bufio.Scanner` 本身的組成中有 `io.Reader` 與 `err`：

``` prettyprint
type Scanner struct {
    r            io.Reader 
    ...略
    err          error
    ...略
}
```

若你查看 `Scan` 方法的實作，會傳回 `false` 的情況之一，就是 `Scanner` 的 `err` 不是 `nil`：

``` prettyprint
    ...略
    if s.err != nil {
        // Shut it down.
        s.start = 0
        s.end = 0
        return false
    }
    ...略
```

Go 不以特定語法處理錯誤（例如 Java 使用 `try..catch`），正因為錯誤發生是傳回錯誤，也就會有許多方式可以檢查錯誤，這邊只是談到幾個可用的設計，重點在於觀察程式碼的需求，適時地重構，看看如何以設計的方式，優雅地處理錯誤，而不是避免檢查錯誤，如果一開始沒什麼方向，可以多觀察 Go 程式庫的原始碼實作中是怎麼處理錯誤的。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
