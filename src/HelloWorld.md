<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 來個 Hello, World

</div>

  

我在這邊原本使用的是 Windows 中的 Go 1.13 版本；若你使用目前的 Go 1.26，可至 [Go 的官方網站](https://go.dev/dl/) 下載安裝。

如果想來點不同的安裝方式，可以參考〈[門外漢的 Go 輕量開發環境](http://openhome.cc/Gossip/CodeData/DockerLayman/DockerLayman4.html)〉，在 Raspberry Pi 上的 Docker 容器中建立相關環境，就目前為止。

本文後半段會示範傳統 `GOPATH` 工作方式，這在 Go 1.26 仍有助於理解套件與建構流程，不過新專案通常會優先使用 Go 模組（`go mod init`，可搭配〈[模組入門](Module.html)〉）。

使用官方安裝程式時，通常不需要手動設定 `GOROOT`（安裝程式會處理）；實務上常見只要讓 Go 的 `bin` 目錄在 `PATH` 中即可。

# go run

要撰寫第一個 Hello, World 程式，你可以建立一個 main.go，在當中撰寫以下的內容：

``` prettyprint
package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
    fmt.Println("哈囉！世界！")
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

每個 .go 原始碼，都必須從 `package` 定義開始，而對於包括程式進入點 `main` 函式的 .go 原始碼，必須是在 `package main` 之中，為了要能輸出訊息，這邊使用了 `fmt` 套件（package）之中的 `Println` 函式，開頭的大寫 P 表示這是個公開的函式，可以在套件之外進行呼叫。

Go 的創建者之一也是 UTF-8 的創建者，因此，Go 可以直接處理多國語言，只要你確定編輯器編碼為 UTF-8 就可以了，如果你使用 vim，可以在 vim 的命令模式下輸入 `:set encoding=utf-8`，或者是在 .vimrc 之中增加一行 `set encoding=utf-8`。

Go 可以用直譯的方式來執行程式，第一個 Hello, World 程式就是這麼做的，執行 `go run` 指定你的原始碼檔名就可以了：

``` prettyprint
$ go run main.go
Hello, World
哈囉！世界！
```

# package 與 GOPATH

以下示範的是傳統 `GOPATH` 目錄配置（`src/`、`pkg/`、`bin/`）。在 Go 1.26 的新專案中，通常不需要手動建立這種結構，也不需要把專案放在 `GOPATH` 內，改用模組即可。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

那麼，一開始的 `package` 是怎麼回事？試著先來建立一個 hello.go：

``` prettyprint
package hello

import "fmt"

func HelloWorld() {
    fmt.Println("Hello, World")
}
```

記得，`package` 中定義的函式，名稱必須是以大寫開頭，其他套件外的程式，才能進行呼叫，若函式名稱是小寫，那麼會是套件中才可以使用的函式。

接著，原本的 main.go 修改為：

``` prettyprint
package main

import "hello"

func main() {
    hello.HelloWorld()
}
```

現在顯然地，main.go 中要用到方才建立的 `hello` 套件中的 `HelloWorld` 函式，這時 `package` 的設定就會發揮一下效用，你得將 hello.go 移到 src/hello 目錄之中，也就是目錄名稱必須符合 `package` 設定之名稱。

同樣地，你可以將 main.go 移到 src/main 目錄之中，以符合 `package` 的設定。

而 src 的位置，必須是在 `GOROOT` 或者是 `GOPATH` 的路徑中可以找到，當 Go 需要某套件中的元素時，會分別到這兩個環境變數的目錄之中，查看 src 中是否有相應於套件的原始碼存在。

為了方便，通常會設定 `GOPATH`，例如，指向目前的工作目錄：

``` prettyprint
set GOPATH=c:\workspace\go-exercise
```

如果沒有設定 `GOPATH` 的話，Go 預設會是使用者目錄的 go 目錄，雖然目前 `GOPATH` 中只一個目錄，不過 `GOPATH` 中可以設定數個目錄，現在我的 go-exercise 目錄底下會有這些東西：

``` prettyprint
go-exercise
          └─src
              ├─hello
              │      hello.go
              │
              └─main
                      main.go
```

接著在 go 目錄中執行指令 `go run src/main/main.go` 的話，你就會看到 Hello, World 了。

# go build

如果想編譯原始碼為可執行檔，那麼可以使用 `go build`，例如，直接在 go 目錄中執行 `go build src/main/main.go`，就會在執行指令的目錄下，產生一個名稱為 main.exe 的可執行檔，可執行檔的名稱是來自己指定的原始碼檔案主檔名，執行產生出來的可執行檔就會顯示 Hello, World。

你也可以建立一個 bin 目錄，然後執行 `go build -o bin/main.exe src/main/main.go`，這樣產生出來的可執行檔，就會被放在 bin 底下。

# go install

每次使用 `go build`，都是從原始碼編譯為可執行檔，這比較沒有效率，如果想要編譯時更有效率一些，可以使用 `go install`，例如，在目前既有的目錄與原始碼架構之下，於 go 目錄中執行 `go install hello` 的話，你就會發現有以下的內容：

``` prettyprint
go-exercise
        ├─bin
        │      main.exe
        │
        ├─pkg
        │  └─windows_amd64
        │          hello.a
        │
        └─src
            ├─hello
            │      hello.go
            │
            └─main
                    main.go
```

`go install packageName` 表示要安裝指定名稱的套件，如果是 `main` 套件，那麼會在 bin 中產生可執行檔，如果是公用套件，那麼會在 pkg 目錄的 `$GOOS`\_`$GOARCH` 目錄中產生 .a 檔案，你可以使用 `go env` 來查看 Go 使用到的環境變數，例如：

（補充：上面這段是以傳統 `GOPATH` 工作模式來理解；在現代模組模式下，編譯快取主要在 `GOCACHE`，模組原始碼快取在 `GOMODCACHE`，不一定會看到同樣的 `pkg/$GOOS_$GOARCH/*.a` 使用方式。）

``` prettyprint
set GO111MODULE=
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\Justin\AppData\Local\go-build
set GOENV=C:\Users\Justin\AppData\Roaming\go\env
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=C:\Users\Justin\go
set GOPRIVATE=
set GOPROXY=https://proxy.golang.org,direct
set GOROOT=C:\Winware\Go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=C:\Winware\Go\pkg\tool\windows_amd64
set GCCGO=gccgo
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\Justin\AppData\Local\Temp\go-build282125542=/tmp/go-build -gno-record-gcc-switches
```

.a 檔案是編譯過後的套件，因此，你看到的 hello.a，就是 hello.go 編譯之後的結果，如果編譯時需要某個套件，而對應的 .a 檔案存在，且原始碼自上次編譯後未曾經過修改，那麼就會直接使用 .a 檔案，而不是從原始碼開始編譯起。

# os.Args

那麼，如果想在執行 Go 程式時使用命令列引數呢？可以使用 `os` 套件的 `Args`，例如，寫一個 main.go：

``` prettyprint
package main

import "os"
import "fmt"

func main() {
    fmt.Printf("Command: %s\n", os.Args[0])
    fmt.Printf("Hello, %s\n", os.Args[1])
}
```

`os.Args` 是個陣列，索引從 0 開始，索引 0 會是編譯後的可執行檔名稱，索引 1 開始會是你提供的引數，例如，在執行過 go build 或 go install 之後，如下直接執行編譯出來的執行檔，會產生的訊息是…

``` prettyprint
$ ./bin/main Justin
Command: ./bin/main
Hello, Justin
```

# go doc

`fmt` 的 Printf，就像是 C 的 `printf`，可用的格式控制可參考 [Package fmt](https://pkg.go.dev/fmt) 的說明。實際上，Go 本身附帶了說明文件，可以執行 `go doc <pkg> <sym>[.<method>]` 來查詢說明。例如：

``` prettyprint
$ go doc fmt.Printf
func Printf(format string, a ...interface{}) (n int, err error)

    Printf formats according to a format specifier and writes to standard
    output. It returns the number of bytes written and any write error
    encountered.
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
