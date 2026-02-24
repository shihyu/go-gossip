<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 從標準輸入、輸出認識 io

</div>

  

若要輸出訊息至主控台，可以透過 `fmt` 的 `Print`、`Println`、`Printf` 等函式，如果要從主控台讀取使用者輸入，可以透過 `fmt` 的 `Scanf`、`Scanln` 等函式。例如：

``` go
package main

import "fmt"

func main() {
    fmt.Print("輸入名稱 年齡：")
    var name string
    var age int
    fmt.Scanf("%s %d", &name, &age)
    fmt.Printf("嗨！%s！今年 %d 歲了啊？", name, age)
}
```

`%s`、`%d` 是格式符號，在 Go 中稱為 verb，Go 可用的 verb 可以在 [`fmt`](https://pkg.go.dev/fmt/) 套件的文件中找到。

`Scanf` 就類似 C 語言中的 `scanf`，可以格式化地取得輸入，底下是個範例：

``` go
輸入名稱 年齡：Justin 45
嗨！Justin！今年 45 歲了啊？
```

在按下 Enter 鍵後，實際上還有個 CR（carriage return）字元還未掃描，如果只是要取得空白分隔的輸入，並以換行作為結束，可以使用 `Scanln`：

``` go
package main

import "fmt"

func main() {
    fmt.Print("輸入空白分隔的文字")
    var text1, text2 string
    fmt.Scanln(&text1, &text2)
    fmt.Println(text1)
    fmt.Println(text2)
}
```

如果是 `Scan` 的話，也是掃描以空白區隔的輸入，按下 Enter 鍵的 CR 字元，也會被視為空白。

`Println`、`Printf` 會使用標準輸出（Standout），如果想使用標準錯誤（Standard err）呢？可以透過 `Fprint`、`Fprintln`、`Fprintf` 等函式，第一個引數指定 `os.Stderr`。例如：

``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Fprintln(os.Stderr, "輸出至標準錯誤")
}
```

`os` 套件的 `Stderr` 代表標準錯誤，而 `Stdin`、`Stdout` 代表標準輸入與輸出，它們的型態是 `*os.File`，若願意的話，也可以直接操作它們，例如 `File` 定義了 `Read` 與 `Write` 方法，可以指定一個型態為 `byte[]` 的 slice，`Read` 會讀入同樣長度的資料至 slice，後者可以將同等長度的資料輸出。例如：

``` go
package main

import "os"

func main() {
    buf := make([]byte, 5);
    os.Stdout.Write([]byte("輸入五個數字："))
    os.Stdin.Read(buf)
    os.Stdout.Write(buf)
}   
```

實際上，[`os.File`](https://pkg.go.dev/os/#File) 可用的方法不只有 `Read`、`Write`，先留意這兩個方法的目的在於，這兩個方法分別符合 `io.Reader`、`io.Writer` 定義的行為：

``` go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

如果察看 `fmt` 的 `Fprint`、`Fprintln`、`Fprintf` 等函式，可以發現它們第一個參數宣告的型態並不是 `*os.File`，而是 `io.Writer`：

``` go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
```

類似地，Fscan 字樣開頭的幾個函式，第一個參數接受的是 `io.Reader`：

``` go
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
```

這表示，`fmt` 套件中這些函式，並不只能用於標準輸入、輸出或錯誤，例如，`strings.NewReader` 函式，可以指定字串，傳回 `*Reader`，這表示 `fmt` 的 `Fscanf` 等函式，可以從字串讀取輸入。例如：

``` go
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    data := `Justin 45
             Monica 42
             Irene 12`
    r := strings.NewReader(data)
    var name string
    var age int
    for {
        if _, err := fmt.Fscanln(r, &name, &age); err == io.EOF {
            break
        }
        fmt.Printf("%s: %d\n", name, age)
    }
}  
```

`Fscanln` 會傳回掃描的筆數，如果筆數少於指定的掃描數量，`err` 會指出原因，在檔案讀取結束（End of file）時，`err` 會是 `io.EOF`，在上例中，資料來源是個格式確定的字串，因此僅簡單地判斷 `err` 是否為 `io.EOF` 來結束掃描。

`os.File` 不過是具有 `io.Reader`、`io.Writer` 的行為罷了，`os.File` 代表檔案，也就是說 `Fprint`、`Fprintln`、`Fprintf`、`Fscan`、`Fscanln`、`Fscanf` 等函式，也可以用在檔案讀寫，其實標準輸入、輸出、錯誤等，也是被視為檔案的，這在 `os` 的 [file.go](https://go.dev/src/os/file.go) 可以看到：

``` go
var (
    Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
    Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```

因此 IO 之類的操作，在 Go 中非常靈活，一切都看 API 上可接受行為而定，不受型態之限制，這之後再從實際的例子中來談。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
