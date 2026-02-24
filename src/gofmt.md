<div id="main" role="main" style="height: auto !important;">

<div class="header">

# gofmt 格式化原始碼

</div>

  

如果你是個有點責任感的開發者，在新接觸一門語言的時候，應該會問一個問題：「我該用什麼格式寫程式？」所以了，在 Go 裏要用什麼格式寫程式？這個問題可以直接請 `gofmt` 來幫你解答。

# 使用 gofmt

使用 `gofmt` 最簡單的方式之一，就是直接執行 `gofmt`，這會接受你在標準輸入（Standard input）鍵入的的程式碼，輸入完成後按下 Ctrl + Z，`gofmt` 就會告訴你怎麼要用什麼格式，例如，來個 Hello, World：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/gofmt-1.JPG" class="pure-img-responsive" alt="gofmt" />

</div>

</div>

在上頭的例子中，我故意製作了一些其他的格式慣例，而從輸出中可以看到 `gofmt` 建議的格式會是什麼樣子，例如，Go 建議的格式是使用 Tab 縮排，你鍵入的程式碼不用是完整的程式，也可以只是個陳述句，例如：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/gofmt-2.JPG" class="pure-img-responsive" alt="gofmt" />

</div>

</div>

你也可以指定檔案，格式化後的結果會輸出至標準輸出（Standard output），或者是一個目錄，這會遞迴地將其中的 .go 檔案讀入並格式化後，輸出至標準輸出，也可以加上 `-w` 指定以格式化後的結果重寫原有的 .go 文件。

有些格式在 Go 中是強制的，例如，大括號 `{}` 必須是右上左下的形式，因此，如果你將大括號置於同一側，執行 `gofmt` 就會得到錯誤訊息：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/gofmt-3.JPG" class="pure-img-responsive" alt="gofmt" />

</div>

</div>

# gofmt 簡單重構

`gofmt` 也可以使用 `-r` 指定規則來實現簡單的重構，例如在〈[Command gofmt](https://pkg.go.dev/cmd/gofmt/)〉文件說明中，有個 `gofmt -r '(a) -> a' -l *.go` 可以列出 .go 檔案中有多餘括號的檔案名稱（透過 `-l` 引數來列出名稱），要直接移除 .go 檔案中多餘的括號並重寫原有的 .go 檔案，可以使用 `gofmt -r '(a) -> a' -w *.go`。

`-r` 接受的規則是 `pattern -> replacement`，其中 `pattern` 與 `replacement` 必須是合法的 Go 語法，而單一、小寫的字元會被作為萬用字元（Wildcard），因此，如果有個原始碼內容是：

``` prettyprint
package goexample

func Hello(who string) {
    var helloWho = ("Hello, ") + (who)
}
```

執行過後，會產生以下的結果：

``` prettyprint
package goexample

func Hello(who string) {
    var helloWho = "Hello, " + who
}
```

再來看個無聊的例子，如果你的程式碼是：

``` prettyprint
package goexample

func Hello(who string) {
    var helloWho = who + "Hello, "
}
```

若你想要 `gofmt` 幫你改成：

``` prettyprint
package goexample

func Hello(who string) {
    var helloWho = "Hello, " + who
}
```

你可以執行 `gofmt -r 'a + "Hello, " -> "Hello, " + a' -w *.go`，甚至 `gofmt -r 'a + b -> b + a' -w` 來達到這個目的。

`gofmt` 還有個 `-s` 引數，可以嘗試為你簡化原始碼，你可以看看〈[Command gofmt](https://pkg.go.dev/cmd/gofmt/)〉文件中的說明，瞭解它會做哪些簡化，文件中也談到，簡化後的 Go 原始碼，可能會與舊版的 Go 不相容。

至於方才提及的 `goimports`，在 Go 1.18+ / 1.26 的常見做法是使用 `go install` 搭配版本號來安裝，例如：

``` prettyprint
go install golang.org/x/tools/cmd/goimports@latest
```

# go fmt

`go` 本身也可以附帶 `fmt`，也就是使用 `go fmt` 的方式來進行程式碼的格式化，`go fmt` 內部使用 `gofmt`，可以使用 `-n` 來顯示要被使用或已被使用的指令：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/gofmt-5.JPG" class="pure-img-responsive" alt="go fmt" />

</div>

</div>

可以看到，`go fmt` 包裝了 `gofmt -l -w` 指令，簡化了常用的指令輸入，你只要指定套件就可以了。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
