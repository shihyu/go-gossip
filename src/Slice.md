<div id="main" role="main" style="height: auto !important;">

<div class="header">

# 底層為陣列的 slice

</div>

  

在〈[身為複合值的陣列](http://openhome.cc/Gossip/Go/Array.html)〉中看過陣列，有的場合需要陣列，然而，若只想處理陣列中某片區域，或者以更高階的觀點看待一片資料（而不是從固定長度的陣列觀點），那麼可以使用 slice。

# 建立一個 slice

如果需要一個 slice，可以使用 `make` 函式，舉個例子來說，可以如下建立一個長度與容量皆為 5 的 slice，並傳回 `slice` 的參考，型態為 `[]int`：

``` prettyprint
package main

import "fmt"

func main() {
    s1 := make([]int, 5)
    s2 := s1
    fmt.Println(s1) // [0 0 0 0 0]
    fmt.Println(s2) // [0 0 0 0 0]
    s1[0] = 1
    fmt.Println(s1) // [1 0 0 0 0]
    fmt.Println(s2) // [1 0 0 0 0]
    s2[1] = 2
    fmt.Println(s1) // [1 2 0 0 0]
    fmt.Println(s2) // [1 2 0 0 0]
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

如上所示，`s1`、`s2` 會是個參考（Reference），型態是 `[]int`，參考至同一個 slice 實例。

透過 `s1` 或 `s2` 操作時，操作的對象是變數參考之實例，就底層來說，`make([]int, 5)` 在記憶體某位置建立了 slice 實例，而 `s1` 儲存了該位置，如果改變了 `s1` 儲存的位址值，那透過 `s1` 操作時，就會是另一個 slice 實例了。

將變數的參考對象指定給另一個變數時，底層是將儲存的位址值指定給該變數，在上例中，`s2 := s1`，就是將 `s1` 儲存的位址值，指定給 `s2`，因此透過 `s2` 操作的對象，與 `s1` 操作的對象是相同的，透過其中一個名稱來改變 slice 的元素內容，透過另一個名稱取得 slice 的元素值，就會是改變後的值。

上例也可以寫為：

``` prettyprint
package main

import "fmt"

func main() {
    var s1 []int = make([]int, 5)
    var s2 []int    // s2 這時是 nil
    s2 = s1         // 將 s1 的參考對象指定給 s2
    fmt.Println(s1) // [0 0 0 0 0]
    fmt.Println(s2) // [0 0 0 0 0]
    s1[0] = 1
    fmt.Println(s1) // [1 0 0 0 0]
    fmt.Println(s2) // [1 0 0 0 0]
    s2[1] = 2
    fmt.Println(s1) // [1 2 0 0 0]
    fmt.Println(s2) // [1 2 0 0 0]
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

在 Go 中，參考的預設零值都是 `nil`。slice 無法進行 `==` 比較，slice 唯一可以用 `==` 比較的對象是 `nil`，儲存 slice 參考的變數也無法進行 `==` 比較，若真想知道兩個變數參考的是否同一 slice，可以如下透過[反射機制](Reflect.html)來得知：

``` prettyprint
package main

import (
    "fmt"
    "reflect"
)

func main() {
    s1 := make([]int, 5)
    s2 := s1
    fmt.Println(reflect.ValueOf(s1).Pointer() == reflect.ValueOf(s2).Pointer())
}
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

若事先知道 slice 的值，也可以使用 slice 字面常量：

``` prettyprint
package main

import (
    "fmt"
    "reflect"
)

func main() {
    s1 := []int{1, 2, 3, 4, 5}
    a1 := [...]int{1, 2, 3, 4, 5}
    fmt.Println(reflect.TypeOf(s1)) // []int
    fmt.Println(reflect.TypeOf(a1)) // [5]int
}
```

注意到，建立 slice 時，方括號中是沒有 `...` 的，如果方括號中有 `...`，那會是個陣列，而不是個 slice，如上可看到的，`s1` 的型態會是 `[]int`，然而，`a1` 的型態會是 `[5]int`，`s1` 是個參考，可以指向某個 slice 實例，`s1` 本身儲存的位址值可以改變，而 `a1` 本身就是陣列，從 `a1` 的位置開始，有連續 5 個 `int` 空間可用來儲存 `int` 值，`a1` 本身的位置是固定的，無法改變。

使用 slice 字面常量時，還可以初始特定索引處的值。例如：

``` prettyprint
slice := []int{10, 20, 30, 10: 100, 20: 200}
// 顯示 [10 20 30 0 0 0 0 0 0 0 100 0 0 0 0 0 0 0 0 0 200]
fmt.Println(slice)
```

<div class="google-auto-placed" style="width: 100%; height: auto; clear: both; text-align: center;">

</div>

在上面的例子中，索引 0、1、2 被初始為 10、20、30，之後指定索引 10 為 100，索引 20 為 200，其餘未指定處初始為 `int` 零值 0。

# 從陣列或 slice 建立 slice

如果有個現成的陣列，可以從陣列中建立 slice，例如，從陣列的索引 1 到 4（不包括）建立一個 slice 的話，可以如下：

``` prettyprint
package main

import (
    "fmt"
    "reflect"
)

func main() {
    arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    slice := arr[1:4]
    fmt.Println(reflect.TypeOf(arr))   // [10]int
    fmt.Println(reflect.TypeOf(slice)) // []int
    fmt.Println(len(slice))            // 3
    fmt.Println(cap(slice))            // 9

    fmt.Println(slice)   // [2 3 4]
    fmt.Println(arr)     // [1 2 3 4 5 6 7 8 9 10]

    slice[0] = 20
    fmt.Println(slice)   // [20 3 4]
    fmt.Println(arr)     // [1 20 3 4 5 6 7 8 9 10]
}
```

在這邊可以看到，slice 的長度可以使用 `len` 得知，而容量可以使用 `cap` 函式得知，如果從陣列中切出 slice，長度是 slice 可參考的元素長度，而容量預設為從 slice 索引 0 處起算的底層陣列元素長度，如圖所示：

<div class="pure-g">

<div class="pure-u-1">

<img src="images/Slice-1.JPG" class="pure-img-responsive" alt="slice 與陣列" />

</div>

</div>

是的！slice 底層實際上還是個陣列，若兩個 slice 底層是共用同一個陣列，從一個 slice 操作，另一個 slice 取得的值也就會反映變化，也因此在上面的例子中，你透過 `slice[0]` 設定值為 20，底層的陣列也會因而反映出變化，透過 slice 指定索引取得元素值時，不能超出 slice 的長度，不然會出現 index out of range 的錯誤。

注意，單是宣告 `var slice []int` 的話，`slice` 預設零值會是 `nil`，也就是相當於 `var slice []int = nil`，也就是 `slice` 參考至 `nil`，此時 `len(slice)` 與 `cap(slice)` 的結果都會是 0，`fmt.Println` 的顯示會是 \[\]，`==` 用於 slice 時，唯一能用來比較的就是 `nil`。

方才使用 `make([]int, 5)` 函式建立 slice 時，只指定了長度為 5，而容量就預設與長度相同，實際上，可以分別指定容量與長度，例如：

``` prettyprint
package main

import "fmt"

func main() {
    slice := make([]int, 5, 10)
    fmt.Println(slice)       // [0 0 0 0 0]
    fmt.Println(len(slice))  // 5
    fmt.Println(cap(slice))  // 10
}
```

指定索引從陣列中產生 slice時，若省略冒號之後的數字，則建立的 slice，預設可取得至陣列尾端的元素，也就是長度將等於容量，例如，若 `arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}`，那麼 `arr[3:]` 的話，取得的 slice 可以存取的元素為 {4, 5, 6, 7, 8, 9, 10}，長度與容量皆為 7；如果省略冒號之前的數字，預設從索引 0 開始，例如 `arr[:2]` 會取得 {1, 2}，長度為 2，容量為 10；如果是 `arr[:]`，那麼就是取得全部陣列內容了，長度與容量皆為 10。

Go 1.2 開始，可以在 `[]` 中指定三個數字，以冒號區隔，第三個數字指定的是 slice 以原陣列哪個索引作為邊界。例如：

``` prettyprint
package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3, 4, 5}
    slice1 := arr[0:2:4]
    fmt.Println(slice1)      // [1 2]
    fmt.Println(len(slice1)) // 2
    fmt.Println(cap(slice1)) // 4
}
```

第三個數字指定的索引不能超過陣列邊界，不然會發生 invalid slice index 的錯誤。

也可以從 slice 中產生 slice，產生的 slice 底層還是同一個陣列。例如：

``` prettyprint
package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    slice1 := arr[:5]
    slice2 := slice1[:3]

    fmt.Println(slice1) // [1 2 3 4 5]
    fmt.Println(slice2) // [1 2 3]

    slice2[0] = 10
    fmt.Println(slice1) // [10 2 3 4 5]
    fmt.Println(slice2) // [10 2 3]
    fmt.Println(arr)    // [10 2 3 4 5 6 7 8 9 10]
}
```

# slice 的 append

可以使用 `append` 對 slice 附加元素，這會傳回一個 slice 的參考：

``` prettyprint
package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3, 4, 5}
    slice1 := arr[:2]
    fmt.Println(slice1)      // [1 2]
    fmt.Println(len(slice1)) // 2
    fmt.Println(cap(slice1)) // 5

    slice2 := append(slice1, 6)
    fmt.Println(slice2)      // [1 2 6]
    fmt.Println(len(slice2)) // 3
    fmt.Println(cap(slice2)) // 5

    slice2[0] = 10
    fmt.Println(slice1) // [10 2]
    fmt.Println(slice2) // [10 2 6]
    fmt.Println(arr)    // [10 2 6 4 5]
}
```

只要附加的元素沒有超出 slice 的容量，傳回的 slice 參考就會是相同的，底層也是同一陣列，因此，改變了 `slice2[0]` 的值，`slice1`、`arr` 取得結果都有了變化。

如果 `append` 的時候，附加元素超出了 slice 的容量，那麼底層會建立一個新的陣列，容量為原 slice 容量的兩倍加 2，接著將舊陣列內容複製到新陣列，然後將指定的值附加上去，`append` 的結果也會傳回新的 slice 參考。例如：

``` prettyprint
package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3, 4, 5}
    slice1 := arr[:]
    fmt.Println(slice1)      // [1 2 3 4 5]
    fmt.Println(len(slice1)) // 5
    fmt.Println(cap(slice1)) // 5

    slice2 := append(slice1, 6)
    fmt.Println(slice2)      // [1 2 3 4 5 6]
    fmt.Println(len(slice2)) // 6
    fmt.Println(cap(slice2)) // 12

    slice2[0] = 10
    fmt.Println(slice1) // [1 2 3 4 5]
    fmt.Println(slice2) // [10 2 3 4 5 6]
    fmt.Println(arr)    // [1 2 3 4 5]
}
```

在上面的例子中，由於 `slice2` 底層的陣列，與 `slice1` 無關了，因此，透過 `slice2[0]` 修改了值，並不會影響到透過 `slice1` 或 `arr` 取得的值。

如果想用 `append` 來直接附加另一個 slice，可以使用 `...`，將另一個 slice 擴展為一列引數，例如：

``` prettyprint
package main

import "fmt"

func main() {
    slice1 := []int{1, 2, 3}
    slice2 := []int{4, 5, 6}
    fmt.Println(append(slice1, slice2...))  // [1 2 3 4 5 6]
}
```

# slice 的 copy

可以使用 `copy` 函式，將一個 slice 的內容，複製至另一個 slice：

``` prettyprint
package main

import "fmt"

func main() {
    src := []int{1, 2, 3, 4, 5}
    dest := make([]int, len(src), (cap(src)+1)*2)
    fmt.Println(copy(dest, src)) // 5
    fmt.Println(src)             // [1 2 3 4 5]
    fmt.Println(dest)            // [1 2 3 4 5]

    src[0] = 10
    fmt.Println(src)  // [10 2 3 4 5]
    fmt.Println(dest) // [1 2 3 4 5]
}
```

複製時，目的 slice 的容量必須足夠，否則會發生 cap out of range 的錯誤，`copy` 函式若執行成功，會傳回複製的元素個數。

先前提到，可以從 slice 中產生 slice，然而，由於從 slice 中產生 slice，底層仍會是同一個陣列，因此，要小心一些應用場合，對於一個很大的陣列，若不斷地切出新的 slice，底層參考的陣列還是那麼大，想避免這類問題，應自行使用 `make` 建立適當大小的 slice，然後從舊 slice 使用 `copy` 複製元素值，或者使用 `append`，將舊 slice 的內容附加至新 slice，以避免這類問題。

  
  

<div class="ad336-280" style="text-align: center;">

</div>

  

<div class="recommend" style="text-align: center;">

------------------------------------------------------------------------

</div>

</div>
