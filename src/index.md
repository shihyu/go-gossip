<div id="layout" style="height: auto !important;">

<a href="#menu" id="menuLink" class="menu-link"><span></span></a>

<div id="menu">

<div class="pure-menu">

<a href="../" class="pure-menu-heading">回 OPENHOME 首頁</a>

</div>

</div>

<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 語言技術：Go 語言

</div>

原文多以 Go 1.13 為例；目前已補充 Go 1.26 的差異說明（尤其是模組與 `go` 指令行為）。

<div class="ad-2" style="text-align: center;">

<div id="aswift_2_host" style="border: none; height: 280px; width: 800px; margin: 0px; padding: 0px; position: relative; visibility: visible; background-color: transparent; display: inline-block; overflow: visible;">

<div class="iframe">

</div>

</div>

</div>

起步走

語言的起步走，需要的時間最好是長一些，因為慢一點才會快一點 ...

- Go 平台概要

在開始認識語言本身之前，先來瞭解 Go 提供的基本工具集，這是 Go 的一大特色。

- - [來個 Hello, World](HelloWorld.html)
  - [Go 套件管理](Package.html)
  - [gofmt 格式化原始碼](gofmt.html)
  - [go doc 文件即註解](godoc.html)
  - [Go 測試套件](Testing.html)

<!-- -->

- 型態、變數、常數、運算子

也許看似基本，然而沒你想像中的那麼簡單。

- - [認識預定義型態](PreDeclaredType.html)
  - [變數宣告、常數宣告](VariableConstantDeclaration.html)
  - [位元組構成的字串](String.html)

- - [身為複合值的陣列](Array.html)
  - [底層為陣列的 slice](Slice.html)
  - [成對鍵值的 map](Map.html)

<!-- -->

- 運算與流程控制

`Go 有指標，switch` 很有彈性，沒有 `while`，存在 `goto` ... XD

- - [運算子](Operator.html)
  - [if ... else、switch 條件式](IfElseSwitch.html)
  - [for 迴圈](For.html)
  - [break、continue、goto](BreakContinueGoto.html)

  

函式、結構與介面

        封裝演算、定義行為、組織程式元件。

- 函式

- - [函式入門](Function.html)
  - [泛型入門（Go 1.18+）](Generics.html)
  - [一級函式](FirstClassFunction.html)
  - [匿名函式與閉包](Closure.html)
  - [defer、panic、recover](DeferPanicRecover.html)

<!-- -->

- 結構

把相關的東西放在一起。

- - [結構入門](Struct.html)
  - [結構與方法](Method.html)
  - [結構組合](StructComposition.html)

<!-- -->

- 介面

將行為定義出來。

- - [介面入門](Interface.html)
  - [型態斷言](TypeAssertion.html)
  - [介面組合](InterfaceComposition.html)

  

常用 API

        從常用 API 中學習如何撰寫 Go 程式碼。

- 基本 IO

從 `io.Reader`、`io.Writer` 開始認識。

- - [從標準輸入、輸出認識 io](StdOutInErr.html)
  - [io.Reader、io.Writer](ReaderWriter.html)
  - [bufio 套件](bufio.html)
  - [檔案操作](File.html)

<!-- -->

- error 處理

到處都在 `if err != nil`？

- - [err 是否 nil？](ErrNil.html)
  - [錯誤的比對](ErrorComparison.html)
  - [errors 套件](errors.html)

<!-- -->

- 資料結構

`sort`、`list`、`heap` 與 `ring` 套件。

- - [sort 套件](Sort.html)
  - [list 套件](List.html)
  - [heap 套件](Heap.html)
  - [ring 套件](Ring.html)

<!-- -->

- 文字

有關字串、位元組、規則表示式等的處理。

- - [strconv、strings 套件](StrconvStrings.html)
  - [bytes 套件](Bytes.html)
  - [unicode 套件](Unicode.html)
  - [編碼轉換](XText.html)
  - [Match 比對](https://openhome.cc/Gossip/Regex/MatchGo.html)
  - [Regexp 實例](https://openhome.cc/Gossip/Regex/RegexpGo.html)

<!-- -->

- 反射

探測資料的結構與相關數值。

- - [反射入門](Reflect.html)
  - [結構欄位標籤](FieldTag.html)

<!-- -->

- 並行

簡單的並行模型。

- - [Goroutine](Goroutine.html)
  - [Channel](Channel.html)

  

其他

        一些雜七雜八的東西，暫時放這分類。

- 相依管理

go module 能終結混亂嗎？

- - [vendor](Vendor.html)
  - [模組入門](Module.html)

<!-- -->

- WebAssembly 支援

Go 也可以在瀏覽器裏跳舞？

- - [哈囉！WebAssembly！](WebAssembly.html)
  - [Go 呼叫 JavaScript](JavaScript.html)
  - [JavaScript 回呼 Go](Callback.html)

  

附錄

- [Go 官方套件說明文件](http://c-faq.com/)
- [How to Write Go Code](https://go.dev/doc/code.html)
- [Go Commands](https://pkg.go.dev/cmd/)
- [Effective Go](https://go.dev/doc/effective_go.html)
- ...

<div class="ad336-280" style="text-align: center;">

<div id="aswift_3_host" style="border: none; height: 280px; width: 336px; margin: 0px; padding: 0px; position: relative; visibility: visible; background-color: transparent; display: inline-block;">

</div>

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

<div id="aswift_4_host" style="border: none; height: 480px; width: 800px; margin: 0px; padding: 0px; position: relative; visibility: visible; background-color: transparent; display: inline-block;">

</div>

</div>

</div>

</div>

<div class="analytics">

</div>

<div id="aswift_0_host" style="border: none; height: 0px; width: 0px; margin: 0px; padding: 0px; position: relative; visibility: visible; background-color: transparent; display: inline-block;">

<div class="iframe">

</div>

</div>

|     |     |
|-----|-----|
|     |     |

<div class="ad-300-flex ad_zone external-add leaderboard_ad_top_responsive logoutAd" style="width: 1px; height: 1px; position: absolute; left: -10000px; top: -10006px;">

</div>

<div class="iframe">

</div>

<div id="aswift_5_host" style="border: none !important; height: 100vh !important; width: 100vw !important; margin: 0px !important; padding: 0px !important; position: relative !important; visibility: visible !important; background-color: transparent !important; display: inline-block !important; inset: auto !important; clear: none !important; float: none !important; max-height: none !important; max-width: none !important; opacity: 1 !important; overflow: visible !important; vertical-align: baseline !important; z-index: auto !important;">

<div class="iframe">

</div>

</div>
