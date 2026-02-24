<div id="main" role="main" style="height: auto !important;">

<div class="header">

# Goroutine

</div>

  

在 Go 中要讓指定的流程並行執行非常簡單，只需要將流程寫在函式中，並在函式加個 `go` 就可以了，這樣我們稱之為啟動一個 Goroutine。

# 使用 Gorutine

先來看個沒有啟用 Goroutine，卻要寫個龜兔賽跑遊戲的例子，你可能是這麼寫的：

``` go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max-min) + min
}

func main() {
    flags := [...]bool{true, false}
    totalStep := 10
    tortoiseStep := 0
    hareStep := 0
    fmt.Println("龜兔賽跑開始...")
    for tortoiseStep < totalStep && hareStep < totalStep {
        tortoiseStep++
        fmt.Printf("烏龜跑了 %d 步...\n", tortoiseStep)
        isHareSleep := flags[random(1, 10)%2]
        if isHareSleep {
            fmt.Println("兔子睡著了zzzz")
        } else {
            hareStep += 2
            fmt.Printf("兔子跑了 %d 步...\n", hareStep)
        }
    }
}
```

由於程式只有一個流程，所以只能將烏龜與兔子的行為混雜在這個流程中撰寫，而且為什麼每次都先遞增烏龜再遞增兔子步數呢？這樣對兔子很不公平啊！如果可以撰寫程式再啟動兩個流程，一個是烏龜流程，一個兔子流程，程式邏輯會比較清楚。

你可以將烏龜的流程與兔子的流程分別寫在一個函式中，並用 `go` 啟動執行：

``` go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max-min) + min
}

func tortoise(totalStep int) {
    for step := 1; step <= totalStep; step++ {
        fmt.Printf("烏龜跑了 %d 步...\n", step)
    }
}

func hare(totalStep int) {
    flags := [...]bool{true, false}
    step := 0
    for step < totalStep {
        isHareSleep := flags[random(1, 10)%2]
        if isHareSleep {
            fmt.Println("兔子睡著了zzzz")
        } else {
            step += 2
            fmt.Printf("兔子跑了 %d 步...\n", step)
        }
    }
}

func main() {
    totalStep := 10

    go tortoise(totalStep)
    go hare(totalStep)

    time.Sleep(5 * time.Second) // 給予時間等待 Goroutine 完成
}
```

現在烏龜的流程與兔子的流程都清楚多了，程式的最後使用 `time.Sleep()` 讓主流程沉睡了五秒鐘，這是因為主流程一結束，所有的 Goroutine 就會停止。

# 使用 sync.WaitGroup

有沒有辦法知道 Goroutine 執行結束呢？實際上沒有任何方法可以得知，除非你主動設計一種機制，可以在 Goroutine 結束時執行通知，使用 Channel 是一種方式，這在之後的文件再說明，這邊先說明另一種方式，也就是使用 `sync.WaitGroup`。

`sync.WaitGroup` 可以用來等待一組 Goroutine 的完成，主流程中建立 `sync.WaitGroup`，並透過 `Add` 告知要等待的 Goroutine 數量，並使用 `Wait` 等待 Goroutine 結束，而每個 Goroutine 結束前，必須執行 `sync.WaitGroup` 的 `Done` 方法。

重點是，`Add` 的數字代表「之後會呼叫 `Done()` 的 Goroutine 數量」，不是程式總共會跑幾步或迴圈會執行幾次。

- `Add` 設太小：可能提早結束，甚至在多呼叫 `Done()` 時發生 `panic: sync: negative WaitGroup counter`
- `Add` 設太大：`wg.Wait()` 會一直卡住不返回

因此，我們可以使用 `sync.WaitGroup` 來改寫以上的範例：

``` go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max-min) + min
}

func tortoise(totalStep int, wg *sync.WaitGroup) {
    defer wg.Done()

    for step := 1; step <= totalStep; step++ {
        fmt.Printf("烏龜跑了 %d 步...\n", step)
    }
}

func hare(totalStep int, wg *sync.WaitGroup) {
    defer wg.Done()

    flags := [...]bool{true, false}
    step := 0
    for step < totalStep {
        isHareSleep := flags[random(1, 10)%2]
        if isHareSleep {
            fmt.Println("兔子睡著了zzzz")
        } else {
            step += 2
            fmt.Printf("兔子跑了 %d 步...\n", step)
        }
    }
}

func main() {
    wg := new(sync.WaitGroup)
    wg.Add(2)

    totalStep := 10

    go tortoise(totalStep, wg)
    go hare(totalStep, wg)

    wg.Wait()
}
```

有個 `runtime.GOMAXPROCS()` 函式，可以設定 Go 同時間能使用的 CPU 數量，它會傳回上一次設定的數字，如果傳入小於 1 的值，不會改變任何設定，因此，可以使用 `runtime.GOMAXPROCS(0)` 知道目前的設定值。想在執行時期得知可用的 CPU 數量，可以使用 `runtime.NumCPU()` 函式，因此，為了確保 Go 會使用全部的 CPU 來運行，可以這麼撰寫：

``` go
runtime.GOMAXPROCS(runtime.NumCPU()) 
```

除了透過 `runtime.GOMAXPROCS()` 設定之外，也可以透過環境變數 `GOMAXPROCS` 來設置，實際上，Go 1.5 已經預設會使用所有的 CPU 核心，不過，仍可以透過 `runtime.GOMAXPROCS()` 函式或環境變數來改變設定。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
