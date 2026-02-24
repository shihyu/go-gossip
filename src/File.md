<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 檔案操作

</div>

  

想要進行目錄、檔案等的操作，基本上就是查看 `os` 套件，可以使用的函式很多，逐一談好像也沒太大意義，基本上若對目錄、檔案以及權限等有所認識，應該查查文件、搜尋一些範例，大致就知道怎麼用吧！

無論如何，輸入輸出中最基本的就是檔案讀寫，至今為止看過，要開啟檔案進行讀取的話，使用的是 `os.Open` 函式，這會以唯讀方式開啟既有的檔案（否則會有 `PathError`）：

``` go
func Open(name string) (*File, error)
```

如果要指定讀寫方式與權限的話，要使用 `os.OpenFile`：

``` go
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```

`flag` 可以指定的常數有：

``` go
const (
    // 必須指定 O_RDONLY、O_WRONLY 或 O_RDWR
    O_RDONLY int = syscall.O_RDONLY // 唯讀
    O_WRONLY int = syscall.O_WRONLY // 唯寫
    O_RDWR   int = syscall.O_RDWR   // 讀寫
    // 接下來這些可以用 | 的方式附加行為
    O_APPEND int = syscall.O_APPEND // 寫入時使用附加方式
    O_CREATE int = syscall.O_CREAT  // 檔案不存在時建立新檔
    O_EXCL   int = syscall.O_EXCL   // 與 O_CREATE 併用，檔案必須不存在
    O_SYNC   int = syscall.O_SYNC   // 以同步 I/O 開啟
    O_TRUNC  int = syscall.O_TRUNC  // 檔案開啟時清空文件
)
```

`perm` 的話是檔案[八進位權限](https://en.wikipedia.org/wiki/Chmod#Octal_modes)，例如 0777；另外，還有個 `os.Create`，實現上就是使用 `OpenFile` 以 0666 的方式建立可讀寫的檔案（清空文件）：

``` go
func Create(name string) (*File, error) {
    return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
```

`Open`、`OpenFile` 或 `Create` 都會傳回 `*os.File`；另外還有個 `NewFile`，多數情況下用不到，主要是在將檔案描述（File descriptor）以 `*os.File` 來表示，例如，`os.Stdin`、`os.Stdout`、`os.Stderr`，在〈[從標準輸入、輸出認識 io](StdOutInErr.html)〉看過它的使用：

``` go
var (
    Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
    Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```

`syscall.Stdin`、`syscall.Stdout`、`syscall.Stderr` 分別是標準輸入、輸出、錯誤的檔案描述，這在 [`syscall` 的文件](https://pkg.go.dev/syscall/#pkg-variables)可以看到：

``` go
var (
    Stdin  = 0
    Stdout = 1
    Stderr = 2
)
```

`os.File` 實作了 `io.Reader`、`io.Writer` 等行為，因此只要知道〈[io.Reader、io.Writer](ReaderWriter.html)〉，剩下的就是查詢文件，看看有哪些方法可以使用，沒什麼特別需要示範的了，倒是若需要簡單的檔案讀寫，可以看看 [`ioutil` 套件](https://pkg.go.dev/io/ioutil/)（歷史 API），其中有些簡便的函式：

``` go
func NopCloser(r io.Reader) io.ReadCloser
func ReadAll(r io.Reader) ([]byte, error)
func ReadDir(dirname string) ([]os.FileInfo, error)
func ReadFile(filename string) ([]byte, error)
func TempDir(dir, prefix string) (name string, err error)
func TempFile(dir, pattern string) (f *os.File, err error)
func WriteFile(filename string, data []byte, perm os.FileMode) error
```

`ReadFile`、`WriteFile` 只要指定檔案名稱等，程式碼上不需要自行建立檔案、緩衝區之類的，這些函式在 [`ioutil` 套件](https://pkg.go.dev/io/ioutil/) 的文件中，都有範例可以參考。

補充（Go 1.16+）：`ioutil` 已棄用（deprecated），功能多已移到 `io` 與 `os`：

``` go
ioutil.ReadAll   -> io.ReadAll
ioutil.ReadFile  -> os.ReadFile
ioutil.WriteFile -> os.WriteFile
ioutil.TempDir   -> os.MkdirTemp
ioutil.TempFile  -> os.CreateTemp
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
