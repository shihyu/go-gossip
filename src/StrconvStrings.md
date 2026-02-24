<div id="main" role="main" style="height: auto !important;">

<div class="header">

# strconv、strings 套件

</div>

  

Go 的字串基本上是個 `[]byte`，在程式語言強弱型別的光譜中，Go 位於強型別的一端，對於字串與其他型態之間的轉換，往往得自行處理，在這方面，[`strconv` 套件](https://pkg.go.dev/strconv/)就提供了不少的函式。

例如，最常用的是將字串剖析為某個型態：

``` go
func ParseBool(str string) (bool, error)
func ParseFloat(s string, bitSize int) (float64, error)
func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (uint64, error)
```

若是剖析失敗，傳回的錯誤會是 `*NumError`：

``` go
type NumError struct {
    Func string // 來源函式（ParseBool、ParseInt、ParseUint、ParseFloat）
    Num  string // 輸入字串
    Err  error  // 失敗的源由（ErrRange、ErrSyntax 等）
}
```

如果要將其他型態附加至字串，可以使用 Append 名稱開頭的函式：

``` go
func AppendBool(dst []byte, b bool) []byte
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
func AppendInt(dst []byte, i int64, base int) []byte
func AppendQuote(dst []byte, s string) []byte
func AppendQuoteRune(dst []byte, r rune) []byte
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
func AppendQuoteToASCII(dst []byte, s string) []byte
func AppendQuoteToGraphic(dst []byte, s string) []byte
func AppendUint(dst []byte, i uint64, base int) []byte
```

以上的附加函式設計上接收 `[]byte`，Go 字串本質上是個 `[]byte`，呼叫這些函式時只要明確型態轉換就可以了，例如：

``` go
b := []byte("bool:")
b = strconv.AppendBool(b, true)
fmt.Println(string(b))
```

對於大量的字串附加處理，可以使用 [`strings` 套件](https://pkg.go.dev/strings/)的 `Builder`，一來操作上比較方便，二來可看看是否可取得較好的效能表現：

``` go
type Builder
    func (b *Builder) Cap() int
    func (b *Builder) Grow(n int)
    func (b *Builder) Len() int
    func (b *Builder) Reset()
    func (b *Builder) String() string
    func (b *Builder) Write(p []byte) (int, error)
    func (b *Builder) WriteByte(c byte) error
    func (b *Builder) WriteRune(r rune) (int, error)
    func (b *Builder) WriteString(s string) (int, error)
```

例如，來個簡單的評測：

``` go
package mypackage

import (
    "testing"
    "strings"
)

func plusAppend() string {
    c := ""
    for i := 0; i < 100000; i++ {
        c += "test"
    }
    return c
}

func buliderAppend() string {
    var b strings.Builder
    for i := 0; i < 100000; i++ {
        b.WriteString("test")
    }
    return b.String()
}

func BenchmarkPlusAppend(b *testing.B) {
    for i := 0; i < b.N; i++ {
        plusAppend()
    }
}

func BenchmarkBuilderAppend(b *testing.B) {
    for i := 0; i < b.N; i++ {
        buliderAppend()
    }
}
```

看一下效能上是否有差異：

``` go
C:\workspace\go-exercise>go test -bench="." mypackage
goos: windows
goarch: amd64
pkg: mypackage
BenchmarkPlusAppend-4                  1        4162865000 ns/op
BenchmarkBuilderAppend-4            1946            655490 ns/op
PASS
ok      mypackage       6.614s
```

如果想將字串當成是個 `io.Reader` 來源，可以使用 `strings.Reader`：

``` go
type Reader
    func NewReader(s string) *Reader
    func (r *Reader) Len() int
    func (r *Reader) Read(b []byte) (n int, err error)
    func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
    func (r *Reader) ReadByte() (byte, error)
    func (r *Reader) ReadRune() (ch rune, size int, err error)
    func (r *Reader) Reset(s string)
    func (r *Reader) Seek(offset int64, whence int) (int64, error)
    func (r *Reader) Size() int64
    func (r *Reader) UnreadByte() error
    func (r *Reader) UnreadRune() error
    func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```

`strings` 還有個 `Replacer`，用於一對一的字串取代：

``` go
type Replacer
    func NewReplacer(oldnew ...string) *Replacer
    func (r *Replacer) Replace(s string) string
    func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
```

什麼是一對一的取代呢？看看官方文件中提到的範例就知道了：

``` go
package main

import (
    "fmt"
    "strings"
)

func main() {
    r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
    fmt.Println(r.Replace("This is <b>HTML</b>!"))
}
```

其他對於字串的比較、分割、大小寫轉換等處理，`strings` 中提供了一系列的函式，[`strings` 套件](https://pkg.go.dev/strings/)的文件中都有程式碼示範。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
