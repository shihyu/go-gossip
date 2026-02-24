<div id="main" role="main" style="height: auto !important;">

<div class="header">

# Go 測試套件

</div>

  

Go 本身附帶了 `testing` 套件，搭配 `go test` 指令，可以自動對套件中的程式碼進行測試，在套件中，測試程式碼必須是 \_test.go 結尾，一個套件中可以有多個 \_test.go，例如，[fmt 套件的原始碼](https://go.dev/src/fmt/) 中，可以看到 export_test.go、fmt_test.go 等，就是測試程式碼。

# 功能測試

想要使用 Go 的 `testing` 套件撰寫測試程式碼，必須 `import "testing"`，在 \_test.go 中撰寫形式 `func TestXxx(t *testing.T)` 的函式，Xxx 可以是任意名稱，例如，在 src/mymath 目錄中，寫個 basic_test.go：

``` go
package mymath

import "testing"

func TestSomething(t *testing.T) {
    // write some test
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

接著只要執行 `go test mymath`，就會自動尋找 `mymath` 套件中的 \_test.go 中 Test 開頭的函式並執行，由於目前沒撰寫任何測試內容，測試是以 PASS 結束。

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Testing-1.JPG" class="pure-img-responsive" alt="Go 測試套件" />

</div>

</div>

如果函式中使用了 `testing` 的 `Error`、`Fail` 等與失敗相關的方法，那麼測試就會失敗，例如：

``` go
package mymath

import "testing"

func TestSomething(t *testing.T) {
    t.Fail()
}
```

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Testing-2.JPG" class="pure-img-responsive" alt="Go 測試套件" />

</div>

</div>

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

如果想要在測試失敗時，留下一些訊息，可以使用 `Error` 方法，例如：

``` go
package mymath

import "testing"

func TestSomething(t *testing.T) {
    t.Error("something wrong")
}
```

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Testing-3.JPG" class="pure-img-responsive" alt="Go 測試套件" />

</div>

</div>

來實際寫個測試，例如，測試一個 `Add` 函式：

``` go
package mymath

import "testing"

func TestAdd(t *testing.T) {
    if Add(1, 2) == 3 {
        t.Log("mymath.Add PASS")
    } else {
        t.Error("mymath.Add FAIL")
    }
}
```

由於目前還沒有撰寫 `Add` 函式，因此若執行 `go test mymath` 的話，會以 \[build failed\] 收場，如果在 basic.go 撰寫了正確的 `Add` 函式：

``` go
package mymath

func Add(a, b int) int {
    return a + b
}
```

不過，如果直接執行 `go test mymath` 的話，只會顯示 ok 等字眼，不會顯示 `Log` 的訊息，想看到 `Log` 的訊息的話，必須加上 `-v` 引數（代表 verbose），例如：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Testing-4.JPG" class="pure-img-responsive" alt="Go 測試套件" />

</div>

</div>

如果 `Log` 之後接上 `Fail` 函式，那麼不加上 `-v`，也會顯示 `Log` 的訊息，實際上，`Error` 函式就是相當於先以 `Log` 顯示指定的訊息，然後再接上 `Fail` 函式。

如果想要略過測試，那麼可以使用 `Skip` 函式，例如：

``` go
package mymath

import "testing"

func TestSomething(t *testing.T) {
    t.Skip()
}

func TestAdd(t *testing.T) {
    if Add(1, 2) == 3 {
        t.Log("mymath.Add PASS")
    } else {
        t.Error("mymath.Add FAIL")
    }
}
```

`TestSomething` 中如果沒有執行 `Skip` 會是兩個 PASS 的測試結果，若如上執行了 `Skip`，會是一個 SKIP 與一個 PASS 的測試結果。例如：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Testing-5.JPG" class="pure-img-responsive" alt="Go 測試套件" />

</div>

</div>

如果你想指定某個測試，可以使用 `-run` 引數，這接受一個正則表示式，例如，若只想執行 `TestAdd`，那麼可以如下：

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Testing-6.JPG" class="pure-img-responsive" alt="Go 測試套件" />

</div>

</div>

# 效能評測

如果想進行效能評測（Benchmark），那麼 \_test.go 中，評測函式必須是 `func BenchmarkXxx(b *testing.B)` 形式，例如：

``` go
package mymath

import "testing"

func TestSomething(t *testing.T) {
    t.Skip()
}

func TestAdd(t *testing.T) {
    if Add(1, 2) == 3 {
        t.Log("mymath.Add PASS")
    } else {
        t.Error("mymath.Add FAIL")
    }
}

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}
```

為了進行評測，被測試的函式要執行多次，以求得每次執行的平均時間，要執行多次函式可以使用迴圈，並以 `b.N` 作為邊界，`b.N` 目標預設是 1000000000，評測預設會在一秒內，以越來越大的 `b.N` 執行迴圈，這是為了讓評測進入穩定狀態，以收集到可靠的評測資料；如果運行時間到了，`b.N` 目標值仍未達成，就以現有收集到的資料來回報評測結果。

你可以在執行 `go test` 時，加上 `-bench` 引數，這個引數後可以使用正則表示式，來指定符合的評測函式名稱，例如，想執行所有評測函式，可以使用 `-bench="."`：

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Testing-7.JPG" class="pure-img-responsive" alt="Go 測試套件" />

</div>

</div>

評測的結果中顯示，達到了 `b.N` 預設目標 100000000 次，平均每次迴圈花了 0.58 奈秒（nanosecond）。

如果只想進行效能評測，可以使用 `-run` 引數，這本來是用來指定要執行的測試函式，只要指定一個不符合任何測試函式的正則表示式，就可以略過所有測試，只執行評測函式了，例如：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Testing-8.JPG" class="pure-img-responsive" alt="Go 測試套件" />

</div>

</div>

方才談到，評測預設的運行時間是一秒，如果在這個時間內，無法達到 `b.N` 的目標值，可以增加這個時間，這要使用 `-benchtime` 引數，指定的格式像是 1h30s，例如：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Testing-9.JPG" class="pure-img-responsive" alt="Go 測試套件" />

</div>

</div>

如果想固定 `b.N` 的值，Go 1.12 以後可以使用 `x` 後置，例如指定執行 100000000000 次（預設 `b.N` 目標的 10 倍）並收集結果：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Testing-10.JPG" class="pure-img-responsive" alt="Go 測試套件" />

</div>

</div>

`-count` 可以指定評測重啟幾次：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Testing-11.JPG" class="pure-img-responsive" alt="Go 測試套件" />

</div>

</div>

想知道更多 Go 測試套件的細節，可以參考 [Package testing](https://pkg.go.dev/testing/) 的說明。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
