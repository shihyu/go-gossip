<div id="main" role="main" style="height: auto !important;">

<div class="header">

# go doc 註解即文件

</div>

  

如果你想查詢套件、函式等的說明，可以使用 `go doc` 指令。

# 查詢文件

如果你想要查詢套件的文件說明，可以使用 `go doc packageName`，例如 `go doc fmt` 可查詢 `fmt` 套件的說明，

<div class="pure-g">

<div class="pure-u-1">

<img src="images/godoc-1.JPG" class="pure-img-responsive" alt="go doc" />

</div>

</div>

可以看到，這顯示了整個套件的說明，通常我們會想要查詢套件中某個函式，這可以使用 `go doc packageName.funcName`，例如，查詢 `fmt` 中的 `Println`，可以使用 `go doc fmt.Println`：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/godoc-2.JPG" class="pure-img-responsive" alt="go doc" />

</div>

</div>

你也可以加上 `-src` 來查詢原始碼，雖然整個套件也可以查詢，不過我想，這直接開 Go 目錄的 src 中原始碼來看比較快，或許加上 `-src` 的機會，會是在查詢函式的原始碼時比較多：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/godoc-3.JPG" class="pure-img-responsive" alt="go doc" />

</div>

</div>

如同 `go xxx` 的指令說明，想要得到 `go doc` 的說明，可以使用 `go help doc` 指令。

<div class="pure-g">

<div class="pure-u-1">

<img src="images/godoc-6.JPG" class="pure-img-responsive" alt="go doc" />

</div>

</div>

# 註解即文件

實際上，`go doc` 的文件說明來自於原始碼中的註解，這樣的概念有點類似 Java 的 JavaDoc，或者是 Python 的 DocString，不過 Go 的理念是讓它更簡單，不使用特殊標記，不使用特別的格式，希望可以在沒有 `go doc` 的場合中，也可以藉由閱讀原始碼中的註解，輕易地得到文件說明。

當然，基本上還是要有一些約定，例如，在函式之前，緊接著函式的註解，中間沒有空白行，就是函式的文件說明來源。

類似地，在套件之前，緊接著套件的註解，就是套件的文件說明來源，通常，一個套件的文件說明，會是來自於套件中，一個 doc.go 中 `package` 宣告前的註解，例如，你可以在 [`fmt` 的原始碼目錄](https://go.dev/src/fmt/) 中，找到一個 [doc.go](https://go.dev/src/fmt/doc.go)，其中除了 `package fmt` 之外，沒有任何原始碼，剩下的只有註解。

除了函式、套件之外，最頂層的型態宣告、變數、常數等前緊接著的註解，都可以是文件的來源，不相鄰的註解則會被 `godoc` 忽略，如果有已知的 Bug，可以使用 `BUG()` 標示，例如 [bytes.go](https://go.dev/src/bytes/bytes.go) 中有個：

``` prettyprint
// BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
func Title(s []byte) []byte {
    ....
```

這會出現在文件的 [Bugs 區段](https://pkg.go.dev/bytes/#pkg-note-BUG)。

如果你想要從註解產生 HTML 文件（使用 `-html` 引數），那麼有幾個簡單的規則（用過 Markdown 的應該感覺有點熟悉），參考一下 Go 的原始碼，應該能很快地掌握。

基本上，`go doc` 會在 `GOROOT` 與 `GOPATH` 中的原始碼查詢註解作為文件，如果想改變查詢時的 Go 目錄，可以使用 `-goroot` 指定。

有關註解與文件間的關係，也可以進一步參考 [Effective Go 的 Commentary](https://go.dev/doc/effective_go.html#commentary)。

# godoc 文件伺服器

Go 1.2rc1 之後，曾經從 `go doc` 改用 `godoc` 指令了，不過，從 [Go 1.5 Release Notes](https://go.dev/doc/go1.5) 中看到，Go 1.5 有個新的 `go doc` 指令，專門用於命令列模式下的文件查詢，這使得 `godoc` 主要剩下文件服器的功能，因而在 Go 1.13 中，`godoc` 被移除。

如果在一個網路受限的環境，又想要在網頁上查詢文件，還是可以安裝 `godoc`（來自 `x/tools`）：

``` prettyprint
go install golang.org/x/tools/cmd/godoc@latest
```

這時執行安裝後的 `godoc`，並附帶一個 `-http` 引數指定連接埠，例如，`godoc -http=:6060`，這會在本機啟動一個 HTTP 伺服器，使用瀏覽器連接 http://localhost:6060（或 http://主機IP:6060）就可以查詢文件：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/godoc-4.JPG" class="pure-img-responsive" alt="godoc" />

</div>

</div>

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
