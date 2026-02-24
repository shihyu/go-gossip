<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 模組入門

</div>

  

Go 在 1.11 時內建了實驗性的模組管理功能，並藉由 `GO111MODULE` 來決定是否啟用，可設定的值是 `auto`（1.11 ~ 1.13 預設）、`on` 與 `off`。

若使用 Go 1.13，當設定值是 `auto`，執行建構指令時，會看看是否有個 go.mod 檔案（用來定義依賴的模組），若有就使用 Go 模組功能，如果沒有 go.mod 檔案，就採用舊式 `GOPATH`、vendor 的特性。

當設定值為 `on` 時，就是始終使用 Go 模組功能（從 1.12 之後，go.mod 可以在必要時再新增），模組下載後會自動安裝至 `GOPATH`。

go.mod 不可以在 `GOPATH` 之中。

設定值為 `off` 時就是不使用 Go 模組功能。

例如，現在有個 pkgfoo 釋出了 [v1.0.0](https://github.com/JustinSDK/pkgfoo/tree/v1.0.0) 版，而你打算基於它寫個 go-exercise，go-exercise 資料夾中有個 src/main/main.go：

``` prettyprint
package main

import "github.com/JustinSDK/pkgfoo"

func main() {
    kgfoo.Hi()
    pkgfoo.Hello()
}
```

現在進入你的 go-exercise 資料夾底下，執行 `go mod init go-exercise`，這會建立一個 go.mod，使用 Go 1.13 的話，預設內容是：

``` prettyprint
module go-exercise

go 1.13
```

從 Go 1.12 開始，預設的 go.mod 中會有版本字段，放置了 go.mod 的資料夾稱為模組根（module root）目錄，通常就是一個 repository 的根目錄，該目錄下的全部套件都屬於該模組（除了那些本身包含 go.mod 檔案的子目錄之外）。

接著執行 `go build -o bin/main.exe src/main/main.go`，這時會自動找出 `import` 陳述，執行了套件的下載並完成建構，此時會顯示以下訊息：

``` prettyprint
go: finding github.com/JustinSDK/pkgfoo v1.0.0
go: downloading github.com/JustinSDK/pkgfoo v1.0.0
go: extracting github.com/JustinSDK/pkgfoo v1.0.0
```

而 go.mod 也有了底下內容：

``` prettyprint
module exercise

go 1.13

require github.com/JustinSDK/pkgfoo v1.0.0
```

go.mod 定義了相依的套件與版本，若是第一次執行 `go build`，那麼總是會下載最新版本，你也可以自行編輯 go.mod 的內容，來取得想要的版本，另外你也會發現多了個 go.sum，其中包含了套件的 hash 等訊息，這用來確認取得的是正確的版本：

``` prettyprint
github.com/JustinSDK/pkgfoo v1.0.0 h1:XOi67njsT9pcRrsT40Oi3LCA3b1TyIxHd6+9ceGwa0U=
github.com/JustinSDK/pkgfoo v1.0.0/go.mod h1:5PAHGmqvfj2XbzxxOeiJJjOflE/p6zTVRFfaiEeSn1w=
```

接著在執行建構出來的可執行檔時會看到：

``` prettyprint
Hi
Helo
```

喔！Hello 少了一個小寫的 l，這是一個小 bug，在修正之後，發佈了 [v1.0.1](https://github.com/JustinSDK/pkgfoo/tree/v1.0.1)：

現在 appfoo 為了要能取得更新，可以使用 `go get -u`，這會昇級到最新的 MINOR 或 PATCH 版本，像是從 1.0.0 到 1.0.1，或者是 1.0.0 到 1.1.0，是的，這採用的是 [Semantic Versioning](https://semver.org/)；若是使用 `go get -u=patch all`，會將用到的套件昇級至最新的 PATCH 版本，像是從 1.0.0 到 1.0.1；若是使用 `go get package@version`，可以指定昇級至某個版本號，例如 `go get github.com/JustinSDK/pkgfoo@v1.0.1`，然而，不建議以此方式昇級至新的 MAJOR 版本，原因後述。

在這邊因為只是小 bug 更新，就使用 `go get -u=patch all`，這會看到底下的訊息：

``` prettyprint
go: finding github.com/JustinSDK/pkgfoo v1.0.1
go: downloading github.com/JustinSDK/pkgfoo v1.0.1
go: extracting github.com/JustinSDK/pkgfoo v1.0.1
```

go.mod 的內容也更新了（go.sum 也會更新）：

``` prettyprint
module go-exercise

go 1.13

require github.com/JustinSDK/pkgfoo v1.0.1
```

重新執行 `go build`，就會顯示正確的訊息了：

``` prettyprint
Hi
Hello
```

假設現在 pkgfoo 中的訊息都改成中文，並更新為 v2.0.0 了，雖然可以使用 `go get github.com/JustinSDK/pkgfoo@v2.0.0` 來下載最新版本，然而會出現 +incompatible 字樣：

``` prettyprint
go: finding github.com/JustinSDK/pkgfoo v2.0.0
go: downloading github.com/JustinSDK/pkgfoo v2.0.0+incompatible
```

雖然可以順利建構，執行時也會是最新版本的結果，然而，若要昇級至最新的 MAJOR 版本，依賴的套件，必須明確地屬於某個模組，因此，pkgfoo 中必須有個 go.mod，並定義版本資訊：

``` prettyprint
module github.com/JustinSDK/pkgfoo/v2
```

go.mod 在加入了 pkgfoo 之後，發佈了 [v2.0.0](https://github.com/JustinSDK/pkgfoo/tree/v2.0.0) ，現在 appfoo 打算使用這 v2.0.0，可以在 `import` 時指定：

``` prettyprint
package main

import "github.com/JustinSDK/pkgfoo/v2"

func main() {
    pkgfoo.Hi()
    pkgfoo.Hello()
}
```

直接 `go build -o bin/main.exe src/main/main.go`，就會看到下載了 v2.0.0：

``` prettyprint
go: finding github.com/JustinSDK/pkgfoo/v2 v2.0.0
go: downloading github.com/JustinSDK/pkgfoo/v2 v2.0.0
go: extracting github.com/JustinSDK/pkgfoo/v2 v2.0.0
```

而且你可以看到 appfoo 的 go.mod 更新為：

``` prettyprint
module go-exercise

go 1.13

require (
    github.com/JustinSDK/pkgfoo v1.0.1
    github.com/JustinSDK/pkgfoo/v2 v2.0.0
)
```

現在它依賴在…兩個版本？是的，事實上，你也可以同時在 appfoo 中使用：

``` prettyprint
package main

import "github.com/JustinSDK/pkgfoo/v2"
import pkgfooV1 "github.com/JustinSDK/pkgfoo"

func main() {
    pkgfoo.Hi()
    pkgfoo.Hello()
    pkgfooV1.Hi()
    pkgfooV1.Hello()        
}
```

不同模組版本的套件，被視為不同的套件，上面的程式執行過時會顯示：

``` prettyprint
嗨
哈囉
Hi
Hello
```

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
