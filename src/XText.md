<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 編碼轉換

</div>

  

不論從哪個面向，都可以看出 Go 獨厚 UTF-8，這可能是因為 Go 的設計者之一 Ken Thompson，也曾經參與了 UTF-8 的設計。

如果文字資料的來源並非 UTF-8 呢？例如，儲存時並非使用 UTF-8 的檔案？解決的方法之一，是將檔案另行儲存為 UTF-8，再使用 Go 來讀取，當然，並非所有的的場合都可以這麼做，另一個方式是，使用 [`golang.org/x/text`](https://pkg.go.dev/golang.org/x/text) 套件。

Go 除了本身自帶的標準套件之外，還有另外一系列官方的擴充套件（常稱 x/ repos），這些套件也是 Go 專案的一部份，只不過在相容性的維護上比較沒那麼嚴格。

在官方擴充套件中，`golang.org/x/text` 主要包含了文字編碼、轉換、國際化、本地化等文字性任務的套件。若是在模組專案中使用，通常直接 `import` 後執行 `go mod tidy` 即可；若要手動加入依賴，也可以使用：

``` prettyprint
go get golang.org/x/text@latest
```

文字編碼的轉換主要由 [`golang.org/x/text/transform`](https://pkg.go.dev/golang.org/x/text/transform) 套件來處理，看看其中的函式或結構方法，都會需要 `Transformer` 介面的實現，例如最基本的 `String`：

``` prettyprint
func String(t Transformer, s string) (result string, n int, err error)
```

`Transformer` 定義的主要是 `Transform` 方法，代表著編碼的轉換：

``` prettyprint
type Transformer interface {
    Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error)
    Reset()
}
```

`dst`、`src` 代表著同一文字兩個不同編碼的位元組，由於 Go 使用 UTF-8，從 UTF-8 轉換為其他編碼，這個動作稱為 `Encode`，從其他編碼轉換為 UTF-8，這個動作稱為 Decode。

Encode、Decode 的動作，分別由 [`golang.org/x/text/encoding`](https://pkg.go.dev/golang.org/x/text/encoding) 套件的 `Encoder`、`Decoder` 來處理，它們都是 `transform.Transformer` 的實現：

``` prettyprint
type Encoder struct {
    transform.Transformer
    ...
}

type Decoder struct {
    transform.Transformer
    ...
}
```

為了便於使用，`encoding` 定義了 `Encoding` 的行為：

``` prettyprint
type Encoding interface {
    NewDecoder() *Decoder
    NewEncoder() *Encoder
}
```

[`golang.org/x/text/encoding`](https://pkg.go.dev/golang.org/x/text/encoding) 套件之中，定義了不同的編碼轉換套件，例如，想處理 Big5（Code Page 950） 編碼轉換的話，需要 `golang.org/x/text/encoding/traditionalchinese` 套件，它的 `Big5` 就實現了 `Encoding`，因此想要獲得 UTF-8 \<-\> Big5 的 `Encoder`、`Decoder`，可以如下：

``` prettyprint
utf8ToBig5 := traditionalchinese.Big5.NewEncoder()
big5ToUtf8 := traditionalchinese.Big5.NewDecoder()
```

因此，若要讀取一個底層為 Big5 編碼的字串，轉換為 UTF-8 編碼字串，可以如下：

``` prettyprint
package main

import (
    "golang.org/x/text/encoding/traditionalchinese"
    "golang.org/x/text/transform"
    "fmt"
)

func main() {
    big5ToUTF8 := traditionalchinese.Big5.NewDecoder()
    big5Test := "\xb4\xfa\xb8\xd5" // 測試的 Big5 編碼
    utf8, _, _ := transform.String(big5ToUTF8, big5Test)
    fmt.Println(utf8) // 顯示「測試」
}
```

要將一個 UTF-8 編碼字串，轉換為 Big5 編碼的字串，可以如下：

``` prettyprint
package main

import (
    "golang.org/x/text/encoding/traditionalchinese"
    "golang.org/x/text/transform"
    "fmt"
)

func main() {
    utf8ToBig5 := traditionalchinese.Big5.NewEncoder()
    big5, _, _ := transform.String(utf8ToBig5, "測試")
    fmt.Printf("%q", big5)  // 顯示 "\xb4\xfa\xb8\xd5"
}
```

`transform` 也定義了 `Reader`、`Writer`，可以用來將 `Transformer` 與 `io.Reader`、`io.Writer` 包裹在一起：

``` prettyprint
type Reader
    func NewReader(r io.Reader, t Transformer) *Reader
    func (r *Reader) Read(p []byte) (int, error)

type Writer
    func NewWriter(w io.Writer, t Transformer) *Writer
    func (w *Writer) Close() error
    func (w *Writer) Write(data []byte) (n int, err error)
```

例如，想要讀取 Big5 文件的話，底下是個示範：

``` prettyprint
package main

import (
    "golang.org/x/text/encoding/traditionalchinese"
    "golang.org/x/text/transform"
    "fmt"
    "io"
    "os"
)

func printBig5(r io.Reader) error {
    var big5R = transform.NewReader(r, traditionalchinese.Big5.NewDecoder())

    b, err := io.ReadAll(big5R)
    fmt.Println(string(b))

    return err
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

    printBig5(f)
}
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
