<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 哈囉！WebAssembly！

</div>

  

Go 1.11 實驗性地加入了 WebAssembly 的支援，這表示你可以使用 Go 來撰寫程式碼，然後令其在網頁中執行，也可以與瀏覽器互動，像是瀏覽器的 JavaScript 環境、DOM 操作等。

對 Go 開發者而言，理想的狀況下，若 Go 封裝的好，最好是可以完全忽略 JavaScript、瀏覽器環境等事實，也不需要知道 WebAssembly 的細節，然而，畢竟目前還是實驗性質，如果能認識 JavaScript、瀏覽器、WebAssembly 的特性，對使用 Go 撰寫程式並編譯為 WebAssembly 來說，是有很大的幫助的。

如果對 JavaScript、瀏覽器的細節有興趣，建議參考〈[ECMAScript 本質部份](https://openhome.cc/Gossip/ECMAScript/)〉，如果對 WebAssembly 的細節有興趣，建議參考〈[WebAssembly](https://openhome.cc/Gossip/WebAssembly/)〉文件。

無論如何，來看個簡單的 Go 程式如何編譯為 WebAssembly，首先，來個簡單的 Go 程式：

``` prettyprint
package main

func main() {
    println("Hello, WebAssembly")
}
```

再簡單也不過，在編譯為 WebAssembly 之後，`println` 的輸出預設會是瀏覽器主控台（console），接下來，若要編譯為 WebAssembly，環境變數 `GOOS` 必須設定為 `js`，`GOARCH` 必須設定為 `wasm`。

如果你是使用 Visual Studio Code，安裝了 [vscode-go](https://github.com/Microsoft/vscode-go) 擴充，可以在 settings.json 中設定：

``` prettyprint
{
    "go.toolsEnvVars": {"GOOS":"js", "GOARCH": "wasm"}
}
```

如果是要在 Visual Studio Code 開啟的終端機中設定環境變數，因為它是基於 Power Shell，可以如下設定環境變數：

``` prettyprint
$env:GOOS="js"
$env:GOARCH="wasm"
```

如果是在 Windows 的命令提示字元，就是使用 `set` 了：

``` prettyprint
SET GOOS=js
SET GOARCH=wasm
```

接下來，可以執行建構：

``` prettyprint
go build -o test.wasm main.go
```

test.wasm 是編譯出來的 WebAssembly 模組位元組碼，除了你撰寫的程式之外，根據〈[Go 1.11 Release Notes](https://goo.gl/YfaETG)〉，編譯出來的 WebAssembly 模組也包含了 goroutine、垃圾收集、maps 等功能的執行環境，最小約在 2MB 左右，壓縮後可以減至 500KB。

接下來就是開個 HTML 檔案，在當中使用 JavaScript，運用 Fetch API、WebAssembly API 等，取得、編譯、初始化 WebAssembly 模組，這些細節在〈[WebAssembly](https://openhome.cc/Gossip/WebAssembly/)〉文件都有談到。

如果想要直接有個現成的載入頁面可以使用，可以複製 Go 安裝目錄的 misc\wasm 中 wasm_exec.html 與 wasm_exec.js 到你的工作目錄之中，wasm_exec.html 裏頭寫的 JavaScript，會使用 Fetch API 來取得 test.wasm，這也是為什麼，方才編譯時指定輸出檔案名稱為 test.wasm 的原因。

如果你有安裝 Node.js，那麼可以直接搭配 wasm_exec.js 來運行 test.wasm，這會顯示 Hello, WebAssembly：

``` prettyprint
node wasm_exec.js test.wasm
```

如果要在瀏覽器中運行，你需要有個 HTTP 伺服器，例如 Node.js 的 `http-server`，在啟動之後，請求你的 http://localhost:8080/wasm_exec.html。

這會看到一個 Run 按鈕，開啟你瀏覽器上的開發者工具，然後按下網頁中的 Run 按鈕，你就會看到開發者工具中的主控台顯示了文字：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/WebAssembly-1.JPG" class="pure-img-responsive" alt="哈囉！WebAssembly！" />

</div>

</div>

要注意的是，在執行完 Go 的 `main` 之後，程式就結束了，就網頁中 Run 按鈕的事件來說，每按一次是會重新跑一次 WebAssembly 模組實例，也就是重新跑一次 `main` 流程，有時這不會是你想要的，這時就要在 Go 中以某種方式，阻斷 `main` 的流程，這之後還會看到。

這只是個體驗，之後的文件還會談到，如何操作瀏覽器中的 JavaScript、DOM，以及 Go 中定義的函式，如何能被 JavaScript 取得呼叫。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
