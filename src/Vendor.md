<div id="main" role="main" style="height: auto !important;">

<div class="header">

# vendor

</div>

  

在只有一個專案的情況下，`GOPATH` 非常合情合理而且簡單，如果有多個專案，各個專案的原始碼也可以放在同一個 `GOPATH` 之中，有著各自的套件結構，使用著來自 `GOPATH` 的非標準套件，此時整個 `GOPATH` 目錄就是一個巨大的 repository，具稱 Google 內部就是這樣的場景，才會有 `GOPATH` 這樣的設計，Go 社群中也有著「如果必須切換 `GOPATH`，大概有哪些地方不對了」的說法。

問題在於，這並不是社群或其他公司中使用 Go 的方式，如果個別專案有個別的套件，比較單純的做法是各個專案有個專用的 `GOPATH`，想要開發哪個專案，就切換至該專案使用的 `GOPATH`，然而很快地，如果有專案相依在這些個別專案上呢？將它們組織為巨大的 repository 是個做法，或者是令 `GOPATH=prj1:prj2:prj3`，prjx 是指向各專案原始碼的路徑，也就是說 `GOPATH` 會是一大串路徑結合後的產物。

在上述的設定中，維持了一個 `GOPATH` 不用切換，新專案可以加入至 `GOPATH` 最前頭，`go get` 的第三方套件會下載到最前面的路徑中，然而，若需要 prj2 也在開發中，若 prj2 需要新的第三方套件時，`go get` 卻會下載到新專案之中；在各自不同的情境中，無論怎麼調整 `GOPATH` 的順序，總是會有各自不同的問題發生。

另一方面，`GOPATH` 本身不涉及套件來源的版本問題，因此，若專案依賴的 repository 被修改了，日後建構專案就會受到影響，對依賴於Github之類來源，而且第三方套件本身非常活躍的專案來說，重新建構專案時無法有穩定的結果，這顯然是個大問題。

例如在〈[Go 套件管理](Package.html)〉中看過的例子，使用 `go get github.com/JustinSDK/goexample`，然後撰寫底下的程式：

``` prettyprint
package main

import "github.com/JustinSDK/goexample"

func main() {
    goexample.Hi()
    goexample.Hello()
}
```

這簡單的程式被發佈為一個範例了，某年某月的某一天，我修改了 goexample 的內容，讓 `Hi`、`Hello` 顯示中文並發佈到 Github 上的檔案庫，有位讀者，依舊照著〈[Go 套件管理](Package.html)〉中的說明進行操作，然而看到的不是英文，而是中文的招呼。

為了避免這個問題，通常會將下載的檔案庫複製出來，例如放到 deps 中：

``` prettyprint
project
    └─src
        ├─deps
        │  └─src
        │      └─github.com
        │          └─JustinSDK
        │              └─goexample
        │                      .gitignore
        │                      hello.go
        │                      hi.go
        │                      LICENSE
        │                      README.md
        │
        └─main
                main.go
```

問題是放到 deps 的檔案庫該怎麼用呢？其中一個方式是修改 `import`：

``` prettyprint
package main

import "deps/src/github.com/JustinSDK/goexample"

func main() {
    goexample.Hi()
    goexample.Hello()
}
```

另一個方式是透過工具修改 `GOPATH` 自動包含 deps 目錄，這類的概念主要成為了 [godep](https://github.com/tools/godep) 等工具早期在管理 Go 套件時的思考出發點。

Go 在 1.5 時實驗性地加入了 vendor，需要透過 `GO15VENDOREXPERIMENT="1"` 來啟用，1.6 預設 `GO15VENDOREXPERIMENT="1"`，1.7 拿掉 `GO15VENDOREXPERIMENT` 環境變數，使得vendor成為正式的內建特性。

簡單來說，如果你的套件中有個 vendor 資料夾，例如：

``` prettyprint
project
    └─src
        └─main
            │  main.go
            │
            └─vender
                └─github.com
                    └─JustinSDK
                        └─goexample
                                .gitignore
                                hello.go
                                hi.go
                                LICENSE
                                README.md
```

對於 `import "github.com/JustinSDK/goexample"` 來說，尋找相依套件的順序會變成 vendor -\> GOROOT 的 src -\> GOPATH 的 src。

在 vendor 推出後，`godep` 也改使用 vendor了，而 [glide](https://github.com/Masterminds/glide) 等工具，也都基於 vendor 了。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
