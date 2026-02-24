<div id="main" role="main" style="height: auto !important;">

<div class="header">

# JavaScript 回呼 Go

</div>

  

在〈[Go 呼叫 JavaScript](JavaScript.html)〉看過如何在 Go 中取得 JavaScript 的函式，然後予以呼叫，若你曾稍微瞭解過〈[WebAssembly](https://openhome.cc/Gossip/WebAssembly/)〉，就會發覺，這跟 WebAssembly 匯入函式至 WebAssembly 的方式不同。

這是 JavaScript 的 wasm_exec.js 以及 Go 的 `syscall/js` 居中之緣故，在 wasm_exec.html 中你也可以看到載入、編譯、實例化 WebAssembly 的過程：

``` prettyprint
if (!WebAssembly.instantiateStreaming) { // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}

const go = new Go();
let mod, inst;
WebAssembly.instantiateStreaming(fetch("test.wasm"), go.importObject).then((result) => {
    mod = result.module;
    inst = result.instance;
    document.getElementById("runButton").disabled = false;
});

async function run() {
    console.clear();
    await go.run(inst);
    inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

Go 有自己的匯入物件，也就是 `go.importObject`，這個物件主要是 JavaScript 環境與 Go 編譯出來的 WebAssembly 之橋樑，將 JavaScript 的值與 Go 的結構實例作了個對應，因此，不用自己匯入某個函式，只要取得某個作為名稱空間的 JavaScript 物件，取得上頭對應的特性，像是函式，就可以在 Go 中操作。

也就是說，如果想要在 Go 中定義函式，然後在 JavaScript 中呼叫，就是將 Go 中定義的函式，設定給某個對應的 JavaScript 物件，之後就可以在 JavaScript 環境中使用了，只不過在定義時，必須留意 JavaScript 與 Go 的型態對應。

在現代版本（例如 Go 1.26）中，可以被 JavaScript 環境呼叫的 Go 函式，通常會使用 `js.Func` 搭配 `js.FuncOf` 來包裝；這個值可以設定到 JavaScript 物件上，之後由 JavaScript 呼叫。

要能被 JavaScript 呼叫的 Go 函式，常見簽章為 `func(this js.Value, args []js.Value) any`，其中 `args` 是呼叫函式時傳入的引數，你可以想像 JavaScript 函式中 `arguments` 的對應型態。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

例如，顯示加總至某個指定 DOM 物件的函式，可以如下定義：

``` prettyprint
package main

import "syscall/js"

func main() {
    // 註冊在 JavaScript 全域
    js.Global().Set("printSumTo", js.FuncOf(printSum))
    // 阻斷 main 流程
    select {}   
}

func printSum(this js.Value, args []js.Value) any {
    c1 := args[0]         // 結果顯示到這個 div 
    numbers := args[1:]   // 接下來是要加總的數字
    c1.Set("innerHTML", sum(numbers))
    return nil
}

func sum(numbers []js.Value) int {
    var sum int
    for _, val := range numbers {
        sum += val.Int()
    }
    return sum
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

`js.FuncOf` 的回呼可傳回值（會轉為對應的 JavaScript 值），不過像事件處理器這類場合，通常仍以副作用方式實現比較常見，例如改變某個 JavaScript 物件的狀態，像是這邊是改變某個 DOM 的 `innerHTML`。

因為 Go 的 `main` 執行完，模組的程式就結束了，這樣 Go 中定義的函式就沒有了，然而，事件會是在之後才發生，因而要被回呼的函式必須存活著，為了這個目的，範例中使用 `select {}` 來阻斷流程，視需求而定，你也可以用別的方式來設計某種阻斷。

至於 JavaScript 的部份，來稍微修改一下 wasm_exec.html：

``` prettyprint
<!doctype html>
<!--
Copyright 2018 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
-->
<html>

<head>
    <meta charset="utf-8">
    <title>Go wasm</title>
</head>

<body>

    <script src="wasm_exec.js"></script>
    <script>
        if (!WebAssembly.instantiateStreaming) { // polyfill
            WebAssembly.instantiateStreaming = async (resp, importObject) => {
                const source = await (await resp).arrayBuffer();
                return await WebAssembly.instantiate(source, importObject);
            };
        }

        const go = new Go();
        let mod, inst;
        WebAssembly.instantiateStreaming(fetch("test.wasm"), go.importObject).then((result) => {
            mod = result.module;
            inst = result.instance;
            document.getElementById("runButton").disabled = false;
        }).then(_ => {   // 實例化模組之後就執行
            console.clear();
            go.run(inst);       
        });
    </script>   

    <script>
        function run() {
            // 呼叫 Go 定義的回呼函式
            printSumTo(document.getElementById('c1'), 
                1, 2, 3, 4, 5);
        }
    </script>

    <button onClick="run();" id="runButton" disabled>Run</button>
    <div id="c1"></div>

</body>
</body>

</html>
```

按下 Run 之後，會呼叫 `runAndPrintSum`，這會先執行 `run` 函式，執行 WebAssembly 模組實例，對應的就是執行 Go 定義的 `main`，因為 `run` 是非同步的，接下來就會執行 `printSumTo`，因此 1 到 5 的加總結果，就會顯示到 `id` 為 `c1` 的 `div` 元素之中。

至於 WebAssembly API 的調整，想要瞭解這部份的話，可以看看〈[WebAssembly](https://openhome.cc/Gossip/WebAssembly/)〉中前三篇的說明。

故且不討論 WebAssembly API 怎麼寫，在自定義的 JavaScript 程式碼中，想要呼叫 Go 中定義的函式，其實感覺就是多了些額外的手續，而且不自然。

如果把一切都帶到 Go 中做，將 Go 中定義的函式，當成是某事件的回呼，會比較單純一些，例如：

``` prettyprint
package main

import (
    "strconv"
    "syscall/js"
)

func main() {
    // 註冊按鈕事件
    dom("runButton").Call("addEventListener", "click", js.FuncOf(cal))
    select {}
}

// 根據 id 取得 DOM 物件
func dom(id string) js.Value {
    return js.Global().Get("document").Call("getElementById", id)
}

// 按下 Run 的事件處理器
func cal(this js.Value, args []js.Value) any {
    n1, _ := inputValue("n1")
    n2, _ := inputValue("n2")
    dom("r").Set("innerHTML", n1+n2)
    return nil
}

// 取得輸入欄位值
func inputValue(id string) (int, error) {
    return strconv.Atoi(dom(id).Get("value").String())
}
```

至於 wasm_exec.html 可以如下調整：

``` prettyprint
<!doctype html>
<!--
Copyright 2018 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
-->
<html>

<head>
    <meta charset="utf-8">
    <title>Go wasm</title>
</head>

<body>
    <input id="n1"> + <input id="n2"> = <span id="r"></span><br>
    <button id="runButton" disabled>Run</button>


    <script src="wasm_exec.js"></script>
    <script>
        if (!WebAssembly.instantiateStreaming) { // polyfill
            WebAssembly.instantiateStreaming = async (resp, importObject) => {
                const source = await (await resp).arrayBuffer();
                return await WebAssembly.instantiate(source, importObject);
            };
        }

        const go = new Go();
        let mod, inst;
        WebAssembly.instantiateStreaming(fetch("test.wasm"), go.importObject).then((result) => {
            mod = result.module;
            inst = result.instance;
            document.getElementById("runButton").disabled = false;
        }).then(_ => {
            console.clear();
            go.run(inst);       
        });
    </script>   
</body>

</html>
```

這樣就可以進行頁面操作，就是個簡單的加法器：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Callback-1.JPG" class="pure-img-responsive" alt="JavaScript 回呼 Go" />

</div>

</div>

（這也許才是 Go 希望的，要你把東西都帶入 Go 中來做，JavaScript 環境的事件會呼叫 Go 的函式，然後在 Go 中計算，在 Go 中改變物件狀態、畫面等。）

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
