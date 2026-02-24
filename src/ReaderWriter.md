<div id="main" role="main" style="height: auto !important;">

<div class="header">

# io.Reader、io.Writer

</div>

  

在〈[從標準輸入、輸出認識 io](StdOutInErr.html)〉中談到了 `io.Reader`、`io.Writer`，在 Go 中，這兩個介面抽象化了輸入、輸出，認識這兩個介面分別定義的 `Read`、`Write` 行為，是掌握 Go 中輸入、輸出的基礎。

`io.Reader` 定義的 `Read` 行為，可以在 [`type Reader`](https://pkg.go.dev/io/#Reader) 查看：

``` prettyprint
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

對於呼叫者來說，`Read` 會將資料讀入 `p`，並傳回讀入的位元組數 `n`，`n` 會是 0 到不大於 `len(p)` 的整數，如果 `n` 不是 0 但不足 `len(p)`，應該先處理已讀取的位元組，這時 `err` 可能不是 `nil`（例如檔案結尾，可能會傳回 `io.EOF`），無論如何，在這之後 `Read`，`n` 會是 0 而 `err` 會是 `io.EOF`。

例如，若要讀取一個文字檔案，其中以 UTF-8 儲存中文，可以如下：

``` prettyprint
package main

import (
    "fmt"
    "io"
    "os"
)

func printUTF8TC(r io.Reader) (err error) {
    var (
        buf = make([]byte, 3)
        n int
    )

    for err == nil {
        n, err = r.Read(buf)
        fmt.Print(string(buf[:n]))
    }
    if err == io.EOF {
        err = nil
    }
    return
}

func main() {
    fmt.Print("檔案來源：")
    var filename string
    fmt.Scanf("%s", &filename)

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    printUTF8TC(f)
}
```

`io.Writer` 定義的 `Write` 行為，可以在 [`type Writer`](https://pkg.go.dev/io/#Writer) 查看：

``` prettyprint
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

`Write` 會將 `p` 輸出並傳回實際輸出的位元組，`n` 會是 0 到不大於 `len(p)` 的整數，如果 `n < len(p)`，那麼 `err` 不會是 `nil`。

來寫個 `Copy` 函式好了，可以將 `io.Reader` 的資料直接寫到 `io.Writer`：

``` prettyprint
package main

import (
    "fmt"
    "io"
    "os"
)

func write(w io.Writer, buf []byte, n int) (err error) {
    nw, ew := w.Write(buf[:n])
    if ew != nil {
        return ew
    }
    if n != nw {
        return io.ErrShortWrite
    }
    return nil
}

func Copy(w io.Writer, r io.Reader) (err error) {
    buf := make([]byte, 32 * 1024)
    for {
        nr, er := r.Read(buf)
        if nr > 0 {
            err = write(w, buf, nr)
            if err != nil {
                return
            }
        }
        if er != nil {
            if er != io.EOF {
                err = er
            }
            return
        }
    }
}

func main() {
    fmt.Print("檔案來源：")
    var filename string
    fmt.Scanf("%s", &filename)

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    Copy(os.Stdout, f)
}
```

在這個例子中，可以將指定的檔案讀入並顯示在主控台中，這是因為 `os.Stdout` 具有 `io.Writer` 的行為。實際上，`io.Copy` 就提供了這個功能：

``` prettyprint
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    fmt.Print("檔案來源：")
    var filename string
    fmt.Scanf("%s", &filename)

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    io.Copy(os.Stdout, f)
}
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
