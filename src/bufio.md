<div id="main" role="main" style="height: auto !important;">

<div class="header">

# bufio 套件

</div>

  

`io.Reader`、`io.Writer` 定義了基於位元組的讀寫行為，然而許多情況下，你會想要基於字串、行來進行讀寫，這可以透過 `bufio` 套件的 `bufio.Reader`、`bufio.Writer` 等達到。

`bufio.Reader` 可以透過 `NewReader`、`NewReaderSize` 指定 `io.Reader` 來建立實例，前者指定預設緩衝區大小 4096 位元組呼叫後者，`bufio.Reader` 在讀取來源時會從底層的 `io.Reader` 將資料讀入，在建立 `bufio.Reader` 實例之後，可以使用的方法有：

``` go
func (b *Reader) Buffered() int
func (b *Reader) Discard(n int) (discarded int, err error)
func (b *Reader) Peek(n int) ([]byte, error)
func (b *Reader) Read(p []byte) (n int, err error)
func (b *Reader) ReadByte() (byte, error)
func (b *Reader) ReadBytes(delim byte) ([]byte, error)
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
func (b *Reader) ReadRune() (r rune, size int, err error)
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
func (b *Reader) ReadString(delim byte) (string, error)
func (b *Reader) Reset(r io.Reader)
func (b *Reader) Size() int
func (b *Reader) UnreadByte() error
func (b *Reader) UnreadRune() error
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

因此對於逐行讀取一個 UTF-8 文字檔案來說，可以簡單地撰寫如下：

``` go
package main

import (
    "bufio"
    "os"
    "fmt"
    "io"
)

func printFile(f *os.File) (err error){
    var (
        r = bufio.NewReader(f)
        line string
    )
    for err == nil {
        line, err = r.ReadString('\n')
        fmt.Println(line)
    }
    if err == io.EOF {
        err = nil
    }
    return
}

func main() {
    var filename string
    fmt.Print("檔案名稱：")
    fmt.Scanf("%s", &filename);

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    printFile(f)
}
```

如果實際上是要讀取之後寫到另一個輸出，使用 `WriteTo` 方法更為方便：

``` go
package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    var filename string
    fmt.Print("檔案名稱：")
    fmt.Scanf("%s", &filename);

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    bufio.NewReader(f).WriteTo(os.Stdout)
}
```

Go 在 `io.WriteTo` 介面定義了 `WriteTo` 行為：

``` go
type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}
```

實際上 `bufio.Reader` 實作了 `io` 中一些介面，`io.WriteTo` 只是其中之一；類似地，如果要建立 `bufio.Writer` 實例，可以透過 `NewWriter`、`NewWriterSize` 函式，建立之後可用的方法如下：

``` go
func (b *Writer) Available() int
func (b *Writer) Buffered() int
func (b *Writer) Flush() error
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
func (b *Writer) Reset(w io.Writer)
func (b *Writer) Size() int
func (b *Writer) Write(p []byte) (nn int, err error)
func (b *Writer) WriteByte(c byte) error
func (b *Writer) WriteRune(r rune) (size int, err error)
func (b *Writer) WriteString(s string) (int, error)
```

`bufio.Writer` 實作了 `io` 中一些介面，像是 `io.ReadFrom`，因此，也可以如下在標準輸出中，顯示讀入的的檔案內容：

``` go
package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    var filename string
    fmt.Print("檔案名稱：")
    fmt.Scanf("%s", &filename);

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    w := bufio.NewWriter(os.Stdout)
    w.ReadFrom(f)
    w.Flush()
}
```

`NewWriter` 預設的緩衝區為 4096 位元組，由於這邊使用標準輸出，在緩衝區未滿前，資料不會寫出，可以使用 `Flush` 來出清緩衝區中的資料。

事實上，對於需要逐行讀取的需求，使用 `bufio.Scanner` 會比較方便，可以使用 `NewScanner` 來建立實例，建立之後有以下的方法可以使用：

``` go
func (s *Scanner) Buffer(buf []byte, max int)
func (s *Scanner) Bytes() []byte
func (s *Scanner) Err() error
func (s *Scanner) Scan() bool
func (s *Scanner) Split(split SplitFunc)
func (s *Scanner) Text() string
```

來看看讀取文字檔案的例子：

``` go
package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    var filename string
    fmt.Print("檔案名稱：")
    fmt.Scanf("%s", &filename);

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
