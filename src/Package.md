<div id="main" role="main" style="height: auto !important;">

<div class="header">

# Go 套件管理

</div>

  

本章主要是早期 `GOPATH` 時代的套件管理方式（歷史脈絡）。在 Go 1.26 的新專案中，通常會優先使用 Go 模組（`go mod init`、`go mod tidy`），可搭配〈[模組入門](Module.html)〉閱讀。

在〈[來個 Hello, World](http://openhome.cc/Gossip/Go/HelloWorld.html)〉中，你已經看到 Go 開發中，一個 workspace 的基本樣貌，你可以看到，裏頭會有 src、pkg、bin 目錄，你會設置 `GOPATH` 環境變數指向這個目錄，這些都是當時的規範，正如〈[How to Write Go Code](https://go.dev/doc/code.html)〉中說到的：

> *The go tool is designed to work with open source code maintained in public repositories. Although you don't need to publish your code, the model for how the environment is set up works the same whether you do or not.*

在〈[來個 Hello, World](http://openhome.cc/Gossip/Go/HelloWorld.html)〉已經稍微瞭解了 `package` 與 `GOPATH` 的關係，原始碼會是在 `GOPATH` 中設定的目錄之 src 中，並有著對照於 `package` 設定名稱之目錄包括著它，當 Go 的工具（`go build`、`go install` 等）需要原始碼時，會到 `GOROOT` 底下，或者是 `GOPATH` 底下，查看是否有相應於套件的原始碼存在，編譯出來的結果，會是在相對應的 pkg 或 bin 底下。

# 本地套件

在當時，為了簡化說明，原始碼主檔名故意與 `package` 設定的名稱同名，這不是必要的，一個相應於 `package` 的目錄底下，可以有許多個原始碼，而每個原始碼開頭，只要 `package` 設定的名稱都與目錄相符就可以了。例如，你可以有個原始碼是 hello.go，位於 src/goexample 底下：

``` go
package goexample

import "fmt"

func Hello() {
    fmt.Println("Hello")
}
```

還可以有個 hi.go，位於 src/goexample 底下：

``` go
package goexample

import "fmt"

func Hi() {
    fmt.Println("Hi")
}
```

也就是說，一個 `package` 可以有數個原始碼檔案，各自組織自己的任務，在執行 `go install goexample` 之後，上面兩個原始碼會在 pkg 目錄的 `$GOOS`\_`$GOARCH` 目錄中產生 goexample.a 檔案。這包括了 `goexample` 套件編譯後的結果，如果想使用 `goexample` 套件的功能，只需要撰寫個 main.go：

``` go
package main

import "goexample"

func main() {
    goexample.Hi()
    goexample.Hello()
}
```

你可以在套件目錄之前增加父目錄，例如，可以建立一個 src/cc/openhome 目錄，然後將方才的 hello.go 與 hi.go 移至該目錄之中，接著執行 `go install cc/openhome/goexample`，那麼，在 pkg 目錄的 \$GOOS\_\$GOARCH 目錄中，會產生對應的 cc/openhome 目錄，其中放置著 goexample.a 檔案，想要使用這個套件的話，可以撰寫個 main.go：

``` go
package main

import "cc/openhome/goexample"

func main() {
    goexample.Hi()
    goexample.Hello()
}
```

# 遠端套件

由於 Go 的 workspace 設置，都必須是如此規範，因此，若你想將原始碼發佈給他人使用時就很方便，例如，你可以建立 src/github.com/JustinSDK 目錄，然後將方才的 goexample 目錄移到 src/github.com/JustinSDK 當中，這麼一來，顯然地，你的 main.go 就要改成：

``` go
package main

import "github.com/JustinSDK/goexample"

func main() {
    goexample.Hi()
    goexample.Hello()
}
```

也就是說，你可以直接將 /src/github.com/JustinSDK/goexample 當作檔案庫（repository）發佈到 Github，那麼，其他人需要你的原始碼時，在當時常會使用 `go get` 指令。我將這個範例發佈在 Github 的 [JustinSDK/goexample](https://github.com/JustinSDK/goexample) 了，因此，你可以執行以下指令：

``` go
go get github.com/JustinSDK/goexample
```

`go get` 會自行判斷該使用的協定，以這邊的例子來說，就會使用 `git` 來複製檔案庫至 src 目錄底下，結果就是 src/github.com/JustinSDK 底下，會有個 goexample 目錄，其中就是原始碼，`go get` 在下載原始碼之後，就會開始進行編譯，因此，你也會在 pkg 目錄中的 \$GOOS\_\$GOARCH 目錄底下，github.com/JustinSDK 中找到編譯好的 .a 檔案。

補充（Go 1.26 現況）：在模組模式下，`go get` 主要用於調整目前模組的依賴版本；若是安裝命令列工具，請改用 `go install module/path/cmd@version`（例如 `@latest`）。

接著，你就可以如上頭的程式撰寫 `import "github.com/JustinSDK/goexample"` 來使用這個套件。

當然，執行 `go install main` 的話，你的 pkg 目錄中的 `$GOOS_$GOARCH` 目錄，會有個 github.com/JustinSDK 目錄，裏頭放置著 goexample.a 檔案，而編譯出來的可執行檔，則會放置在 bin 目錄之中，此時，你的目錄應該會像是：

``` go
go-exercise
        ├─bin
        │      main.exe
        │
        ├─pkg
        │  └─windows_amd64
        │      └─github.com
        │          └─JustinSDK
        │                  goexample.a
        │
        └─src
            ├─github.com
            │  └─JustinSDK
            │      └─goexample
            │              .gitignore
            │              hello.go
            │              hi.go
            │              LICENSE
            │              README.md
            │
            └─main
                    main.go
```

# GOPATH 中多個路徑

如果你在 `GOPATH` 中設定多個路徑，那麼，在哪個路徑底下的 src 找到套件的原始碼，編譯出來的 .a 檔案就會放在哪個路徑底下的 pkg 目錄之中。

如果是包括程式進入點的 `main` 套件，那麼執行 `go install main` 的話，預設會放在找到 `main` 套件原始碼的 bin 目錄之中。你可以設定 `GOBIN`，指定編譯出來的可執行檔放置的目錄。

如果你在 `GOPATH` 中設定多個路徑，那麼，`go get` 複製回來的原始碼，會被放置在 `GOPATH` 中設置的第一個目錄 src 之中，同理，對應的 .a 檔案，也會是 `GOPATH` 中設置的第一個目錄的 pkg 之中。

# 有關 import

在 `import` 時預設會使用套件名稱作為呼叫套件中函式等的前置名稱，你可以在 `import` 時指定別名。例如：

``` go
package main

import f "fmt"

func main() {
    f.Println("哈囉！世界！")
}
```

若指定別名時使用 `.`，就不需要套件名稱作為前置名稱，例如：

``` go
package main

import . "fmt"

func main() {
    Println("哈囉！世界！")
}
```

你不能只是 `import x "x"` 來試圖只執行套件的初始函式，因為 Go 編譯器不允許 `import` 了某個套件而不使用，然而若指定別名時使用 `_`，則不會導入套件，只會執行套件的初始函式，也就是套件中使用 `func init()` 定義的函式。

每個套件可以有多個 `init` 定義在各個不同的原始檔案中，套件被 `import` 時會執行，若是 `main` 套件，則會在所有 `init` 函式執行完畢後，再執行 `main` 函式，Go 執行套件初始化時，不會保證套件中多個 `init` 的執行順序。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
