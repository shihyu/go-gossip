<div id="main" role="main" style="height: auto !important;">

<div class="header">

# bytes 套件

</div>

  

Go 字串的本質是 `[]byte`，如果想基於位元組來處理字串，或者是想處理其他來源的 `[]byte`，可以使用 [`bytes`](https://pkg.go.dev/bytes/) 套件。

因為 Go 字串本質上就是一組 Unicode 碼點的 UTF-8 編碼位元組，[`bytes`](https://pkg.go.dev/bytes/) 與 [`strings`](https://pkg.go.dev/strings/) 套件中提供的函式，有著很大的相似性，只不過前者針對 `[]byte`，後者針對 `string`…唔…好像在說廢話…也就是說…儘管兩者提供的函式在名稱上有重疊，除了函式上的參數或傳回型態不同之外，兩者處理的粒度等也不同，例如 Compare，一個是逐一比較位元組，另一個是逐一比較 Unicode 碼點。

類似地，對於頻繁性的字串操作，可以使用 `strings.Builder`，對於對於頻繁性的位元組操作，可以使用 `bytes.Buffer`：

``` go
type Buffer
    func NewBuffer(buf []byte) *Buffer
    func NewBufferString(s string) *Buffer
    func (b *Buffer) Bytes() []byte
    func (b *Buffer) Cap() int
    func (b *Buffer) Grow(n int)
    func (b *Buffer) Len() int
    func (b *Buffer) Next(n int) []byte
    func (b *Buffer) Read(p []byte) (n int, err error)
    func (b *Buffer) ReadByte() (byte, error)
    func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
    func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)
    func (b *Buffer) ReadRune() (r rune, size int, err error)
    func (b *Buffer) ReadString(delim byte) (line string, err error)
    func (b *Buffer) Reset()
    func (b *Buffer) String() string
    func (b *Buffer) Truncate(n int)
    func (b *Buffer) UnreadByte() error
    func (b *Buffer) UnreadRune() error
    func (b *Buffer) Write(p []byte) (n int, err error)
    func (b *Buffer) WriteByte(c byte) error
    func (b *Buffer) WriteRune(r rune) (n int, err error)
    func (b *Buffer) WriteString(s string) (n int, err error)
    func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)
```

建立 `Buffer` 時可以使用 `NewBuffer` 指定初始的位元組大小，如果你想要處理的是字串的 UTF-8 位元組，可以使用 `NewBufferString`。例如，來簡單地針對中文做百分比編碼：

``` go
package main

import (
    "fmt"
    "bytes"
    "strings"
)

func encodeURI(s string) string {
    buf := bytes.NewBufferString(s)

    var builder strings.Builder
    for {
        b, e := buf.ReadByte()
        if e != nil {
            break
        }
        builder.WriteString(fmt.Sprintf("%%%X", b))
    }

    return builder.String()
}

func main() {
    fmt.Println(encodeURI("良葛格")) // %E8%89%AF%E8%91%9B%E6%A0%BC
}
```

類似地，你也可以透過 `bytes.Reader`，將 `[]byte` 作為來源讀取：

``` go
type Reader
    func NewReader(b []byte) *Reader
    func (r *Reader) Len() int
    func (r *Reader) Read(b []byte) (n int, err error)
    func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
    func (r *Reader) ReadByte() (byte, error)
    func (r *Reader) ReadRune() (ch rune, size int, err error)
    func (r *Reader) Reset(b []byte)
    func (r *Reader) Seek(offset int64, whence int) (int64, error)
    func (r *Reader) Size() int64
    func (r *Reader) UnreadByte() error
    func (r *Reader) UnreadRune() error
    func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
