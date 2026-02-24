<div id="main" role="main" style="height: auto !important;">

<div class="header">

# Go 呼叫 JavaScript

</div>

  

Go 社群中有不少人直言，Go 支援 WebAssembly 就是要取代 Javascript，雖然我個人覺得，這就姑且當成是個崇高的理想就好，不過這也表示，在編譯為 WebAssembly 之後，可以呼叫 JavaScript 或操作 DOM，自然也是 Go 應該照料之事。`syscall/js` 在 Go 1.11 時期以實驗性姿態登場，而在現代版本（例如 Go 1.26）的 `js/wasm` 開發中，仍是重要的橋接套件。

Go 與 JavaScript 畢竟是兩個不同的語言，各擁有不同的資料型態與結構，因而必須先知道，兩個語言間的型態如何對應，這主要定義在 `syscall/js` 套件的 js.go 中。

例如，`js.Value` 結構代表 JavaScript 中的值，定義有 `Get` 與 `Set` 方法，可以對物件上的特性存取；若想存取的對象實際上是陣列，可以使用 `Index`、`SetIndex` 並指定索引；若對象是個函式，可以使用 `Invoke` 指定引數來呼叫，若想呼叫的是物件上的方法，可以使用 `Call` 指定方法名稱與呼叫時的引數等。

在 js.go 中預先定義了一些 `js.Value` 的實例，可以透過公開的 `js.Undefined`、`js.Null`、`js.Global` 等函式呼叫取得。

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

因而，你可以開啟〈[哈囉！WebAssembly！](WebAssembly.html)〉中談到的 wasm_exec.html，在 `<button onClick="run();" id="runButton" disabled>Run</button>` 標籤底下加上 `<div id="c1">Hello, WebAssembly!</div>`：

``` go
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
        WebAssembly API 等... 略
    </script>

    <button onClick="run();" id="runButton" disabled>Run</button>
    <div id="c1">Hello, WebAssembly!</div>
</body>

</html>
```

若想撰寫 Go 來取得對應的 DOM 物件，並在主控台顯示 `innerHTML` 特性值，可以如下撰寫：

``` go
package main

import "syscall/js"

func main() {
    window := js.Global()                       // 取得全域的 window
    doc := window.Get("document")               // 相當於 window.document
    c1 := doc.Call("getElementById", "c1")      // 相當於 document.getElementById('c1')
    innerHTML := c1.Get("innerHTML").String()   // 相當於 c1.innerHTML
    println(innerHTML)
}
```

這邊特意使用數個變數，代逐一對照取得的各是哪個 JavaScript 值，實際上當然可以直接寫成底下：

``` go
package main

import "syscall/js"

func main() {
    innerHTML :=
        js.Global().
            Get("document").
            Call("getElementById", "c1").
            Get("innerHTML").
            String()
    println(innerHTML)
}
```

也就是相當於在 JavaScript 中撰寫 `document.getElementById("c1").innerHTML`；在編譯為 WebAssembly、使用瀏覽器連線至網頁之後，按下 Run 按鈕，就可以取得目標 `c1` 的 `innerHTML`：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/JavaScript-1.JPG" class="pure-img-responsive" alt="Go 呼叫 JavaScript" />

</div>

</div>

類似地，如果想在 Go 中呼叫瀏覽器提供的 `alert` 全域函式，可以如下撰寫：

``` go
package main

import "syscall/js"

func main() {
    alert := js.Global().Get("alert")
    // 相當於 alert('Hello, WebAssembly!')
    alert.Invoke("Hello, WebAssembly!")
}
```

在編譯為 WebAssembly、使用瀏覽器連線至網頁之後，按下 Run 按鈕，會令瀏覽器出現警示對話方塊：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/JavaScript-2.JPG" class="pure-img-responsive" alt="Go 呼叫 JavaScript" />

</div>

</div>

因此，如果有個自定義的 JavaScript 函式，而你想在 Go 中呼叫它，就是看看，那個函式是在哪個物件之上，想辦法取得該物件，之後就可以加以呼叫了，例如：

``` go
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
        WebAssembly API 等... 略
    </script>
    <script>
        function hi_wasm(name) {
            document.getElementById('c1').innerHTML = 'Hi, ' + name;
        }
    </script>

    <button onClick="run();" id="runButton" disabled>Run</button>
    <div id="c1">Hello, WebAssembly!</div>
</body>

</html>
```

在上例中，`hi_wasm` 函式實際上是 `window` 的一個特性，因此在 Go 中可以這麼呼叫：

``` go
package main

import "syscall/js"

func main() {
    hi_wasm := js.Global().Get("hi_wasm")
    hi_wasm.Invoke("WebAssembly")
}
```

在編譯為 WebAssembly、使用瀏覽器連線至網頁之後，按下 Run 按鈕，就會將 `c1` 的文字改變為 Hi, WebAssembly。

如果自訂的 JavaScript 函式有傳回值的話，那會成為 `Invoke` 方法的傳回值，然而記得，JavaScript 的值在 Go 中是對應至 `js.Value`，`Invoke` 的傳回型態正是 `js.Value`，取得之後，就看它代表著什麼 JavaScript 值（數值、字串、陣列、函式？），再進一步操作。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
